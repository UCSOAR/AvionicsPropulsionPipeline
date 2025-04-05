package staticfire

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Helper function to create a string from the current point of a scanner until the end of an LVM header.
//
// Parameters:
//   - scanner: The scanner to read from.
//
// Returns:
//   - string: The string containing the header.
func readUntilEOH(scanner *bufio.Scanner) string {
	var headerBuilder strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if line == AssertedEndOfHeader {
			break
		}

		headerBuilder.WriteString(line)
		headerBuilder.WriteRune('\n')
	}

	return headerBuilder.String()
}

// Helper function to filter out the comment column from a list of column names.
//
// Parameters:
//   - columnNames: The list of column names to filter.
//
// Returns:
//   - []string: The list of column names with the comment column removed.
func filterCommentColumn(columnNames []string) []string {
	if columnNames[len(columnNames)-1] == AssertedCommentColumnName {
		return columnNames[:len(columnNames)-1]
	}

	return columnNames
}

// Parses an LVM file into a cache tree.
//
// Parameters:
//   - reader: The reader to read the LVM content from.
//
// Returns:
//   - CacheTree: The cache tree parsed from the LVM content.
//   - error: An error if the LVM content could not be parsed, or nil if the operation was successful.
func ParseIntoCacheTree(reader io.Reader) (CacheTree, error) {
	lvmScanner := bufio.NewScanner(reader)

	// Assert first line
	if !lvmScanner.Scan() {
		return CacheTree{}, nil
	}

	if line := strings.TrimSpace(lvmScanner.Text()); line != AssertedFirstLine {
		return CacheTree{}, fmt.Errorf("first line %s does not match expected %s", line, AssertedFirstLine)
	}

	// Read forward to the end of the first header (entry header)
	rawEntryHeaderText := readUntilEOH(lvmScanner)
	entryHeader, err := ParseEntryHeader(rawEntryHeaderText)

	if err != nil {
		return CacheTree{}, err
	}

	// Read forward to the end of the second header (channel header)
	rawChannelHeaderText := readUntilEOH(lvmScanner)
	channelHeader, err := ParseChannelHeader(rawChannelHeaderText)

	if err != nil {
		return CacheTree{}, err
	}

	// Count number of X and Y columns
	xColumnCount := 1 // For one X column
	yColumnCount := channelHeader.ChannelCount

	// For multiple X columns
	if entryHeader.XColumns == XColumnsMulti {
		// One X column for each channel
		xColumnCount = channelHeader.ChannelCount
	}

	yColumnMetadata := make([]YColumnMetadata, yColumnCount)
	xColumns := make([]ColumnNode, xColumnCount)
	yColumns := make([]ColumnNode, yColumnCount)

	// Initialize Y column metadata
	for i := range yColumnCount {
		yColumnMetadata[i] = YColumnMetadata{
			Samples:    channelHeader.Samples[i],
			Date:       channelHeader.Dates[i],
			UnitLabel:  channelHeader.YUnitLabels[i],
			XDimension: channelHeader.XDimensions[i],
		}
	}

	// Read column names
	if !lvmScanner.Scan() {
		return CacheTree{}, errors.New("no columns found")
	}

	columnNames := filterCommentColumn(strings.Split(lvmScanner.Text(), string(entryHeader.Separator)))

	// Check that number of columns matches expectations
	totalColumnCount := xColumnCount + yColumnCount

	if len(columnNames) != xColumnCount+yColumnCount {
		return CacheTree{}, fmt.Errorf("expected %d columns, got %d", totalColumnCount, len(columnNames))
	}

	// Read until end of file
	for lvmScanner.Scan() {
		line := lvmScanner.Text()
		values := strings.Split(line, string(entryHeader.Separator))

		// Check if the last column is a comment
		if len(values) > totalColumnCount {
			// Drop the comment column
			values = values[:totalColumnCount]
		} else if len(values) < totalColumnCount {
			return CacheTree{}, fmt.Errorf("too few values in line: %s", line)
		}

		// For each column, parse the value and add it as a new row
		xColumnIndex := 0

		for columnIndex, value := range values {
			floatValue, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return CacheTree{}, fmt.Errorf("failed to parse value: %s", value)
			}

			// A column is an X column if:
			// - There is only one X column and the current column index is 0
			// - There are multiple X columns and the current column index is even
			if (entryHeader.XColumns == XColumnsOne && columnIndex == 0) || (entryHeader.XColumns == XColumnsMulti && columnIndex%2 == 0) {
				xColumns[xColumnIndex].Rows = append(xColumns[xColumnIndex].Rows, floatValue)
				xColumnIndex++
			} else {
				// Otherwise, it is a Y column
				yColumnIndex := columnIndex - xColumnIndex
				yColumns[yColumnIndex].Rows = append(yColumns[yColumnIndex].Rows, floatValue)
			}
		}
	}

	// Generate column names
	xColumnNames := make([]string, 0, xColumnCount)
	yColumnNames := make([]string, 0, yColumnCount)

	for i, columnName := range columnNames {
		if strings.HasPrefix(columnName, AssertedXColumnPrefix) {
			if entryHeader.XColumns == XColumnsOne {
				// For one X column, the X column is the first column
				xColumnNames = append(xColumnNames, columnName)
			} else {
				if i == len(columnNames)-1 {
					return CacheTree{}, fmt.Errorf("the X column name %s is the last column", columnName)
				}

				// The X column associated with the Y column after it
				xColumnNames = append(xColumnNames, "(X) "+columnNames[i+1])
			}
		} else {
			yColumnNames = append(yColumnNames, columnName)
		}
	}

	// generate total rows in lvm file
	var totalRows int

	if len(xColumns) > 0 {
		totalRows = len(xColumns[0].Rows)
	}

	tree := CacheTree{
		PreviewMetadata: PreviewMetadata{
			ResultTimestamp: TimestampMetadata{
				Date: entryHeader.Date,
				Time: entryHeader.Time,
			},
			Operator:     entryHeader.Operator,
			XColumnNames: xColumnNames,
			YColumnNames: yColumnNames,
			TotalRows:    totalRows,
		},
		YColumnMetadata: yColumnMetadata,
		XColumnNodes:    xColumns,
		YColumnNodes:    yColumns,
	}

	fmt.Print(tree)

	return tree, nil
}

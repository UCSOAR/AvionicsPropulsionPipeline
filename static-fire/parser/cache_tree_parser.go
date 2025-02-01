package parser

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	caching "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching"
)

// Move a scanner forward until the end of the header.
// Returns the scanned text.
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

func filterCommentColumn(columnNames []string) []string {
	if columnNames[len(columnNames)-1] == "Comment" {
		return columnNames[:len(columnNames)-1]
	}

	return columnNames
}

func ParseIntoCacheTree(lvmFile multipart.File) (caching.CacheTree, error) {
	scanner := bufio.NewScanner(lvmFile)

	// Assert first line
	if !scanner.Scan() {
		return caching.CacheTree{}, nil
	}

	if line := strings.TrimSpace(scanner.Text()); line != AssertedFirstLine {
		return caching.CacheTree{}, fmt.Errorf("First line %s does not match expected %s", line, AssertedFirstLine)
	}

	// Read forward to the end of the first header (entry header)
	rawEntryHeaderText := readUntilEOH(scanner)
	entryHeader, err := ParseEntryHeader(rawEntryHeaderText)

	if err != nil {
		return caching.CacheTree{}, err
	}

	// Read forward to the end of the second header (channel header)
	rawChannelHeaderText := readUntilEOH(scanner)
	channelHeader, err := ParseChannelHeader(rawChannelHeaderText)

	if err != nil {
		return caching.CacheTree{}, err
	}

	// Count number of X and Y columns
	xColumnCount := 1 // For one X column
	yColumnCount := channelHeader.ChannelCount

	// For multiple X columns
	if entryHeader.XColumns == XColumnsMulti {
		// One X column for each channel
		xColumnCount = channelHeader.ChannelCount
	}

	xColumns := make([]caching.XColumnNode, xColumnCount)
	yColumns := make([]caching.YColumnNode, yColumnCount)

	// Read column names
	if !scanner.Scan() {
		return caching.CacheTree{}, fmt.Errorf("No columns found")
	}

	columnNames := filterCommentColumn(strings.Split(scanner.Text(), string(entryHeader.Seperator)))

	// Check that number of columns matches expectations
	totalColumnCount := xColumnCount + yColumnCount

	if len(columnNames) != xColumnCount+yColumnCount {
		return caching.CacheTree{}, fmt.Errorf("Expected %d columns, got %d", totalColumnCount, len(columnNames))
	}

	// Read until end of file
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, string(entryHeader.Seperator))

		// Check if the last column is a comment
		if len(values) > totalColumnCount {
			// Drop the comment column
			values = values[:totalColumnCount]
		} else if len(values) < totalColumnCount {
			return caching.CacheTree{}, fmt.Errorf("Too few values in line: %s", line)
		}

		// For each column, parse the value and add it as a new row
		xColumnIndex := 0

		for columnIndex, value := range values {
			floatValue, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return caching.CacheTree{}, fmt.Errorf("Failed to parse value: %s", value)
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
				yColumns[yColumnIndex] = caching.YColumnNode{
					Rows:       append(yColumns[yColumnIndex].Rows, floatValue),
					Samples:    channelHeader.Samples[yColumnIndex],
					Date:       channelHeader.Dates[yColumnIndex],
					UnitLabel:  channelHeader.YUnitLabels[yColumnIndex],
					XDimension: channelHeader.XDimensions[yColumnIndex],
				}
			}
		}
	}

	// Generate column names
	xColumnNames := make([]string, xColumnCount)
	yColumnNames := make([]string, yColumnCount)

	for i, columnName := range columnNames {
		if strings.HasPrefix(columnName, AssertedXColumnPrefix) {
			if i == len(columnNames)-1 {
				return caching.CacheTree{}, fmt.Errorf("X column name %s is the last column", columnName)
			}

			xColumnNames = append(xColumnNames, "X-"+columnNames[i+1])
		} else {
			yColumnNames = append(yColumnNames, columnName)
		}
	}

	tree := caching.CacheTree{
		PreviewMetadata: caching.PreviewMetadata{
			ResultTimestamp: caching.TimestampMetadata{
				Date: entryHeader.Date,
				Time: entryHeader.Time,
			},
			Operator:     entryHeader.Operator,
			XColumnNames: xColumnNames,
			YColumnNames: yColumnNames,
		},
		XColumnNodes: xColumns,
		YColumnNodes: yColumns,
	}

	return tree, nil
}

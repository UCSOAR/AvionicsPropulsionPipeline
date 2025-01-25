package staticFireParser

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Represents a parsed seperated value file.
type ParsedSv struct {
	ColumnCount uint64
	ColumnNames []string    // Guaranteed to have the same length as `ColumnCount`
	Data        [][]float64 // Guaranteed to have the same column count as `ColumnCount` (row major order)
}

// Parses a seperated value file.
// Returns a struct representing the parsed seperated value file.
func ParseSv(rawSvText string, delimiter rune) (ParsedSv, error) {
	reader := strings.NewReader(rawSvText)
	scanner := bufio.NewScanner(reader)

	// Read columns
	if !scanner.Scan() {
		return ParsedSv{}, fmt.Errorf("No columns found")
	}

	columns := strings.Split(scanner.Text(), string(delimiter))
	columnCount := uint64(len(columns))

	// Read data
	var data [][]float64

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, string(delimiter))

		var row []float64

		for _, value := range values {
			floatValue, err := strconv.ParseFloat(value, 64)

			// Replace invalid values with 0
			if err != nil {
				floatValue = 0
			}

			row = append(row, floatValue)
		}

		// Fill rest of row with 0
		rowColumns := uint64(len(row))

		for i := rowColumns; i < columnCount; i++ {
			row = append(row, 0)
		}

		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		return ParsedSv{}, err
	}

	sv := ParsedSv{
		ColumnCount: columnCount,
		ColumnNames: columns,
		Data:        data,
	}

	return sv, nil
}

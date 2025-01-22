package staticFireParser

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ParsedSv struct {
	ColumnCount uint64
	ColumnNames []string
	Data        [][]float64 // Row major order
}

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

			if err != nil {
				floatValue = math.NaN()
			}

			row = append(row, floatValue)
		}

		// Fill rest of row with NaN
		rowColumns := uint64(len(row))

		for i := rowColumns; i < columnCount; i++ {
			row = append(row, math.NaN())
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

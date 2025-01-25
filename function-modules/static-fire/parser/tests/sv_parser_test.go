package staticFireParserTests

import (
	"reflect"
	"testing"

	staticFireParser "example.com/static-fire-parser"
)

func TestValidSvParsesCorrectly(t *testing.T) {
	rawSvText := `Column1	Column2	Column3
1.1	2.2	3.3
4.4	5.5	6.6
7.7	8.8	9.9`

	expected := staticFireParser.ParsedSv{
		ColumnCount: 3,
		ColumnNames: []string{"Column1", "Column2", "Column3"},
		Data: [][]float64{
			{1.1, 2.2, 3.3},
			{4.4, 5.5, 6.6},
			{7.7, 8.8, 9.9},
		},
	}

	result, err := staticFireParser.ParseSv(rawSvText, '\t')

	if err != nil {
		t.Errorf("ParseSv() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseSv() = %v, want %v", result, expected)
	}
}

func TestValidSvWithNaNParsesCorrectly(t *testing.T) {
	rawSvText := `Column1	Column2	Column3
1.1	2.2	3.3
4.4	5.5	fish
7.7	8.8	9.9`

	expected := staticFireParser.ParsedSv{
		ColumnCount: 3,
		ColumnNames: []string{"Column1", "Column2", "Column3"},
		Data: [][]float64{
			{1.1, 2.2, 3.3},
			{4.4, 5.5, 0},
			{7.7, 8.8, 9.9},
		},
	}

	result, err := staticFireParser.ParseSv(rawSvText, '\t')

	if err != nil {
		t.Errorf("ParseSv() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseSv() = %v, want %v", result, expected)
	}
}

func TestIncompleteSvShouldFillWithZero(t *testing.T) {
	rawSvText := `Column1	Column2	Column3
1.1	2.2	3.3
4.4	5.5	6.6
7.7	8.8`

	expected := staticFireParser.ParsedSv{
		ColumnCount: 3,
		ColumnNames: []string{"Column1", "Column2", "Column3"},
		Data: [][]float64{
			{1.1, 2.2, 3.3},
			{4.4, 5.5, 6.6},
			{7.7, 8.8, 0},
		},
	}

	result, err := staticFireParser.ParseSv(rawSvText, '\t')

	if err != nil {
		t.Errorf("ParseSv() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseSv() = %v, want %v", result, expected)
	}
}

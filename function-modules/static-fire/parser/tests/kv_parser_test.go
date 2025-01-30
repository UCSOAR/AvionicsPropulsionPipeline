package tests

import (
	"reflect"
	"testing"

	parser "example.com/static-fire/parser"
)

func TestValidHeaderShouldParseCorrectly(t *testing.T) {
	rawHeaderText := `Key1 Value1
Key2 Value2 Value3`
	expected := parser.ParsedKvHeader{
		Kv: map[string][]string{
			"Key1": {"Value1"},
			"Key2": {"Value2", "Value3"},
		},
	}
	result, err := parser.ParseKv(rawHeaderText)

	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, want %v", result, expected)
	}
}

func TestInvalidHeaderLineShouldError(t *testing.T) {
	rawHeaderText := `InvalidLine`

	_, err := parser.ParseKv(rawHeaderText)

	if err == nil {
		t.Errorf("Expected invalid header line error")
	}
}

func TestParseDuplicateKeyShouldError(t *testing.T) {
	rawHeaderText := `Key1 Value1
Key1 Value2`

	_, err := parser.ParseKv(rawHeaderText)

	if err == nil {
		t.Errorf("Expected duplicate key error")
	}
}

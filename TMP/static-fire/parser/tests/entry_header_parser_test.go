package tests

import (
	"reflect"
	"testing"

	parser "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser"
)

func TestValidEntryHeaderTextParsesCorrectly(t *testing.T) {
	rawHeaderText := `Writer_Version 2
Reader_Version 2
Separator Tab
Decimal_Separator .
Multi_Headings Yes
X_Columns Multi
Time_Pref Absolute
Operator TestOperator
Date 2025/01/18
Time 10:45:47.0352557312499836422`

	multiHeadings, err := parser.ParseMultiHeadingsValue("Yes")

	if err != nil {
		t.Errorf("ParseMultiHeadingsValue() error = %v", err)
		return
	}

	xColumns, err := parser.ParseXColumnsValue("Multi")

	if err != nil {
		t.Errorf("ParseXColumnsValue() error = %v", err)
		return
	}

	expected := parser.ParsedEntryHeader{
		Seperator:     '\t',
		MultiHeadings: multiHeadings,
		XColumns:      xColumns,
		Operator:      "TestOperator",
		Date:          "2025/01/18",
		Time:          "10:45:47.0352557312499836422",
	}

	result, err := parser.ParseEntryHeader(rawHeaderText)

	if err != nil {
		t.Errorf("ParseEntryHeader() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseEntryHeader() = %v, want %v", result, expected)
	}
}

func TestIncompleteEntryHeaderShouldError(t *testing.T) {
	rawHeaderText := `Writer_Version 2
Reader_Version 2
Separator Tab
Decimal_Separator .`

	_, err := parser.ParseEntryHeader(rawHeaderText)

	if err == nil {
		t.Errorf("Expected error for incomplete header")
	}
}

func TestIncorrectWriterVersionShouldError(t *testing.T) {
	rawHeaderText := `Writer_Version 1
Reader_Version 2
Separator Tab
Decimal_Separator .
Multi_Headings Yes
X_Columns Multi
Time_Pref Absolute
Operator TestOperator
Date 2025/01/18
Time 10:45:47.0352557312499836422`

	_, err := parser.ParseEntryHeader(rawHeaderText)

	if err == nil {
		t.Errorf("Expected error for incorrect Writer_Version")
	}
}

func TestIncorrectReaderVersionShouldError(t *testing.T) {
	rawHeaderText := `Writer_Version 2
Reader_Version 1
Separator Tab
Decimal_Separator .
Multi_Headings Yes
X_Columns Multi
Time_Pref Absolute
Operator TestOperator
Date 2025/01/18
Time 10:45:47.0352557312499836422`

	_, err := parser.ParseEntryHeader(rawHeaderText)

	if err == nil {
		t.Errorf("Expected error for incorrect Reader_Version")
	}
}

func TestIncorrectTimePrefShouldError(t *testing.T) {
	rawHeaderText := `Writer_Version 2
Reader_Version 2
Separator Tab
Decimal_Separator .
Multi_Headings No
X_Columns Multi
Time_Pref Sometime
Operator TestOperator
Date 2025/01/18
Time 10:45:47.0352557312499836422`

	_, err := parser.ParseEntryHeader(rawHeaderText)

	if err == nil {
		t.Errorf("Expected error for incorrect Time_Pref")
	}
}

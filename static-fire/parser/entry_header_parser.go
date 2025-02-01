package parser

import "fmt"

type MultiHeadingsValue uint8

const (
	MultiHeadingsUnknown MultiHeadingsValue = 0
	MultiHeadingsYes     MultiHeadingsValue = 1
	MultiHeadingsNo      MultiHeadingsValue = 2
)

func ParseMultiHeadingsValue(multiHeadingsText string) (MultiHeadingsValue, error) {
	switch multiHeadingsText {
	case "Yes":
		return MultiHeadingsYes, nil
	case "No":
		return MultiHeadingsNo, nil
	default:
		return MultiHeadingsUnknown, fmt.Errorf("Invalid multi headings value: %s", multiHeadingsText)
	}
}

func (m MultiHeadingsValue) String() string {
	switch m {
	case 1:
		return "Yes"
	case 2:
		return "No"
	default:
		return "Unknown"
	}
}

type XColumnsValue uint8

const (
	XColumnsUnknown XColumnsValue = 0
	XColumnsOne     XColumnsValue = 1
	XColumnsMulti   XColumnsValue = 2
)

func ParseXColumnsValue(xColumnsText string) (XColumnsValue, error) {
	switch xColumnsText {
	case "One":
		return XColumnsOne, nil
	case "Multi":
		return XColumnsMulti, nil
	default:
		return XColumnsUnknown, fmt.Errorf("Invalid X columns value: %s", xColumnsText)
	}
}

func (x XColumnsValue) String() string {
	switch x {
	case 1:
		return "One"
	case 2:
		return "Multi"
	default:
		return "Unknown"
	}
}

type FieldSeperator rune

const (
	FieldSeperatorUnknown FieldSeperator = 0
	FieldSeperatorTab     FieldSeperator = '\t'
	FieldSeperatorSpace   FieldSeperator = ' '
)

func ParseFieldSeperator(seperatorText string) (FieldSeperator, error) {
	switch seperatorText {
	case "Tab":
		return FieldSeperatorTab, nil
	case "Space":
		return FieldSeperatorTab, nil
	default:
		return FieldSeperatorUnknown, fmt.Errorf("Invalid seperator: %s", seperatorText)
	}
}

type ParsedEntryHeader struct {
	Seperator     FieldSeperator     `json:"seperator"`
	MultiHeadings MultiHeadingsValue `json:"multiHeadings"`
	XColumns      XColumnsValue      `json:"xColumns"`
	Operator      string             `json:"operator"`
	Date          string             `json:"date"`
	Time          string             `json:"time"`
}

// Parses only the text that contains the entry header section.
// Returns a struct representing the parsed entry header.
func ParseEntryHeader(rawHeaderText string) (ParsedEntryHeader, error) {
	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return ParsedEntryHeader{}, err
	}

	// Ensure all required keys are present
	requiredKeys := []string{
		"Writer_Version",
		"Reader_Version",
		"Separator",
		"Decimal_Separator",
		"Multi_Headings",
		"X_Columns",
		"Time_Pref",
		"Operator",
		"Date",
		"Time",
	}

	for _, key := range requiredKeys {
		if _, ok := parsedHeader.Kv[key]; !ok {
			return ParsedEntryHeader{}, fmt.Errorf("Missing key: %s", key)
		}
	}

	// Handle assertions
	if parsedHeader.Kv["Writer_Version"][0] != AssertedWriterVersion {
		return ParsedEntryHeader{}, fmt.Errorf("Writer_Version is not %s", AssertedWriterVersion)
	}

	if parsedHeader.Kv["Reader_Version"][0] != AssertedReaderVersion {
		return ParsedEntryHeader{}, fmt.Errorf("Reader_Version is not %s", AssertedReaderVersion)
	}

	if parsedHeader.Kv["Decimal_Separator"][0] != AssertedDecimalSeparator {
		return ParsedEntryHeader{}, fmt.Errorf("Decimal_Separator is not %s", AssertedDecimalSeparator)
	}

	if parsedHeader.Kv["Time_Pref"][0] != AssertedTimePreferance {
		return ParsedEntryHeader{}, fmt.Errorf("Time_Preference is not %s", AssertedTimePreferance)
	}

	// Create entry header structure
	seperator, err := ParseFieldSeperator(parsedHeader.Kv["Separator"][0])

	if err != nil {
		return ParsedEntryHeader{}, err
	}

	multiHeadings, err := ParseMultiHeadingsValue(parsedHeader.Kv["Multi_Headings"][0])

	if err != nil {
		return ParsedEntryHeader{}, err
	}

	xColumns, err := ParseXColumnsValue(parsedHeader.Kv["X_Columns"][0])

	if err != nil {
		return ParsedEntryHeader{}, err
	}

	entryHeader := ParsedEntryHeader{
		Seperator:     seperator,
		MultiHeadings: multiHeadings,
		XColumns:      xColumns,
		Operator:      parsedHeader.Kv["Operator"][0],
		Date:          parsedHeader.Kv["Date"][0],
		Time:          parsedHeader.Kv["Time"][0],
	}

	return entryHeader, nil
}

package staticfire

import "fmt"

var requiredEntryHeaderKeys = [...]string{
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

// Parses only the text that contains the entry header section.
// Returns a struct representing the parsed entry header.
func ParseEntryHeader(rawHeaderText string) (ParsedEntryHeader, error) {
	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return ParsedEntryHeader{}, err
	}

	// Ensure all required keys are present
	for _, key := range requiredEntryHeaderKeys {
		if _, ok := parsedHeader.Kv[key]; !ok {
			return ParsedEntryHeader{}, fmt.Errorf("missing key: %s", key)
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

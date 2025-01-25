package staticFireParser

import "fmt"

type ParsedEntryHeader struct {
	Seperator      rune
	IsMultiHeading bool
	HasOneXColumn  bool
	Operator       string
	Date           string
	Time           string
}

// Helper function for simply converting the word for a seperator into a rune
func seperatorFromText(seperatorText string) rune {
	if seperatorText == "Tab" {
		return '\t'
	}

	return ' '
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
	entryHeader := ParsedEntryHeader{
		Seperator:      seperatorFromText(parsedHeader.Kv["Separator"][0]),
		IsMultiHeading: parsedHeader.Kv["Multi_Headings"][0] == "Yes",
		HasOneXColumn:  parsedHeader.Kv["X_Columns"][0] == "One",
		Operator:       parsedHeader.Kv["Operator"][0],
		Date:           parsedHeader.Kv["Date"][0],
		Time:           parsedHeader.Kv["Time"][0],
	}

	return entryHeader, nil
}

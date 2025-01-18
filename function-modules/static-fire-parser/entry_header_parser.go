package staticFireParser

import "fmt"

type EntryHeader struct {
	Seperator        rune
	DecimalSeparator rune
	HasOneXColumn    bool
	Operator         string
	Date             string
	Time             string
}

func seperatorFromText(seperatorText string) rune {
	if seperatorText == "Tab" {
		return '\t'
	}

	return ' '
}

func ParseEntryHeader(rawHeaderText string) (EntryHeader, error) {
	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return EntryHeader{}, err
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
			return EntryHeader{}, fmt.Errorf("Missing key: %s", key)
		}
	}

	// Handle assertions
	if parsedHeader.Kv["Writer_Version"][0] != AssertedWriterVersion {
		return EntryHeader{}, fmt.Errorf("Writer_Version is not %s", AssertedWriterVersion)
	}

	if parsedHeader.Kv["Reader_Version"][0] != AssertedReaderVersion {
		return EntryHeader{}, fmt.Errorf("Reader_Version is not %s", AssertedReaderVersion)
	}

	if parsedHeader.Kv["Multi_Headings"][0] != AssertedMultiHeadings {
		return EntryHeader{}, fmt.Errorf("Multi_Headings is not %s", AssertedMultiHeadings)
	}

	if parsedHeader.Kv["Time_Pref"][0] != AssertedTimePreferance {
		return EntryHeader{}, fmt.Errorf("Time_Preference is not %s", AssertedTimePreferance)
	}

	// Create entry header structure
	entryHeader := EntryHeader{
		Seperator:        seperatorFromText(parsedHeader.Kv["Separator"][0]),
		DecimalSeparator: rune(parsedHeader.Kv["Decimal_Separator"][0][0]),
		HasOneXColumn:    parsedHeader.Kv["X_Columns"][0] == "One",
		Operator:         parsedHeader.Kv["Operator"][0],
		Date:             parsedHeader.Kv["Date"][0],
		Time:             parsedHeader.Kv["Time"][0],
	}

	return entryHeader, nil
}

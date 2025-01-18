package staticFireParser

import (
	"bufio"
	"fmt"
	"strings"
)

type ParsedLvm struct {
	EntryHeader    ParsedEntryHeader
	ChannelHeader  ParsedChannelHeader
	CsvColumnCount uint64
	CsvColumns     []string
	CsvData        [][]float64
}

func ParseMainLvm(rawLvmText string) (ParsedLvm, error) {
	reader := strings.NewReader(rawLvmText)
	scanner := bufio.NewScanner(reader)

	// Assert first line
	if line := scanner.Text(); line != AssertedFirstLine {
		return ParsedLvm{}, fmt.Errorf("First line does not match expected: %s", line)
	}

	// Read forward to the end of the header
	// ...

	lvm := ParsedLvm{}

	return lvm, nil
}

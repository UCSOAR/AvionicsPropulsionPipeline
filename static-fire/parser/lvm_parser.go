package parser

import (
	"bufio"
	"fmt"
	"strings"
)

type HeaderMetadata struct {
	EntryHeader   ParsedEntryHeader
	ChannelHeader ParsedChannelHeader
}

// Represents a parsed LabVIEW Measurement File (LVM).
type ParsedLvm struct {
	Headers HeaderMetadata
	SvData  ParsedSv // Seperated value data
}

// Move a scanner forward until the end of the header.
// Returns the scanned text.
func readUntilEOH(scanner *bufio.Scanner) string {
	var headerBuilder strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if line == AssertedEndOfHeader {
			break
		}

		headerBuilder.WriteString(line)
		headerBuilder.WriteRune('\n')
	}

	return headerBuilder.String()
}

// Move a scanner forward until the end of the file.
// Returns the scanned text.
func readUntilEOF(scanner *bufio.Scanner) string {
	var dataBuilder strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		dataBuilder.WriteString(line)
		dataBuilder.WriteRune('\n')
	}

	return dataBuilder.String()
}

// Parses a raw LVM file into a struct representing the parsed data.
func ParseLvm(rawLvmText string) (ParsedLvm, error) {
	reader := strings.NewReader(rawLvmText)
	scanner := bufio.NewScanner(reader)

	// Assert first line
	if !scanner.Scan() {
		return ParsedLvm{}, fmt.Errorf("Failed to read first line")
	}

	if line := strings.TrimSpace(scanner.Text()); line != AssertedFirstLine {
		return ParsedLvm{}, fmt.Errorf("First line %s does not match expected %s", line, AssertedFirstLine)
	}

	// Read forward to the end of the first header (entry header)
	rawEntryHeaderText := readUntilEOH(scanner)
	entryHeader, err := ParseEntryHeader(rawEntryHeaderText)

	if err != nil {
		return ParsedLvm{}, err
	}

	// Read forward to the end of the second header (channel header)
	rawChannelHeaderText := readUntilEOH(scanner)
	channelHeader, err := ParseChannelHeader(rawChannelHeaderText)

	if err != nil {
		return ParsedLvm{}, err
	}

	// Read the rest of the file which is expected to be the SV data
	rawSvText := readUntilEOF(scanner)
	svData, err := ParseSv(rawSvText, rune(entryHeader.Seperator))

	if err != nil {
		return ParsedLvm{}, err
	}

	// Construct parsed LVM
	lvm := ParsedLvm{
		Headers: HeaderMetadata{
			EntryHeader:   entryHeader,
			ChannelHeader: channelHeader,
		},
		SvData: svData,
	}

	return lvm, nil
}

package staticFireParser

import (
	"bufio"
	"fmt"
	"strings"
)

type ParsedLvm struct {
	EntryHeader   ParsedEntryHeader
	ChannelHeader ParsedChannelHeader
	SvData        ParsedSv
}

func readUntilEOH(scanner *bufio.Scanner) string {
	var headerBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if line == AssertedReaderVersion {
			break
		}

		headerBuilder.WriteString(line)
	}

	return headerBuilder.String()
}

func readUntilEOF(scanner *bufio.Scanner) string {
	var dataBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		dataBuilder.WriteString(line)
	}

	return dataBuilder.String()
}

func ParseMainLvm(rawLvmText string) (ParsedLvm, error) {
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
	svData, err := ParseSv(rawSvText, entryHeader.Seperator)

	if err != nil {
		return ParsedLvm{}, err
	}

	// Construct parsed LVM
	lvm := ParsedLvm{
		EntryHeader:   entryHeader,
		ChannelHeader: channelHeader,
		SvData:        svData,
	}

	return lvm, nil
}

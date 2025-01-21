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

	for {
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

func ParseMainLvm(rawLvmText string) (ParsedLvm, error) {
	reader := strings.NewReader(rawLvmText)
	scanner := bufio.NewScanner(reader)

	// Assert first line
	if line := scanner.Text(); line != AssertedFirstLine {
		return ParsedLvm{}, fmt.Errorf("First line does not match expected: %s", line)
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

	// Construct parsed LVM
	lvm := ParsedLvm{
		EntryHeader:   entryHeader,
		ChannelHeader: channelHeader,
	}

	return lvm, nil
}

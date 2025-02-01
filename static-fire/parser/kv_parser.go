package parser

import (
	"bufio"
	"errors"
	"strings"
)

// Represents a parsed key-value header.
type ParsedKvHeader struct {
	Kv map[string]([]string) // Map of header keys to a list of values
}

// Parses a generic key-value header.
func ParseKv(rawHeaderText string) (ParsedKvHeader, error) {
	reader := strings.NewReader(rawHeaderText)
	scanner := bufio.NewScanner(reader)

	// Initialize the map
	header := ParsedKvHeader{
		Kv: make(map[string]([]string)),
	}

	// Read the header text line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split line at first tab or space
		parts := strings.Fields(line)

		if len(parts) < 2 {
			return ParsedKvHeader{}, errors.New("Invalid header line: " + line)
		}

		// Add key-value pair to the map
		key := parts[0]
		rest := parts[1:]

		// Key cannot already exist
		if header.Kv[key] != nil {
			return ParsedKvHeader{}, errors.New("Duplicate key: " + key)
		}

		header.Kv[key] = rest
	}

	if err := scanner.Err(); err != nil {
		return ParsedKvHeader{}, err
	}

	return header, nil
}

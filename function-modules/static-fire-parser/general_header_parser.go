package staticFireParser

import (
	"bufio"
	"errors"
	"strings"
)

type ParsedHeader struct {
	// Map of header keys to a list of values
	Kv map[string]([]string)
}

func ParseKv(rawHeaderText string) (ParsedHeader, error) {
	reader := strings.NewReader(rawHeaderText)
	scanner := bufio.NewScanner(reader)

	// Initialize the map
	header := ParsedHeader{
		Kv: make(map[string]([]string)),
	}

	// Read the header text line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split line at first tab or space
		parts := strings.Fields(line)

		if len(parts) < 2 {
			return ParsedHeader{}, errors.New("Invalid header line: " + line)
		}

		// Add key-value pair to the map
		key := parts[0]
		rest := parts[1:]

		// Key cannot already exist
		if header.Kv[key] != nil {
			return ParsedHeader{}, errors.New("Duplicate key: " + key)
		}

		header.Kv[key] = rest
	}

	if err := scanner.Err(); err != nil {
		return ParsedHeader{}, err
	}

	return header, nil
}

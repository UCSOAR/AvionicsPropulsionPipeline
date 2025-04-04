package staticfire

import (
	"fmt"
	"strconv"
)

// Parses only the text that contains the channel header section.
// Returns a struct representing the parsed channel header.
func ParseChannelHeader(rawHeaderText string) (ParsedChannelHeader, error) {
	requiredChannelHeaderKeys := [...]string{
		"Channels",
		"Samples",
		"Date",
		"Time",
		"Y_Unit_Label",
		"X_Dimension",
		"X0",
		"Delta_X",
	}

	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return ParsedChannelHeader{}, err
	}

	// Ensure all required keys are present
	for _, key := range requiredChannelHeaderKeys {
		if _, ok := parsedHeader.Kv[key]; !ok {
			return ParsedChannelHeader{}, fmt.Errorf("missing key: %s", key)
		}
	}

	// Attempt to parse channel count into an unsigned integer
	channelCount, err := strconv.Atoi(parsedHeader.Kv["Channels"][0])

	if err != nil {
		return ParsedChannelHeader{}, fmt.Errorf("failed to parse Channels: %s", parsedHeader.Kv["Channels"][0])
	}

	// Ensure all arrays have the length of the channel count
	for i := 1; i < len(requiredChannelHeaderKeys); i++ {
		if len(parsedHeader.Kv[requiredChannelHeaderKeys[i]]) != channelCount {
			return ParsedChannelHeader{}, fmt.Errorf("length of %s does not match channel count", requiredChannelHeaderKeys[i])
		}
	}

	// Parse samples
	samples := make([]int, channelCount)

	for i := range channelCount {
		sampleCount, err := strconv.Atoi(parsedHeader.Kv["Samples"][i])

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("failed to parse Samples: %s", parsedHeader.Kv["Samples"][i])
		}

		samples[i] = sampleCount
	}

	// Get dates and times
	dates := parsedHeader.Kv["Date"]
	times := parsedHeader.Kv["Time"]

	// Get Y unit labels and X dimensions
	yUnitLabels := parsedHeader.Kv["Y_Unit_Label"]
	xDimensions := parsedHeader.Kv["X_Dimension"]

	// Parse initial Xs
	initialXs := make([]float64, channelCount)

	for i := range channelCount {
		initialX, err := strconv.ParseFloat(parsedHeader.Kv["X0"][i], 64)

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("failed to parse X0: %s", parsedHeader.Kv["X0"][i])
		}

		initialXs[i] = initialX
	}

	// Parse delta Xs
	deltaXs := make([]float64, channelCount)

	for i := range channelCount {
		deltaX, err := strconv.ParseFloat(parsedHeader.Kv["Delta_X"][i], 64)

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("failed to parse Delta_X: %s", parsedHeader.Kv["Delta_X"][i])
		}

		deltaXs[i] = deltaX
	}

	// Create and return the parsed channel header
	channelHeader := ParsedChannelHeader{
		ChannelCount: channelCount,
		Samples:      samples,
		Dates:        dates,
		Times:        times,
		YUnitLabels:  yUnitLabels,
		XDimensions:  xDimensions,
		InitialXs:    initialXs,
		DeltaXs:      deltaXs,
	}

	return channelHeader, nil
}

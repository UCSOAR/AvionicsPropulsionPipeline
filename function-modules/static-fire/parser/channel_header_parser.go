package parser

import (
	"fmt"
	"strconv"
)

// Represents a parsed LVM channel header.
// The data stored is relevant to the purposes of this project.
// It is guaranteed that all arrays will have the same length as `ChannelCount`.
type ParsedChannelHeader struct {
	ChannelCount uint64    `json:"channelCount"`
	Samples      []uint64  `json:"samples"`
	Dates        []string  `json:"dates"`
	Times        []string  `json:"times"`
	YUnitLabels  []string  `json:"yUnitLabels"`
	XDimensions  []string  `json:"xDimensions"`
	InitialXs    []float64 `json:"initialXs"`
	DeltaXs      []float64 `json:"deltaXs"`
}

// Parses only the text that contains the channel header section.
// Returns a struct representing the parsed channel header.
func ParseChannelHeader(rawHeaderText string) (ParsedChannelHeader, error) {
	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return ParsedChannelHeader{}, err
	}

	// Ensure all required keys are present
	requiredKeys := []string{
		"Channels",
		"Samples",
		"Date",
		"Time",
		"Y_Unit_Label",
		"X_Dimension",
		"X0",
		"Delta_X",
	}

	for _, key := range requiredKeys {
		if _, ok := parsedHeader.Kv[key]; !ok {
			return ParsedChannelHeader{}, fmt.Errorf("Missing key: %s", key)
		}
	}

	// Attempt to parse channel count into an unsigned integer
	channelCount, err := strconv.ParseUint(parsedHeader.Kv["Channels"][0], 10, 64)

	if err != nil {
		return ParsedChannelHeader{}, fmt.Errorf("Failed to parse Channels: %s", parsedHeader.Kv["Channels"][0])
	}

	// Ensure all arrays have the length of the channel count
	for i := 1; i < len(requiredKeys); i++ {
		if uint64(len(parsedHeader.Kv[requiredKeys[i]])) != channelCount {
			return ParsedChannelHeader{}, fmt.Errorf("Length of %s does not match channel count", requiredKeys[i])
		}
	}

	// Parse samples
	samples := make([]uint64, channelCount)

	for i := uint64(0); i < channelCount; i++ {
		sampleCount, err := strconv.ParseUint(parsedHeader.Kv["Samples"][i], 10, 64)

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("Failed to parse Samples: %s", parsedHeader.Kv["Samples"][i])
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

	for i := uint64(0); i < channelCount; i++ {
		initialX, err := strconv.ParseFloat(parsedHeader.Kv["X0"][i], 64)

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("Failed to parse X0: %s", parsedHeader.Kv["X0"][i])
		}

		initialXs[i] = initialX
	}

	// Parse delta Xs
	deltaXs := make([]float64, channelCount)

	for i := uint64(0); i < channelCount; i++ {
		deltaX, err := strconv.ParseFloat(parsedHeader.Kv["Delta_X"][i], 64)

		if err != nil {
			return ParsedChannelHeader{}, fmt.Errorf("Failed to parse Delta_X: %s", parsedHeader.Kv["Delta_X"][i])
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

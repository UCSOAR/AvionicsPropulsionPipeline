package staticFireParser

import (
	"fmt"
	"strconv"
)

type ChannelHeader struct {
	ChannelCount uint64
	Samples      []uint64
	Dates        []string
	Times        []string
	YunitLabels  []string
	Xdimensions  []string
	InitialXs    []float64
	DeltaXs      []float64
}

func ParseChannelHeader(rawHeaderText string) (ChannelHeader, error) {
	parsedHeader, err := ParseKv(rawHeaderText)

	if err != nil {
		return ChannelHeader{}, err
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
			return ChannelHeader{}, fmt.Errorf("Missing key: %s", key)
		}
	}

	channelCount, err := strconv.ParseUint(parsedHeader.Kv["Channels"][0], 10, 64)

	if err != nil {
		return ChannelHeader{}, fmt.Errorf("Failed to parse Channels: %s", parsedHeader.Kv["Channels"][0])
	}

	// Ensure all arrays have the length of the channel count
	for i := 1; i < len(requiredKeys); i++ {
		if uint64(len(parsedHeader.Kv[requiredKeys[i]])) != channelCount {
			return ChannelHeader{}, fmt.Errorf("Length of %s does not match channel count", requiredKeys[i])
		}
	}

	// Parse samples
	samples := make([]uint64, channelCount)

	for i := uint64(0); i < channelCount; i++ {
		sampleCount, err := strconv.ParseUint(parsedHeader.Kv["Samples"][i], 10, 64)

		if err != nil {
			return ChannelHeader{}, fmt.Errorf("Failed to parse Samples: %s", parsedHeader.Kv["Samples"][i])
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
			return ChannelHeader{}, fmt.Errorf("Failed to parse X0: %s", parsedHeader.Kv["X0"][i])
		}

		initialXs[i] = initialX
	}

	// Parse delta Xs
	deltaXs := make([]float64, channelCount)

	for i := uint64(0); i < channelCount; i++ {
		deltaX, err := strconv.ParseFloat(parsedHeader.Kv["Delta_X"][i], 64)

		if err != nil {
			return ChannelHeader{}, fmt.Errorf("Failed to parse Delta_X: %s", parsedHeader.Kv["Delta_X"][i])
		}

		deltaXs[i] = deltaX
	}

	channelHeader := ChannelHeader{
		ChannelCount: channelCount,
		Samples:      samples,
		Dates:        dates,
		Times:        times,
		YunitLabels:  yUnitLabels,
		Xdimensions:  xDimensions,
		InitialXs:    initialXs,
		DeltaXs:      deltaXs,
	}

	return channelHeader, nil
}

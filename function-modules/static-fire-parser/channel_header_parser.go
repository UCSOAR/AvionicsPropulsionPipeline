package staticFireParser

type ChannelHeader struct {
	channelCount uint
	samples      []uint
	dates        []string
	times        []string
	yUnitLabels  []string
	xDimensions  []string
	initialXs    []float64
	deltaXs      []float64
}

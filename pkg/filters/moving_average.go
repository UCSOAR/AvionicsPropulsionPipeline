package filters

// MovingAvgFilter applies a simple moving average filter to smooth yRows.
func MovingAvgFilter(xRows []float64, yRows []float64, windowSize int) []float64 {
	if len(xRows) != len(yRows) || windowSize <= 0 || len(yRows) < windowSize {
		return nil // Ensure input arrays have the same length and valid window size
	}

	smoothed := make([]float64, len(yRows))

	// Apply the moving average for each point
	for i := range yRows {
		sum := 0.0
		count := 0

		// Calculate the sum of the points within the window
		for j := max(0, i-(windowSize/2)); j < min(len(yRows), i+(windowSize/2)+1); j++ {
			sum += yRows[j]
			count++
		}

		smoothed[i] = sum / float64(count)
	}

	return smoothed
}

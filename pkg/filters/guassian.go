package filters

import "math"

// GaussianFilter applies a Gaussian filter to smooth the data.
// Now accepts sigma as a parameter so users can control smoothing.
func GaussianFilter(xRows []float64, yRows []float64, sigma float64) []float64 {
	if len(yRows) == 0 {
		return nil // Ensure there is data to process
	}

	if sigma <= 0 {
		sigma = 1.0 // fallback to default if invalid sigma is provided
	}

	windowSize := int(6*sigma) + 1 // window size based on sigma
	smoothed := make([]float64, len(yRows))

	// Generate the Gaussian kernel
	kernel := generateGaussianKernel(windowSize, sigma)

	// Apply the Gaussian filter
	for i := range yRows {
		sum := 0.0
		weightedSum := 0.0

		for j := -windowSize / 2; j <= windowSize/2; j++ {
			idx := i + j

			if idx >= 0 && idx < len(yRows) {
				weight := kernel[j+(windowSize/2)]
				sum += weight * yRows[idx]
				weightedSum += weight
			}
		}

		smoothed[i] = sum / weightedSum
	}

	return smoothed
}

// generateGaussianKernel creates a Gaussian kernel with a given sigma.
func generateGaussianKernel(size int, sigma float64) []float64 {
	kernel := make([]float64, size)
	sum := 0.0

	for i := range size {
		x := float64(i - size/2)
		kernel[i] = math.Exp(-0.5 * (x * x) / (sigma * sigma))
		sum += kernel[i]
	}
	// Normalize the kernel so it sums to 1
	for i := range size {
		kernel[i] /= sum
	}

	return kernel
}

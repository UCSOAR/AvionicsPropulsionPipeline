package filters

import "math"

// SomeFilter applies a simple moving average filter to smooth yRows.
func MovingAvgFilter(xRows []float64, yRows []float64) []float64 {
    windowSize := 3 // Defines the smoothing window size (this can be adjusted)
    
    if len(xRows) != len(yRows) || windowSize <= 0 || len(yRows) < windowSize {
        return nil // Ensure input arrays have the same length and valid window size
    }

    smoothed := make([]float64, len(yRows))

    // Apply the moving average for each point
    for i := 0; i < len(yRows); i++ {
        sum := 0.0
        count := 0
        // Calculate the sum of the points within the window
        for j := max(0, i-windowSize/2); j < min(len(yRows), i+windowSize/2+1); j++ {
            sum += yRows[j]
            count++
        }
        smoothed[i] = sum / float64(count)
    }

    return smoothed
}

// Helper functions to handle boundary conditions
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


// Gaussian filter applies a Gaussian filter to smooth the data.
func GaussianFilter(xRows []float64, yRows []float64) []float64 {
    if len(yRows) == 0 {
        return nil // Ensure there is data to process
    }

    sigma := 1.0 // Set the standard deviation for the Gaussian filter (can be adjusted)
    windowSize := int(6*sigma) + 1 // window size based on sigma
    smoothed := make([]float64, len(yRows))

    // Generate the Gaussian kernel
    kernel := generateGaussianKernel(windowSize, sigma)

    // Apply the Gaussian filter
    for i := 0; i < len(yRows); i++ {
        sum := 0.0
        weightSum := 0.0
        for j := -windowSize / 2; j <= windowSize/2; j++ {
            idx := i + j
            if idx >= 0 && idx < len(yRows) {
                weight := kernel[j+windowSize/2]
                sum += weight * yRows[idx]
                weightSum += weight
            }
        }
        smoothed[i] = sum / weightSum
    }

    return smoothed
}

// generateGaussianKernel creates a Gaussian kernel with a given sigma.
func generateGaussianKernel(size int, sigma float64) []float64 {
    kernel := make([]float64, size)
    sum := 0.0
    for i := 0; i < size; i++ {
        x := float64(i - size/2)
        kernel[i] = math.Exp(-0.5 * (x * x) / (sigma * sigma))
        sum += kernel[i]
    }
    // Normalize the kernel so it sums to 1
    for i := 0; i < size; i++ {
        kernel[i] /= sum
    }
    return kernel
}


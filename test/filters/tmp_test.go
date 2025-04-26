package tests

import (
	"reflect"
	"soarpipeline/pkg/filters"
	"testing"
)

// Test for Moving Average Filter
func TestMovingAvgFilter(t *testing.T) {
	xRows := []float64{1, 2, 3, 4, 5}
	yRows := []float64{10, 20, 30, 40, 50}

	expected := []float64{15, 20, 30, 40, 45} // Expected moving average with windowSize 3

	result := filters.MovingAvgFilter(xRows, yRows)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MovingAvgFilter() = %v, expected %v", result, expected)
	}
}

// Test for Gaussian Filter
func TestGaussianFilter(t *testing.T) {
	xRows := []float64{1, 2, 3, 4, 5}
	yRows := []float64{10, 20, 30, 40, 50}

	sigma := 1.0 
	result := filters.GaussianFilter(xRows, yRows, sigma) // pass it in!

	// Ensure the output length matches the input length
	if len(result) != len(yRows) {
		t.Errorf("GaussianFilter() returned incorrect length: got %d, expected %d", len(result), len(yRows))
	}

	// Check if the values are smoothed (not the same as original)
	if reflect.DeepEqual(result, yRows) {
		t.Errorf("GaussianFilter() did not modify the input data")
	}
}

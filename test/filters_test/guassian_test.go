package filters_test

import (
	"reflect"
	"soarpipeline/pkg/filters"
	"testing"
)

// Test for Gaussian Filter.
func TestGaussianFilter(t *testing.T) {
	xRows := []float64{1, 2, 3, 4, 5}
	yRows := []float64{10, 20, 30, 40, 50}

	sigma := 1.0
	result := filters.GaussianFilter(xRows, yRows, sigma)

	// Ensure the output length matches the input length.
	if len(result) != len(yRows) {
		t.Errorf("GaussianFilter() returned incorrect length: got %d, expected %d", len(result), len(yRows))
	}

	// Check if the values are smoothed (not the same as original).
	if reflect.DeepEqual(result, yRows) {
		t.Errorf("GaussianFilter() did not modify the input data")
	}
}

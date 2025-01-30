package parser

import (
	caching "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching"
)

func (lvm *ParsedLvm) ToCacheTree() (caching.CacheTree, error) {
	// Just return this tree for now
	tree := caching.CacheTree{
		PreviewMetadata: caching.PreviewMetadata{
			ResultTimestamp: caching.TimestampMetadata{
				Date: "2025-01-01",
				Time: "12:00:00",
			},
			ProcessedTimestamp: caching.TimestampMetadata{
				Date: "2025-01-01",
				Time: "12:00:00",
			},
			Operator:     "Test Operator",
			XColumnNames: []string{"x-ax", "x-ay", "x-az"},
			YColumnNames: []string{"ax", "ay", "az"},
		},
		XColumnNodes: []caching.XColumnNode{
			{
				Dimension: "Time",
				Rows:      []float64{0.0, 0.1, 0.2, 0.3, 0.4},
			},
		},
		YColumnNodes: []caching.YColumnNode{
			{
				Samples:   5,
				Date:      "2025-01-01",
				UnitLabel: "m/s^2",
				Rows:      []float64{0.0, 0.1, 0.2, 0.3, 0.4},
			},
		},
	}

	return tree, nil
}

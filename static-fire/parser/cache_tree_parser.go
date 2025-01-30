package parser

import (
	caching "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching"
)

func (lvm *ParsedLvm) ToCacheTree() (caching.CacheTree, error) {
	return caching.CacheTree{}, nil
}

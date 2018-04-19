package text

import (
	"github.com/sckelemen/indexer/core"
)

type TextIndexer struct {
}

func (ti TextIndexer) IndexerType() core.IndexerType {
	return TextIndexerType
}

const (
	TextIndexerType core.IndexerType = "text_indexer"
)

func NewIndexer() TextIndexer {
	return TextIndexer{}
}

package symbol

import (
	"github.com/sckelemen/indexer/core"
)

type SymbolIndexer struct {
}

func (si SymbolIndexer) IndexerType() core.IndexerType {
	return SymbolIndexerType
}

const (
	SymbolIndexerType core.IndexerType = "roslyn_indexer"
)

func NewIndexer() SymbolIndexer {
	return SymbolIndexer{}
}

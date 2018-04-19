package regex

import (
	"github.com/sckelemen/indexer/core"
)

type RegexIndexer struct {
}

func (ri RegexIndexer) IndexerType() core.IndexerType {
	return RegexIndexerType
}

const (
	RegexIndexerType core.IndexerType = "regex_indexer"
)

func NewIndexer() RegexIndexer {
	return RegexIndexer{}
}

package core

func NewIndexer() Indexer {
	return Indexer{}
}

type Indexer struct {
	indexers []IIndexer
}

func (i Indexer) AddIndexer(indexer IIndexer) bool {
	i.indexers = append(i.indexers, indexer)
	return true
}

package core

type IIndexer interface {
	IndexerType() IndexerType
}

type IndexerType string

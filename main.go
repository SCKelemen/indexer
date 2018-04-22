package main

import (
	"github.com/SCKelemen/indexer/regex"
	"github.com/SCKelemen/indexer/symbol"
	"github.com/SCKelemen/indexer/text"
	"github.com/sckelemen/indexer/core"
)

func main() {
	indexer := core.NewIndexer()
	ti := text.NewIndexer()
	ri := regex.NewIndexer()
	si := symbol.NewIndexer()
	indexer.AddIndexer(ti)
	indexer.AddIndexer(ri)
	indexer.AddIndexer(si)
}

// add trigrams

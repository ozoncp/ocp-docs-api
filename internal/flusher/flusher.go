package flusher

import (
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/repo"
	"github.com/ocp-docs-api/internal/utils"
)

type Flusher interface {
	Flush(docs []document.Document) []document.Document
}

type flusher struct {
	repo repo.Repo
	chunkSize int
}

func New(docsRepo repo.Repo, chunkSize int) Flusher {
	return &flusher {
		    repo : docsRepo,
		    chunkSize: chunkSize,
	}
}

func (f *flusher) Flush(docs []document.Document) []document.Document {
	chunks, err := utils.SplitDocumentSlice(docs, f.chunkSize)
	if err != nil {
		return docs
	}

	for i := 0; i < len(chunks); i++ {
		if err := f.repo.AddDocs(chunks[i]); err != nil {
			return docs[i * int(f.chunkSize):]
		}
	}

	return nil
}

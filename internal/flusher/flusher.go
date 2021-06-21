package flusher

import (
	"context"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/repo"
	"github.com/ocp-docs-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, docs []document.Document) ([]document.Document, []uint64, error)
}

type flusher struct {
	repo      repo.Repo
	chunkSize int
}

func New(docsRepo repo.Repo, chunkSize int) Flusher {
	return &flusher{
		repo:      docsRepo,
		chunkSize: chunkSize,
	}
}

//  return not added docs and IDs of successful added
func (f *flusher) Flush(ctx context.Context, docs []document.Document) ([]document.Document, []uint64, error) {
	chunks, err := utils.SplitDocumentSlice(docs, f.chunkSize)
	successFullIds := make([]uint64,0, len(docs))
	if err != nil {
		return docs, successFullIds, err
	}
	for i := 0; i < len(chunks); i++ {
		ids, err := f.repo.AddDocs(ctx, chunks[i])
		successFullIds = append(successFullIds, ids...)
		if err != nil {
			return docs[i*int(f.chunkSize):], successFullIds, err
		}
	}

	return nil, successFullIds, nil
}

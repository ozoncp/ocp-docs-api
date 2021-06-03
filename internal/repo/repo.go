package repo

import "github.com/ocp-docs-api/internal/models/document"

type Repo interface {
	AddDocs(docs []document.Document)  error
	RemoveDoc(taskId uint64) error
	DescribeDoc(taskId uint64) (*document.Document, error)
	ListDocs(limit, offset uint64) ([]document.Document, error)
}
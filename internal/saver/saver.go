package saver

import (
	"github.com/ocp-docs-api/internal/flusher"
	"github.com/ocp-docs-api/internal/models/document"
)

type Saver interface {
	Init()
	Save(doc document.Document)
	Close()
}

type saver struct {
	capacity int
	f flusher.Flusher
	loseAllData bool
	data []document.Document
}

func (s *saver) Save(doc document.Document) {

}

func (s *saver) Close() {

}

func (s *saver) Init() {

}

func New() Saver {
	return &saver{}
}
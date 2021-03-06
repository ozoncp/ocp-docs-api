package saver

import (
	"context"
	"github.com/ocp-docs-api/internal/alarmer"
	"github.com/ocp-docs-api/internal/flusher"
	"github.com/ocp-docs-api/internal/models/document"
)

type SaveStrategy int

const (
	DropAll SaveStrategy = iota
	DropOne
)

type Saver interface {
	Init()
	Save(doc document.Document)
	Close()
}

type saver struct {
	capacity int
	f        flusher.Flusher
	a        alarmer.Alarmer
	data     []document.Document
	done     chan struct{}
	strategy SaveStrategy
	docCh    chan document.Document
	ctx      context.Context
}

func New(ctx context.Context, capacity int, f flusher.Flusher, a alarmer.Alarmer, strategy SaveStrategy) Saver {
	done := make(chan struct{})
	data := make([]document.Document, 0, capacity)
	docCh := make(chan document.Document)
	return &saver{
		ctx:      ctx,
		capacity: capacity,
		f:        f,
		a:        a,
		strategy: strategy,
		data:     data,
		done:     done,
		docCh:    docCh,
	}
}

func (s *saver) Save(doc document.Document) {
	s.docCh <- doc
}

func (s *saver) Close() {
	close(s.done)
}

func (s *saver) flushing() {
	for {
		select {
		case <-s.a.Alarm():
			flushRes := s.f.Flush(s.ctx, s.data)
			if flushRes != nil {
				s.data = flushRes
			} else {
				s.data = s.data[:0]
			}
		case <-s.done:
			s.data = s.f.Flush(s.ctx, s.data)
			s.a.Close()
			return
		case task := <-s.docCh:
			if len(s.data) == s.capacity {
				switch s.strategy {
				case DropAll:
					s.data = s.data[:0]
					s.data = append(s.data, task)
				case DropOne:
					s.data = append(s.data[1:], task)
				}
			} else {
				s.data = append(s.data, task)
			}
		}
	}
}

func (s *saver) Init() {
	go s.flushing()
}

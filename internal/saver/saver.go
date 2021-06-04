package saver

import (
	"fmt"
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
	f flusher.Flusher
	a alarmer.Alarmer
	data []document.Document
	done chan struct{}
	strategy SaveStrategy
	docCh chan document.Document
}

func New(capacity int, f flusher.Flusher, a alarmer.Alarmer, strategy SaveStrategy) Saver {
	done := make(chan struct{})
	data := make([]document.Document, capacity)
	docCh := make(chan document.Document)
	return &saver{
		capacity: capacity,
		f: f,
		a: a,
		strategy: strategy,
		data: data,
		done: done,
		docCh: docCh,
	}
}

func (s *saver) Save(doc document.Document) {
	s.docCh <- doc
}

func (s *saver) Close() {
	s.done <- struct{}{}
	close(s.done)
}

func (s *saver) overFlowLogic() {
	switch s.strategy {
	case DropAll:
		s.data = make([]document.Document, s.capacity)
	case DropOne:
		copy(s.data[0:], s.data[1:])
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *saver) flushing() {
	for {
		select {
		case _, ok := <-s.a.Alarm():
			if ok {
				s.data = s.f.Flush(s.data)
				fmt.Println("Flush")
			} else {
				fmt.Println("Alarm channel is not available")
			}
		case <-s.done:
			fmt.Println("Close Saver")
			s.a.Close()
			return
		case task := <-s.docCh:
			if len(s.data) == cap(s.data) {
				s.overFlowLogic()
			}
			s.data = append(s.data, task)
			fmt.Println("push task")
		}
	}
}

func (s *saver) Init() {

}


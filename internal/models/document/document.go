package document

import "fmt"

type Document struct {
	Id         uint64
	Name       string
	Link       string
	SourceLink string
}

func (doc *Document) String() string {
	return fmt.Sprintf("{Id: %v, Name: %v, Link: %v, SourceLink: %v}",
		doc.Id, doc.Name, doc.Link, doc.SourceLink)
}

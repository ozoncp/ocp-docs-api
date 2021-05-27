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
func New(id uint64, Name string, link string, sourceLink string) *Document {
	doc := Document{Id: id, Name: Name, Link: link, SourceLink: sourceLink}
	return &doc
}

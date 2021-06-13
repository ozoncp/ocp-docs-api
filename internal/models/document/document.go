package document

import "fmt"

type Document struct {
	Id         uint64 `db:"id"`
	Name       string `db:"name"`
	Link       string `db:"link"`
	SourceLink string `db:"source_link"`
}

func (doc *Document) String() string {
	return fmt.Sprintf("{Id: %v, Name: %v, Link: %v, SourceLink: %v}",
		doc.Id, doc.Name, doc.Link, doc.SourceLink)
}

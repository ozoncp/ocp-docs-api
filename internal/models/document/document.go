package document

import "fmt"

type Document struct {
	Id         uint64
	DocName    string
	Link       string
	SourceLink string
}

func (doc *Document) String() string {
	return fmt.Sprintf("{Id: %v, DocName: %v, Link: %v, SourceLink: %v}",
		doc.Id, doc.DocName, doc.Link, doc.SourceLink)
}
func New(id uint64, docName string, link string, sourceLink string) *Document {
	doc := Document{Id: id, DocName: docName, Link: link, SourceLink: sourceLink}
	return &doc
}
func (doc *Document) getName() string {
	return doc.DocName
}

func (doc *Document) getLink() string {
	return doc.Link
}

func (doc *Document) getSource() string {
	return doc.SourceLink
}

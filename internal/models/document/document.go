package document

import "fmt"

type Document struct {
	Id      uint64
	DocName string
	Link    string
}

func (doc *Document) String() string {
	return fmt.Sprintf("{ Id: %v, DocName: %v, Link:%v }", doc.Id, doc.Link, doc.Link)
}
func New(id uint64, docName string, link string) *Document {
	doc := Document{Id: id, DocName: docName, Link: link}
	return &doc
}

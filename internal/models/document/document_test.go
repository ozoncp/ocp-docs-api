package document

import (
	"fmt"
	"testing"
)

func TestNewForDocument(t *testing.T) {
	id := uint64(1)
	docName := "Doc1"
	link := "Url1"
	sourceLink := "Url2"
	doc := New(id, docName, link, sourceLink)
	if doc != nil {
		return
	} else {
		t.Error("New() method generate nil object")
	}
}

func TestGetLink(t *testing.T) {
	link := "Url1"
	doc := New(1, "Doc1", link, "Url2")
	if doc.getLink() != link {
		t.Error("result isn't correct")
	}
}

func TestGetName(t *testing.T) {
	name := "Doc1"
	doc := New(1, name, "Url1", "Url2")
	if doc.getName() != name {
		t.Error("result isn't correct")
	}
}

func TestGetSource(t *testing.T) {
	sourceLink := "Url2"
	doc := New(1, "Doc1", "Url1", sourceLink)
	if doc.getSource() != sourceLink {
		t.Error("result isn't correct")
	}
}

func TestGetString(t *testing.T) {
	ref := "{Id: 1, DocName: Doc1, Link: Url1, SourceLink: Url2}"
	doc := New(1, "Doc1", "Url1", "Url2")
	if doc.String() != ref {
		fmt.Println(doc.String())
		t.Errorf("string serialization isn't correct. \n Result   : %s\n Expected : %s ", doc.String(), ref)
	}

}

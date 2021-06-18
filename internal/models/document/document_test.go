package document

import (
	"fmt"
	"testing"
)

func TestNewForDocument(t *testing.T) {
	id := uint64(1)
	Name := "Doc1"
	link := "Url1"
	sourceLink := "Url2"
	doc := &Document{id, Name, link, sourceLink}
	if doc != nil {
		return
	} else {
		t.Error("New() method generate nil object")
	}
}

func TestGetString(t *testing.T) {
	ref := "{Id: 1, Name: Doc1, Link: Url1, SourceLink: Url2}"
	doc := &Document{1, "Doc1", "Url1", "Url2"}
	if doc.String() != ref {
		fmt.Println(doc.String())
		t.Errorf("string serialization isn't correct. \n Result   : %s\n Expected : %s ", doc.String(), ref)
	}
}

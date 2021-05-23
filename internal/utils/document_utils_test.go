package utils

import (
	"github.com/ocp-docs-api/internal/models/document"
	"reflect"
	"testing"
)

func generateSimpleDocSlice() []document.Document {
	sliceDoc := []document.Document{
		{Id: 1, DocName: "test1", Link: "link1"},
		{Id: 2, DocName: "test2", Link: "link2"},
		{Id: 3, DocName: "test3", Link: "link3"},
		{Id: 4, DocName: "test4", Link: "link4"},
		{Id: 5, DocName: "test5", Link: "link5"},
	}
	return sliceDoc
}

func TestSplitDocumentSliceSliceWithTail(t *testing.T) {
	sliceDoc := generateSimpleDocSlice()
	result := SplitDocumentSlice(sliceDoc, 3)
	ref := [][]document.Document{
		{{Id: 1, DocName: "test1", Link: "link1"},
			{Id: 2, DocName: "test2", Link: "link2"},
			{Id: 3, DocName: "test3", Link: "link3"}},
		{{Id: 4, DocName: "test4", Link: "link4"},
			{Id: 5, DocName: "test5", Link: "link5"},
		},
	}
	if !reflect.DeepEqual(ref, result) {
		t.Error("Fail: ", result)
		return
	}
}

func TestConvertDocumentSliceToMap(t *testing.T) {
	sliceDoc := generateSimpleDocSlice()
	result, err := ConvertDocumentSliceToMap(sliceDoc)
	if err == nil {
		ref := map[uint64]document.Document{
			1: {Id: 1, DocName: "test1", Link: "link1"},
			2: {Id: 2, DocName: "test2", Link: "link2"},
			3: {Id: 3, DocName: "test3", Link: "link3"},
			4: {Id: 4, DocName: "test4", Link: "link4"},
			5: {Id: 5, DocName: "test5", Link: "link5"},
		}
		if !reflect.DeepEqual(ref, result) {
			t.Error("Fail: ", result)
			return
		}
	} else {
		t.Error("Fail, err is not nil ")
	}

}

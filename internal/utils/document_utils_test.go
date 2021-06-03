package utils

import (
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitDocumentSlice(t *testing.T) {
	t.Run("SliceWithTail", SplitDocSliceSliceWithTail)
	t.Run("WithChunkZero", SplitDocSliceWithChunkZero)
	t.Run("WithEmptyStruct", SplitDocSliceWithEmptyStruct)
	t.Run("WithNilStruct", SplitDocSliceWithNilStruct)
}

func SplitDocSliceSliceWithTail(t *testing.T) {
	sliceDoc := generateSimpleDocSlice()
	chunkSize := 3
	result, err := SplitDocumentSlice(sliceDoc, chunkSize)
	assert.Empty(t, err)
	ref := [][]document.Document{
		{{Id: 1, Name: "test1", Link: "link1"},
			{Id: 2, Name: "test2", Link: "link2"},
			{Id: 3, Name: "test3", Link: "link3"}},
		{{Id: 4, Name: "test4", Link: "link4"},
			{Id: 5, Name: "test5", Link: "link5"},
		},
	}
	assert.Equal(t, ref, result)
}

func SplitDocSliceWithChunkZero(t *testing.T) {
	sliceDoc := generateSimpleDocSlice()
	result, err := SplitDocumentSlice(sliceDoc, 0)
	assert.Equal(t, err, errorChunkSizeIsInvalid)
	assert.Empty(t, result)
}

func SplitDocSliceWithEmptyStruct(t *testing.T) {
	sliceDoc := []document.Document{}
	result, err := SplitDocumentSlice(sliceDoc, 3)
	assert.Empty(t, err)
	assert.Equal(t, result, [][]document.Document{})
}

func SplitDocSliceWithNilStruct(t *testing.T) {
	chunkSize := 1
	result, err := SplitDocumentSlice(nil, chunkSize)
	assert.Equal(t, err, errorInputSliceIsNil)
	assert.Empty(t, result)
}


func TestConvertDocumentSliceToMap(t *testing.T) {
	t.Run("withSimpleInput", ConvertSimpleInput)
	t.Run("withNilInput", ConvertDocumentSliceToMapNilInput)
	t.Run("withSameKeys", ConvertDocumentSliceToMapSameKeys)
}

func ConvertSimpleInput(t *testing.T) {
	sliceDoc := generateSimpleDocSlice()
	result, err := ConvertDocumentSliceToMap(sliceDoc)
	ref := map[uint64]document.Document{
		1: {Id: 1, Name: "test1", Link: "link1"},
		2: {Id: 2, Name: "test2", Link: "link2"},
		3: {Id: 3, Name: "test3", Link: "link3"},
		4: {Id: 4, Name: "test4", Link: "link4"},
		5: {Id: 5, Name: "test5", Link: "link5"},
	}
	assert.Nil(t, err)
	assert.Equal(t, ref, result)
}

func ConvertDocumentSliceToMapNilInput(t *testing.T) {
	result, err := ConvertDocumentSliceToMap(nil)
	assert.Equal(t, err, errorInputSliceIsNil)
	assert.Empty(t, result)
}

func ConvertDocumentSliceToMapSameKeys(t *testing.T) {
	sliceDoc := []document.Document{
		{Id: 1, Name: "test1", Link: "link1"},
		{Id: 1, Name: "test2", Link: "link2"},
	}
	result, err := ConvertDocumentSliceToMap(sliceDoc)
	assert.Empty(t, result)
	assert.Equal(t, err, errorKeyAlreadyExistInMap)
}

func generateSimpleDocSlice() []document.Document {
	sliceDoc := []document.Document{
		{Id: 1, Name: "test1", Link: "link1"},
		{Id: 2, Name: "test2", Link: "link2"},
		{Id: 3, Name: "test3", Link: "link3"},
		{Id: 4, Name: "test4", Link: "link4"},
		{Id: 5, Name: "test5", Link: "link5"},
	}
	return sliceDoc
}

package utils

import (
	"errors"
	"github.com/ocp-docs-api/internal/models/document"
)

func SplitDocumentSlice(source []document.Document, chunkSize uint) [][]document.Document {
	if source == nil || chunkSize == 0 {
		return nil
	}
	lenSrc := uint(len(source))
	numOfChunks := lenSrc / chunkSize
	if lenSrc%chunkSize != 0 {
		numOfChunks++
	}
	result := make([][]document.Document, numOfChunks)
	for i := uint(0); i < numOfChunks-1; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		result[i] = source[start:end]
	}
	result[numOfChunks-1] = source[(numOfChunks-1)*chunkSize:]
	return result
}

func ConvertDocumentSliceToMap(source []document.Document) (map[uint64]document.Document, error) {
	if source == nil {
		return nil, errors.New("Inpit slice is nil")
	}
	documentMap := make(map[uint64]document.Document, len(source))
	for _, val := range source {
		if _, found := documentMap[val.Id]; found {
			return nil, errors.New("Key already exist in a map")
		}
		documentMap[val.Id] = val
	}
	return documentMap, nil
}

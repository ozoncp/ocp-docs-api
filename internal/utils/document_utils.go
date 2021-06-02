package utils

import (
	"errors"
	"github.com/ocp-docs-api/internal/models/document"
)

var (
	inputSliceIsNil = errors.New("inpit slice is nil")
	chinkSizeIsInvalid = errors.New("chunkSize is invalid")
	keyAlreadyExistInAMap = errors.New("key already exist in a map")
)

func SplitDocumentSlice(source []document.Document, chunkSize int) ([][]document.Document, error) {
	if source == nil {
		return nil, inputSliceIsNil
	}

	if chunkSize <= 0 {
		return nil, chinkSizeIsInvalid
	}

	lenSrc := len(source)
	numOfChunks := lenSrc / chunkSize
	if lenSrc%chunkSize != 0 {
		numOfChunks++
	}

	if numOfChunks == 0 {
		return [][]document.Document{}, nil
	}

	result := make([][]document.Document, numOfChunks)
	for i := 0; i < numOfChunks-1; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		result[i] = source[start:end]
	}
	result[numOfChunks-1] = source[(numOfChunks-1)*chunkSize:]
	return result, nil
}

func ConvertDocumentSliceToMap(source []document.Document) (map[uint64]document.Document, error) {
	if source == nil {
		return nil, inputSliceIsNil
	}
	documentMap := make(map[uint64]document.Document, len(source))
	for _, val := range source {
		if _, found := documentMap[val.Id]; found {
			return nil, keyAlreadyExistInAMap
		}
		documentMap[val.Id] = val
	}
	return documentMap, nil
}

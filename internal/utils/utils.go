package utils

func SwapKeyValue(source map[int]string) map[string]int {
	var result = map[string]int{}
	for key, value := range source {
		result[value] = key
	}
	return result
}

func findInStrSlice(source []string, sym string) bool {
	for _, val := range source {
		if val == sym {
			return true
		}
	}
	return false
}

func FilterSliceByValues(source []string, filter []string) []string {
	var result []string
	for _, val := range source {
		if !findInStrSlice(filter, val) {
			result = append(result, val)
		}
	}
	return result
}

func SplitSliceIntoChunks(source []int, chunkSize int) [][]int {
	numOfChunks := len(source)/chunkSize
	if len(source)%chunkSize != 0 {
		numOfChunks += 1
	}
	result := make([][]int, numOfChunks)

	for i := 0; i < numOfChunks - 1; i++ {
		result[i] = source[i * chunkSize: (i+1) * chunkSize]
	}
	result[numOfChunks - 1] = source[(numOfChunks - 1) * chunkSize: len(source)]
	return result
}
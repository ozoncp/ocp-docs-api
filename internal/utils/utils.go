package utils

func SwapKeyValue(source map[int]string) map[string]int {
	if source == nil {
		return nil
	}
	var result = map[string]int{}
	for key, value := range source {
		if _, found := result[value]; found {
			errMsg := "Error: Result map already contain the key " + value
			panic(errMsg)
		} else {
			result[value] = key
		}
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

func SplitSlice(source []int, chunkSize uint) [][] int {
	if source == nil || chunkSize == 0 {
		return nil
	}
	lenSrc := uint(len(source))
	numOfChunks := lenSrc/chunkSize
	if lenSrc%chunkSize != 0 {
		numOfChunks++
	}
	result := make([][]int, numOfChunks)
	for i := uint(0); i < numOfChunks - 1; i++ {
		start := i * chunkSize
		end := (i+1) * chunkSize
		result[i] = source[start:end]
	}
	result[numOfChunks - 1] = source[(numOfChunks - 1) * chunkSize:]
	return result
}
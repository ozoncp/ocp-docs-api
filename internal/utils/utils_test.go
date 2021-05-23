package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSwapKeyValueCorrect(t *testing.T) {
	inputMap := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}
	ref := map[string]int{
		"zero": 0,
		"one":  1,
		"two":  2,
	}
	res := SwapKeyValue(inputMap)
	if !reflect.DeepEqual(res, ref) {
		t.Error("Fail", res)
		return
	}
}

func TestPanicInSwapKey(t *testing.T) {
	inputMap := map[int]string{
		0: "zero",
		1: "zero",
	}
	assert.Panics(t, func() { SwapKeyValue(inputMap) }, "Fail")
}

func TestFilterSliceByValues(t *testing.T) {
	inSlice := []string{"a", "b", "c", "d", "e", "g"}
	filter := []string{"a", "b", "c"}
	ref := []string{"d", "e", "g"}
	result := FilterSliceByValues(inSlice, filter)
	if !reflect.DeepEqual(ref, result) {
		t.Error("Fail", result)
		return
	}
}

func TestSplitSliceWithTail(t *testing.T) {
	sliceInt := []int{1, 2, 3, 4, 5}
	result := SplitSlice(sliceInt, 3)
	ref := [][]int{{1, 2, 3}, {4, 5}}
	if !reflect.DeepEqual(ref, result) {
		t.Error("Fail: ", result)
		return
	}
}

func TestSplitSliceWithoutTail(t *testing.T) {
	sliceInt := []int{1, 2, 3, 4, 5, 6}
	result := SplitSlice(sliceInt, 2)
	ref := [][]int{{1, 2}, {3, 4}, {5, 6}}
	if !reflect.DeepEqual(ref, result) {
		t.Error("Fail: ", result)
		return
	}
}

func TestThatSplitSliceCOnnectedWithOriginalSlice(t *testing.T) {
	sliceInt := []int{1, 2, 3, 4, 5}
	result := SplitSlice(sliceInt, 3)
	result[0][0] = 0
	ref := []int{0, 2, 3, 4, 5}
	if !reflect.DeepEqual(ref, sliceInt) {
		t.Error("Fail: ", result)
		return
	}
}

package hash

import (
	"testing"
	"lib"
	"fmt"
)

func TestTopKFrequent(t *testing.T) {
	lib.PrintFunc()

	arr := []string{"i", "love", "leetcode", "i", "love", "coding"}
	result := TopKFrequent(arr, 4)
	fmt.Printf("topk:%v\n", result)
}

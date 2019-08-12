package sort

import (
	"fmt"
	"testing"
)

func TestFindTopK(t *testing.T) {
	arr := []int{3, 5, 4, 1, 2, 6}
	lastK := FindTopK(arr, len(arr),4)
	fmt.Printf("lastK:%v\n", lastK)
}

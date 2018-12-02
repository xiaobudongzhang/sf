package _5_array

import (
	"testing"
	"fmt"
)

func TestArray_Insert(t *testing.T) {
	fmt.Printf(" TestArray_Insert start .....\n")
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	arr.Print()
	fmt.Printf(" TestArray_Insert end .....\n")
}

func TestArray_Delete(t *testing.T) {
	fmt.Printf(" TestArray_Delete start .....\n")
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	arr.Delete(4)
	arr.Print()
	fmt.Printf(" TestArray_Delete end .....\n")
}

func TestArray_Print(t *testing.T) {
	fmt.Printf(" TestArray_Print start .....\n")
	capacity := 10
	arr := NewArray(uint(capacity))

	for i:= 0; i< capacity; i++ {
		err := arr.Insert(uint(i), i+1)

		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
	fmt.Printf(" TestArray_Print end .....\n")
}
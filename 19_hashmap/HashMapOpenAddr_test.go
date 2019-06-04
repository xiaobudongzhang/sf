package _9_hashmap

import (
	"fmt"
	"lib"
	"testing"
)

func TestHashMap_Insert(t *testing.T) {
	lib.PrintFunc()

	hashTable := NewHashTable(10)
	hashTable.Insert(1)
	hashTable.Insert(2)
	hashTable.Insert(3)
	hashTable.Insert(33)
	hashTable.Insert(333)
	hashTable.Insert(4)
	hashTable.Insert(5)
	hashTable.Insert(6)
	hashTable.Insert(7)
	hashTable.Insert(8)

	hashTable.Print()

	findIndex := hashTable.Find(33)
	fmt.Printf("findIndex：%v\n", findIndex)

	hashTable.Delete(33)
	hashTable.Print()

	findIndex2 := hashTable.Find(33)
	fmt.Printf("findIndex2：%v\n", findIndex2)

	findIndex3 := hashTable.Find(2)
	fmt.Printf("findIndex3：%v\n", findIndex3)

}

package _5_array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity < 1 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("Out of index range")
	}
	return this.data[index], nil
}

func (this *Array) Insert(index uint, v int) error {
	if this.isIndexOutOfRange(index) {
		return errors.New(fmt.Sprintf("%d %s", index, "out of index range"))
	}
	if this.Len() >= uint(cap(this.data)) {
		return errors.New("Full array")
	}
	for i := this.Len(); i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("Out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

func (this *Array) Print() {
	for i := uint(0); i < this.Len(); i++ {
		fmt.Printf("%d \n", this.data[i])
	}
}

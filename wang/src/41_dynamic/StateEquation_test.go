package _1_dynamic

import (
	"testing"
	"lib"
	"fmt"
)

func TestMindistP(t *testing.T) {
	lib.PrintFunc()

	min := MindistP()
	fmt.Printf("%v\n", min)
}
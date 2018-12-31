package _7_greedy

import (
	"testing"
	"lib"
)

func TestSpace(t *testing.T) {
	lib.PrintFunc()
	point := []*Node{
		{l:6,r:8},
		{l:2,r:4},
		{l:3,r:5},
		{l:1,r:5},
		{l:5,r:9},
		{l:8,r:10},
	}
	Space(point)
}
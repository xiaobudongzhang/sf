package common

import "math"

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}


func Height(p *Node) int {
	lHeight := 0
	rHeight := 0

	if p == nil {
		return 0
	}

	lHeight = Height(p.left)
	rHeight = Height(p.right)

	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

func Height2(p *Node) float64 {
	if p == nil {
		return 0
	}
	return math.Max(Height2(p.left), Height2(p.right)) + 1
}


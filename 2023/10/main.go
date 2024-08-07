package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value       rune
	Next        *Node
	Previous    *Node
	IsConnected bool
	X           int
	Y           int
}

func NewNode(value rune, x, y int) *Node {
	node := &Node{Value: value, IsConnected: false, X: x, Y: y}
	return node
}

func (n *Node) GetNextLoopNode(up, left, right, down *Node) {
	// Find next node based on previous node, and what comes next based on pipe
	switch n.Value {
	case '|':
		if up == n.Previous {
			// if came from up node, next must be down etc.
			n.Next = down
		} else {
			n.Next = up
		}
	case '-':
		if left == n.Previous {
			n.Next = right
		} else {
			n.Next = left
		}
	case 'L':
		if up == n.Previous {
			n.Next = right
		} else {
			n.Next = up
		}
	case 'J':
		if up == n.Previous {
			n.Next = left
		} else {
			n.Next = up
		}
	case '7':
		if left == n.Previous {
			n.Next = down
		} else {
			n.Next = left
		}
	case 'F':
		if right == n.Previous {
			n.Next = down
		} else {
			n.Next = right
		}
	case 'S':
		if up != nil && (up.Value == 'F' || up.Value == '|' || up.Value == '7') {
			n.Next = up
		} else if right != nil && (right.Value == '-' || right.Value == 'J' || right.Value == '7') {
			n.Next = right
		} else if down != nil && (down.Value == '|' || down.Value == 'L' || down.Value == 'J') {
			n.Next = down
		} else {
			n.Next = left
		}
	default:
		n.Next = nil
	}

	// Setup link between nodes
	if n.Next != nil {
		n.Next.Previous = n
		n.Next.IsConnected = true
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%c", n.Value)
}

type NodeMap [][]*Node

func NewNodeMap(nodeLines []string) NodeMap {
	result := make([][]*Node, len(nodeLines))

	for i, line := range nodeLines {
		result[i] = make([]*Node, len(line))

		for i2, v := range line {
			result[i][i2] = NewNode(v, i2, i)
		}
	}

	return result
}

func FindLoop(nodeMap NodeMap) *Node {
	// find starting node
	x, y := 0, 0
	var current *Node
	for ; y < len(nodeMap); y++ {
		x = 0
		for ; x < len(nodeMap[y]); x++ {
			if nodeMap[y][x].Value == 'S' {
				current = nodeMap[y][x]
				break
			}
		}
		if current != nil {
			break
		}
	}

	// iterate around nodes finding next node in loop until reaching starting node again
	for current.Next == nil || current.Previous == nil {
		var up, down, left, right *Node

		// get nodes around current node
		if current.Y > 0 {
			up = nodeMap[current.Y-1][current.X]
		}
		if current.Y < len(nodeMap)-1 {
			down = nodeMap[current.Y+1][current.X]
		}
		if current.X > 0 {
			left = nodeMap[current.Y][current.X-1]
		}
		if current.X < len(nodeMap[current.Y])-1 {
			right = nodeMap[current.Y][current.X+1]
		}
		current.GetNextLoopNode(up, left, right, down)
		fmt.Printf("v: %s x: %d y: %d; n: %s x: %d y: %d\n", current, current.X, current.Y, current.Next, current.Next.X, current.Next.Y)
		current = current.Next
	}

	// return starting node
	return current

}

func FindFurthestPointInLoop(node *Node) int {
	current := node.Next
	result := 1
	for current.Value != 'S' {
		result += 1
		current = current.Next
	}
	return result / 2
}

func part1(file string) int {
	split := strings.Split(file, "\n")
	nodeMap := NewNodeMap(split)
	startingNode := FindLoop(nodeMap)
	return FindFurthestPointInLoop(startingNode)
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d", part1(string(f)))

}

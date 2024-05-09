package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Val   string
	Left  string
	Right string
}

func (n *Node) String() string {
	return fmt.Sprintf("val: %s, left: %s, right: %s", n.Val, n.Left, n.Right)
}

// strip: helper function to remove all non alpha characters to parse node correctly
func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('A' <= b && b <= 'Z') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// parseNetwork: build tree of nodes from input network strings
func parseNetwork(network []string) map[string]*Node {
	m := make(map[string]*Node)
	for _, s := range network {
		// trim string of punctuation and parse fields
		s = strip(s)
		f := strings.Fields(s)
		m[f[0]] = &Node{
			Val:   f[0],
			Left:  f[1],
			Right: f[2],
		}
	}
	return m
}

func getLengthToTarget(networkMap map[string]*Node, directions, start, target string) int {
	node := networkMap[start]
	if node == nil { // starting node not in network map
		return -1
	}
	fmt.Printf("initial node: %s\n", node)
	length := 0
	for {
		for i := 0; i < len(directions); i++ {
			d := directions[i]
			if d == 'L' {
				node = networkMap[node.Left]
				fmt.Printf("after left turn, node: %s\n", node)
			} else {
				node = networkMap[node.Right]
				fmt.Printf("after right turn, node: %s\n", node)
			}
			if node.Val == target { // found node, return how many directions taken
				return length + (i + 1)
			}
		}
		length += len(directions) // could not find in current set of directions, iterate through again
	}
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	directions := split[0]
	network := split[2:]

	networkMap := parseNetwork(network)

	return getLengthToTarget(networkMap, directions, "AAA", "ZZZ")
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(string(input)))
}

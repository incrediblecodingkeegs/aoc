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
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// parseNetwork: return map of all nodes from input network strings and list of starting nodes
func parseNetwork(network []string) (map[string]*Node, []*Node) {
	m := make(map[string]*Node)
	l := make([]*Node, 0)
	for _, s := range network {
		// trim string of punctuation and parse fields
		s = strip(s)
		f := strings.Fields(s)
		val := f[0]
		node := &Node{
			Val:   val,
			Left:  f[1],
			Right: f[2],
		}
		m[f[0]] = node

		// node is starting node if node.Val ends with an 'A'
		if val[2] == 'A' {
			l = append(l, node)
		}

	}
	//fmt.Printf("starting nodes: %+v length: %d\n", l, len(l))
	return m, l
}

func getTargetHits(targetHitChan chan<- int, node *Node, networkMap map[string]*Node, directions string) {
	//fmt.Printf("initial node: %s\n", node)
	length := 1
	for {
		for i := 0; i < len(directions); i++ {
			d := directions[i]
			if d == 'L' {
				node = networkMap[node.Left]
				//fmt.Printf("after left turn, node: %s\n", node)
			} else {
				node = networkMap[node.Right]
				//fmt.Printf("after right turn, node: %s\n", node)
			}
			if node.Val[2] == 'Z' { // node is on target node, iterate num directions taken
				//fmt.Printf("sending target hit at %d\n", length+1)
				targetHitChan <- length + i
			}
		}
		length += len(directions) // could not find in current set of directions, iterate through again
	}
}

func getLoopLength(node *Node, networkMap map[string]*Node, directions string) int {
	//fmt.Printf("initial node: %s\n", node)
	length := 1
	for {
		for i := 0; i < len(directions); i++ {
			d := directions[i]
			if d == 'L' {
				node = networkMap[node.Left]
				//fmt.Printf("after left turn, node: %s\n", node)
			} else {
				node = networkMap[node.Right]
				//fmt.Printf("after right turn, node: %s\n", node)
			}
			if node.Val[2] == 'Z' { // node is on target node, iterate num directions taken
				fmt.Printf("loop length = %d\n", length+1)
				return length + i
			}
		}
		length += len(directions) // could not find in current set of directions, iterate through again
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(i []int) int {
	if len(i) == 0 {
		return -1
	} else if len(i) == 1 {
		return i[0]
	} else {
		a := i[0]
		b := i[1]

		i[1] = a * b / GCD(a, b)
		return LCM(i[1:])
	}

}

func getLengthToTarget(networkMap map[string]*Node, directions string, startingNodes []*Node) int {
	loopLengths := make([]int, 0)
	for _, node := range startingNodes {
		loopLengths = append(loopLengths, getLoopLength(node, networkMap, directions))
	}

	return LCM(loopLengths)

}

func part2(input string) int {
	split := strings.Split(input, "\n")
	directions := split[0]
	network := split[2:]

	networkMap, startingNodes := parseNetwork(network)

	return getLengthToTarget(networkMap, directions, startingNodes)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part2: %d\n", part2(string(input)))
}

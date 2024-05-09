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

func getLengthToTarget(networkMap map[string]*Node, directions string, startingNodes []*Node) int {
	var targetHitChans = make([]chan int, len(startingNodes))
	//var wg sync.WaitGroup

	// start worker threads
	for i := 0; i < len(startingNodes); i++ {
		targetHitChans[i] = make(chan int)
		go getTargetHits(targetHitChans[i], startingNodes[i], networkMap, directions)
	}

	// find common target hits
	hcd := 0
	targetHits := make([][]int, len(startingNodes))
	for {
		// read results from channels
		for i := 0; i < len(targetHitChans); i++ {
			// read off channel
			//fmt.Printf("reading off channel %d\n", i)
			targetHit := <-targetHitChans[i]
			targetHits[i] = append(targetHits[i], targetHit)
			//fmt.Printf("target hit at %d\n", targetHit)
		}

		//fmt.Printf("targetHits: %+v\n", targetHits)

		cleanupRequired := false
		numMatches := 0
		for _, index := range targetHits {
			if len(index) == 0 { // need all workers to deliver results first
				cleanupRequired = false
				break
			}

			// remove all values that can no longer be the possible index
			if hcd == 0 {
				hcd = index[0]
			} else if index[0] > hcd {
				hcd = index[0]
				// cleanup structure as there are values in here that cannot be the correct answer
				cleanupRequired = true
			} else if index[0] == hcd {
				numMatches += 1
			}
		}

		if numMatches >= 3 {
			fmt.Printf("matches: %d len(targetHits[0]): %d hcd: %d\n", numMatches, len(targetHits[0]), hcd)
			//	fmt.Printf("targetHits[0]: %+v\n", targetHits[0])
			//	fmt.Printf("targetHits[1]: %+v\n", targetHits[1])
			//	fmt.Printf("targetHits[2]: %+v\n", targetHits[2])
			//	fmt.Printf("targetHits[3]: %+v\n", targetHits[3])
			//	fmt.Printf("targetHits[4]: %+v\n", targetHits[4])
			//	fmt.Printf("targetHits[5]: %+v\n", targetHits[5])
		}
		if numMatches == len(startingNodes) {
			return hcd
		}

		// garbage collect structure
		// remove all values lower than the highest common denominator as they cannot be the correct answer
		if cleanupRequired {
			//fmt.Printf("cleaning up...\n")
			for i, index := range targetHits {
				if numMatches >= 2 {
					//fmt.Printf("targetHits[%d] before cleanup: %+v\n", i, targetHits[i])
				}
				for len(index) > 0 && index[0] < hcd {
					index = index[1:]
				}
				targetHits[i] = index
				if numMatches >= 2 {
					//fmt.Printf("targetHits[%d] after cleanup: %+v\n", i, targetHits[i])
				}
			}

			//fmt.Printf("targetHits after cleanup: %+v\n", targetHits)
		}
	}
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

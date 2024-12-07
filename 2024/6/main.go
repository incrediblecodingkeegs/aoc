package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	isVisited         bool
	isObstacle        bool
	visitedDirections map[direction]bool
	triedObstacleHere bool
}

func (n *node) String() string {
	//if n.triedObstacleHere {
	//	return "0"
	//} else if n.isVisited {
	if n.isVisited {
		return "X"
	} else if n.isObstacle {
		return "#"
	} else {
		return "."
	}
}

func printNodes(nodes [][]*node) {
	for _, nodeRow := range nodes {
		rowString := ""
		for _, n := range nodeRow {
			rowString += n.String()
		}
		fmt.Println(rowString)
	}
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

type guard struct {
	x int
	y int
	d direction
}

func (g *guard) setNextPosition(nodes [][]*node) bool {
	for {
		switch g.d {
		case up:
			if g.y > 0 { // if not at top of map
				nextNode := nodes[g.y-1][g.x]
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = right
				} else {
					g.y--
					nextNode.isVisited = true
					return false
				}
			} else {
				return true
			}
		case right:
			if g.x < len(nodes) { // if not at edge of map
				nextNode := nodes[g.y][g.x+1]
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = down
				} else {
					g.x++
					nextNode.isVisited = true
					return false
				}
			} else {
				return true
			}
		case down:
			if g.y < len(nodes)-1 { // if not at edge of map
				nextNode := nodes[g.y+1][g.x]
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = left
				} else {
					g.y++
					nextNode.isVisited = true
					return false
				}
			} else {
				return true
			}
		case left:
			if g.x > 0 { // if not at edge of map
				nextNode := nodes[g.y][g.x-1]
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = up
				} else {
					g.x--
					nextNode.isVisited = true
					return false
				}
			} else {
				return true
			}
		}
	}
}

func (g *guard) wouldLoop(nodes [][]*node) bool {
	objectPlaced := false
	for {
		switch g.d {
		case up:
			if g.y > 0 { // if not at edge of map
				nextNode := nodes[g.y-1][g.x]
				if !objectPlaced && !nextNode.triedObstacleHere && !nextNode.isObstacle {
					nextNode.isObstacle = true
					objectPlaced = true
					break
				} else if !objectPlaced {
					return false // have already tried placing an obstacle here, or object already there
				}
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = right
				} else {
					// check if next node has already been visited with the same direction
					if nextNode.visitedDirections[g.d] {
						return true
					}
					nextNode.isVisited = true
					nextNode.visitedDirections[g.d] = true
					g.y--
				}
			} else { // guard exits
				return false
			}
		case right:
			if g.x < len(nodes)-1 { // if not at edge of map
				nextNode := nodes[g.y][g.x+1]
				if !objectPlaced && !nextNode.triedObstacleHere && !nextNode.isObstacle {
					nextNode.isObstacle = true
					objectPlaced = true
					break
				} else if !objectPlaced {
					return false // have already tried placing an obstacle here, or object already there
				}
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = down
				} else {
					// check if next node has already been visited with the same direction
					if nextNode.visitedDirections[g.d] {
						return true
					}
					nextNode.isVisited = true
					nextNode.visitedDirections[g.d] = true
					g.x++
				}
			} else { // guard exits
				return false
			}
		case down:
			if g.y < len(nodes)-1 { // if not at edge of map
				nextNode := nodes[g.y+1][g.x]
				if !objectPlaced && !nextNode.triedObstacleHere && !nextNode.isObstacle {
					nextNode.isObstacle = true
					objectPlaced = true
					break
				} else if !objectPlaced {
					return false // have already tried placing an obstacle here, or object already there
				}
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = left
				} else {
					// check if next node has already been visited with the same direction
					if nextNode.visitedDirections[g.d] {
						return true
					}
					nextNode.isVisited = true
					nextNode.visitedDirections[g.d] = true
					g.y++
				}
			} else { // guard exits
				return false
			}
		case left:
			if g.x > 0 { // if not at edge of map
				nextNode := nodes[g.y][g.x-1]
				if !objectPlaced && !nextNode.triedObstacleHere && !nextNode.isObstacle {
					nextNode.isObstacle = true
					objectPlaced = true
					break
				} else if !objectPlaced {
					return false // have already tried placing an obstacle here, or object already there
				}
				if nextNode.isObstacle { // if next node is obstacle, turn guard
					g.d = up
				} else {
					// check if next node has already been visited with the same direction
					if nextNode.visitedDirections[g.d] {
						return true
					}
					nextNode.isVisited = true
					nextNode.visitedDirections[g.d] = true
					g.x--
				}
			} else { // guard exits
				return false
			}
		}
	}

}

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Printf("Part One: %d\n", partOne(string(f)))
	fmt.Printf("Part Two: %d\n", partTwo(string(f)))
}

func partOne(input string) int {
	nodes, g := parseMap(input)
	//fmt.Println("nodes starting:")
	//printNodes(nodes)
	for {
		if g.setNextPosition(nodes) {
			break
		}
	}

	result := 0
	//fmt.Println("nodes final:")
	//printNodes(nodes)
	for _, nodeRow := range nodes {
		for _, n := range nodeRow {
			if n.isVisited {
				result += 1
			}
		}
	}

	return result
}

func partTwo(input string) int {
	currentNodes, g := parseMap(input)
	startingNodes := deepCopyMap(currentNodes)
	var futureGuard *guard
	result := 0
	for {

		// setup new copies to check loops in the future
		futureNodes := deepCopyMap(startingNodes)
		futureGuard = copyGuard(&g)

		// check if guard looped
		guardLooped := futureGuard.wouldLoop(futureNodes)

		// check if guard exited
		exited := g.setNextPosition(currentNodes)

		if guardLooped {
			fmt.Printf("\nloop %d:\n", result)
			printNodes(futureNodes)

			result += 1
			// update starting nodes to indicate to future that an obstacle has been tried here
			currentNodes[g.y][g.x].triedObstacleHere = true
			startingNodes[g.y][g.x].triedObstacleHere = true
		}
		if exited {
			break
		}
	}
	return result

}

func deepCopyMap(input [][]*node) [][]*node {
	output := make([][]*node, len(input))
	for i, nodes := range input {
		output[i] = make([]*node, len(nodes))
		for j, n := range nodes {
			nodeCopy := *n
			if n.visitedDirections != nil {
				directionsCopy := make(map[direction]bool)
				for k, v := range n.visitedDirections {
					directionsCopy[k] = v
				}
				nodeCopy.visitedDirections = directionsCopy
			}
			output[i][j] = &nodeCopy
		}
	}
	return output
}

func copyGuard(input *guard) *guard {
	output := *input
	return &output
}

func parseMap(input string) ([][]*node, guard) {
	inputLines := strings.Split(input, "\n")
	nodes := make([][]*node, len(inputLines))
	var gp guard

	for i, line := range inputLines {
		nodes[i] = make([]*node, len(line))
		for j, r := range line {
			switch r {
			case '.':
				nodes[i][j] = &node{isVisited: false, isObstacle: false}
			case '#':
				nodes[i][j] = &node{isVisited: false, isObstacle: true}
			case '^':
				nodes[i][j] = &node{isVisited: true, isObstacle: false}
				gp = guard{
					x: j,
					y: i,
					d: up,
				}
			case '>':
				nodes[i][j] = &node{isVisited: true, isObstacle: false}
				gp = guard{
					x: j,
					y: i,
					d: right,
				}
			case 'v':
				nodes[i][j] = &node{isVisited: true, isObstacle: false}
				gp = guard{
					x: j,
					y: i,
					d: down,
				}
			case '<':
				nodes[i][j] = &node{isVisited: true, isObstacle: false}
				gp = guard{
					x: j,
					y: i,
					d: left,
				}
			default:
				nodes[i][j] = &node{isVisited: false, isObstacle: false}
			}
			nodes[i][j].visitedDirections = make(map[direction]bool, 4)
		}
	}
	return nodes, gp
}

package main

import (
	"testing"
)

func TestNewNodeMap(t *testing.T) {
	type args struct {
		nodeLines []string
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			name: "simple__want_correct_lengths",
			args: args{
				nodeLines: []string{
					"-L|F7",
					"7S-7|",
					"L|7||",
					"-L-J|",
					"L|-JF",
				},
			},
			wantX: 5,
			wantY: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNodeMap(tt.args.nodeLines)
			//for i, nodes := range got {
			//	fmt.Printf("%d: %+v\n", i, nodes)
			//}
			if len(got[0]) != tt.wantX {
				t.Errorf("NewNodeMap() = %v, wantX %v", len(got[0]), tt.wantX)
			}
			if len(got) != tt.wantY {
				t.Errorf("NewNodeMap() = %v, wantY %v", len(got), tt.wantY)
			}
			if got[0][0].Value != rune(tt.args.nodeLines[0][0]) {
				t.Errorf("NewNodeMap() not loading nodes correctly; want[0][0] %c, got %c", got[0][0].Value, rune(tt.args.nodeLines[0][0]))
			}
		})
	}
}

func TestFindLoop(t *testing.T) {
	testNodeMap := NewNodeMap([]string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	})

	type args struct {
		nodeMap NodeMap
	}
	tests := []struct {
		name            string
		args            args
		wantNodesInLoop [][]bool
	}{
		{
			name: "testNodeMap__want_correct_loop",
			args: args{nodeMap: testNodeMap},
			wantNodesInLoop: [][]bool{
				{false, false, false, false, false},
				{false, true, true, true, false},
				{false, true, false, true, false},
				{false, true, true, true, false},
				{false, false, false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindLoop(tt.args.nodeMap)
			for y, bools := range tt.wantNodesInLoop {
				for x, b := range bools {
					if b != tt.args.nodeMap[y][x].IsConnected {
						t.Errorf("FindLoop() nodeMap[%d][%d].isConnected = %v, wanted %v", y, x, tt.args.nodeMap[y][x].IsConnected, b)
					}
				}
			}
		})
	}
}

func TestFindFurthestPointInLoop(t *testing.T) {
	simpleNodeMap := NewNodeMap([]string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	})
	complexNodeMap := NewNodeMap([]string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	})
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple_loop__want_4",
			args: args{node: FindLoop(simpleNodeMap)},
			want: 4,
		},
		{
			name: "complex_loop__want_8",
			args: args{node: FindLoop(complexNodeMap)},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFurthestPointInLoop(tt.args.node); got != tt.want {
				t.Errorf("FindFurthestPointInLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

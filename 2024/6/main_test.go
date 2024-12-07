package main

import (
	"testing"
)

const (
	example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
)

func Test_parseMap(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		wantDirection direction
		wantNodeRows  int
		wantNodeCols  int
	}{
		{
			name: "example_map__want_success",
			args: args{
				input: example,
			},
			wantDirection: up,
			wantNodeRows:  10,
			wantNodeCols:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodes, gotDirection := parseMap(tt.args.input)
			if tt.wantNodeRows != len(gotNodes) {
				t.Errorf("parseMap() len(gotNodes) = %d, want %d", len(gotNodes), tt.wantNodeRows)
			}
			if tt.wantNodeCols != len(gotNodes[0]) {
				t.Errorf("parseMap() len(gotNodes[0]) = %d, want %d", len(gotNodes[0]), tt.wantNodeRows)
			}
			if tt.wantDirection != gotDirection.d {
				t.Errorf("parseMap() gotDirection.d = %d, want %d", gotDirection.d, tt.wantDirection)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example_part_one__want_41",
			args: args{
				input: example,
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.input); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example_part_two__want_6",
			args: args{
				input: example,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.input); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

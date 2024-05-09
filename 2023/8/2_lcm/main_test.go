package main

import (
	"reflect"
	"testing"
)

func Test_parseNetwork(t *testing.T) {
	type args struct {
		network []string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]*Node
		want1 []*Node
	}{
		{
			name: "simpleExample__want_success",
			args: args{
				network: []string{
					"AAA = (BBB, BBB)",
					"BBB = (AAA, ZZZ)",
					"ZZZ = (ZZZ, ZZZ)",
				},
			},
			want: map[string]*Node{
				"AAA": {
					Val:   "AAA",
					Left:  "BBB",
					Right: "BBB",
				},
				"BBB": {
					Val:   "BBB",
					Left:  "AAA",
					Right: "ZZZ",
				},
				"ZZZ": {
					Val:   "ZZZ",
					Left:  "ZZZ",
					Right: "ZZZ",
				},
			},
			want1: []*Node{
				{
					Val:   "AAA",
					Left:  "BBB",
					Right: "BBB",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseNetwork(tt.args.network)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNetwork() = %+v, want %+v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseNetwork() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}

func Test_getLengthToTarget(t *testing.T) {
	type args struct {
		network       []string
		directions    string
		startingNodes []*Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example__want_2",
			args: args{
				network: []string{
					"11A = (11B, XXX)",
					"11B = (XXX, 11Z)",
					"11Z = (11B, XXX)",
					"22A = (22B, XXX)",
					"22B = (22C, 22C)",
					"22C = (22Z, 22Z)",
					"22Z = (22B, 22B)",
					"XXX = (XXX, XXX)",
				},
				directions: "LR",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			networkMap, startingNodes := parseNetwork(tt.args.network)
			if got := getLengthToTarget(networkMap, tt.args.directions, startingNodes); got != tt.want {
				t.Errorf("getLengthToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10_15__want_30",
			args: args{
				i: []int{10, 15},
			},
			want: 30,
		}, {
			name: "1_to_10__want_2520",
			args: args{
				i: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 2520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.i); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

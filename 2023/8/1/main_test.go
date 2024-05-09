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
		name string
		args args
		want map[string]*Node
	}{
		{
			name: "simpleExample",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNetwork(tt.args.network); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNetwork() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_getLengthToTarget(t *testing.T) {
	type args struct {
		networkMap map[string]*Node
		directions string
		start      string
		target     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example__want_2",
			args: args{
				networkMap: parseNetwork([]string{
					"AAA = (BBB, CCC)",
					"BBB = (DDD, EEE)",
					"CCC = (ZZZ, GGG)",
					"DDD = (DDD, DDD)",
					"EEE = (EEE, EEE)",
					"GGG = (GGG, GGG)",
					"ZZZ = (ZZZ, ZZZ)",
				}),
				directions: "RL",
				start:      "AAA",
				target:     "ZZZ",
			},
			want: 2,
		}, {
			name: "example__want_6",
			args: args{
				networkMap: parseNetwork([]string{
					"AAA = (BBB, BBB)",
					"BBB = (AAA, ZZZ)",
					"ZZZ = (ZZZ, ZZZ)",
				}),
				directions: "LLR",
				start:      "AAA",
				target:     "ZZZ",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLengthToTarget(tt.args.networkMap, tt.args.directions, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("getLengthToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

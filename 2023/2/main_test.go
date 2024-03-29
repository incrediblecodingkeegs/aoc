package main

import (
	"reflect"
	"testing"
)

func Test_parseGameString(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name string
		args args
		want *Game
	}{
		{
			name: "happy_path_1",
			args: args{
				game: "Game 1: 3 maxBlue, 4 maxRed; 1 maxRed, 2 maxGreen, 6 maxBlue; 2 maxGreen",
			},
			want: &Game{
				id:       1,
				maxRed:   4,
				maxGreen: 2,
				maxBlue:  6,
			},
		}, {
			name: "happy_path_2",
			args: args{
				game: "Game 3: 8 maxGreen, 6 maxBlue, 20 maxRed; 5 maxBlue, 4 maxRed, 13 maxGreen; 5 maxGreen, 1 maxRed",
			},
			want: &Game{
				id:       3,
				maxRed:   20,
				maxGreen: 13,
				maxBlue:  6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGameString(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGameString() = %v, want %v", got, tt.want)
			}
		})
	}
}

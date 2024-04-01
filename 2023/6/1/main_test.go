package main

import (
	"reflect"
	"testing"
)

func Test_parseRaces(t *testing.T) {
	type args struct {
		timeString     string
		distanceString string
	}
	tests := []struct {
		name string
		args args
		want []*Race
	}{
		{
			name: "parse_races__want_success",
			args: args{
				timeString:     "Time:      7  15   30",
				distanceString: "Distance:  9  40  200",
			},
			want: []*Race{
				{
					time:     7,
					distance: 9,
				}, {
					time:     15,
					distance: 40,
				}, {
					time:     30,
					distance: 200,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRaces(tt.args.timeString, tt.args.distanceString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumWinningChances(t *testing.T) {
	type args struct {
		race *Race
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple_race__want_success",
			args: args{
				race: &Race{
					time:     7,
					distance: 9,
				},
			},
			want: 4,
		}, {
			name: "simple_race_2__want_success",
			args: args{
				race: &Race{
					time:     15,
					distance: 40,
				},
			},
			want: 8,
		}, {
			name: "simple_race_#__want_success",
			args: args{
				race: &Race{
					time:     30,
					distance: 200,
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumWinningChances(tt.args.race); got != tt.want {
				t.Errorf("getNumWinningChances() = %v, want %v", got, tt.want)
			}
		})
	}
}

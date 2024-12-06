package main

import "testing"

func Test_partOne(t *testing.T) {
	type args struct {
		reports [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.reports); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		reports [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
					{1, 4, 3, 2, 1},
					{36, 34, 36, 33, 30},
					{5, 8, 10, 13, 15, 16, 17, 21},
					{82, 85, 87, 88, 91, 93, 97, 94},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.reports); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

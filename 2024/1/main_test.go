package main

import "testing"

func Test_partOne(t *testing.T) {
	type args struct {
		col1 []int
		col2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				col1: []int{3, 4, 2, 1, 3, 3},
				col2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.col1, tt.args.col2); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		col1 []int
		col2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				col1: []int{3, 4, 2, 1, 3, 3},
				col2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.col1, tt.args.col2); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

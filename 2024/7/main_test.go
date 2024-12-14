package main

import "testing"

func Test_partOneEvaluate(t *testing.T) {
	type args struct {
		answer int
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example_1__want_true",
			args: args{
				answer: 190,
				values: []int{10, 19},
			},
			want: true,
		}, {
			name: "example_2__want_true",
			args: args{
				answer: 3267,
				values: []int{81, 40, 27},
			},
			want: true,
		}, {
			name: "example_3__want_false",
			args: args{
				answer: 83,
				values: []int{17, 5},
			},
			want: false,
		}, {
			name: "example_4__want_true",
			args: args{
				answer: 292,
				values: []int{11, 6, 16, 20},
			},
			want: true,
		}, {
			name: "edge_1__want_true",
			args: args{
				answer: 1460,
				values: []int{11, 6, 16, 20, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOneEvaluate(tt.args.answer, tt.args.values); got != tt.want {
				t.Errorf("partOneEvaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwoEvaluate(t *testing.T) {
	type args struct {
		answer int
		values []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example_1__want_true",
			args: args{
				answer: 190,
				values: []int{10, 19},
			},
			want: true,
		}, {
			name: "example_2__want_true",
			args: args{
				answer: 3267,
				values: []int{81, 40, 27},
			},
			want: true,
		}, {
			name: "example_3__want_false",
			args: args{
				answer: 83,
				values: []int{17, 5},
			},
			want: false,
		}, {
			name: "example_4__want_true",
			args: args{
				answer: 156,
				values: []int{15, 6},
			},
			want: true,
		}, {
			name: "example_5__want_true",
			args: args{
				answer: 7290,
				values: []int{6, 8, 6, 15},
			},
			want: true,
		}, {
			name: "example_6__want_true",
			args: args{
				answer: 192,
				values: []int{17, 8, 14},
			},
			want: true,
		}, {
			name: "edge_1__want_true",
			args: args{
				answer: 1460,
				values: []int{11, 6, 16, 20, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwoEvaluate(tt.args.answer, tt.args.values); got != tt.want {
				t.Errorf("partOneEvaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7_old 18 13
292: 11 6 16 20`
)

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
			name: "example_1__want_3749",
			args: args{
				input: example,
			},
			want: 3749,
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
			name: "example_1__want_11387",
			args: args{
				input: example,
			},
			want: 11387,
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

func Test_concat(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "15_6__want_156",
			args: args{
				a: 15,
				b: 6,
			},
			want: 156,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concat(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

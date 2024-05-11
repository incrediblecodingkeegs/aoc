package main

import "testing"

func Test_getNextValue(t *testing.T) {
	type args struct {
		history []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0_3_6_9_12_15__want_18",
			args: args{
				history: []int{0, 3, 6, 9, 12, 15},
			},
			want: 18,
		}, {
			name: "1_3_6_10_15_21__want_28",
			args: args{
				history: []int{1, 3, 6, 10, 15, 21},
			},
			want: 28,
		}, {
			name: "10_13_16_21_30_45__want_68",
			args: args{
				history: []int{10, 13, 16, 21, 30, 45},
			},
			want: 68,
		}, {
			name: "9_6_3_0_-3_-6_-9__want_-12",
			args: args{
				history: []int{9, 6, 3, 0, -3, -6, -9},
			},
			want: -12,
		}, {
			name: "-2_6_32_90_203_421_863_1799_3804_8041_16764_34188_67988_131949_250881_470219_875478_1632236_3067792_5834670__want_11229253",
			args: args{
				history: []int{-2, 6, 32, 90, 203, 421, 863, 1799, 3804, 8041, 16764, 34188, 67988, 131949, 250881, 470219, 875478, 1632236, 3067792, 5834670},
			},
			want: 11229253,
		}, {
			name: "13_22_39_70_140_304_660_1370_2693_5028_8955_15248_24816_38506_56676_78416_100261_114202_104759_44834__want_-109988",
			args: args{
				//
				//						 41  -37 71
				//					   9   50  13  84
				// 					 6   15  65  78   162
				// 				   8   14  39  104  182  344
				//				 9 	 17  31  70   174  356  710   1323
				history: []int{13, 22, 39, 70, 140, 304, 660, 1370, 2693, 5028, 8955, 15248, 24816, 38506, 56676, 78416, 100261, 114202, 104759, 44834},
			},
			want: -109988,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextValue(tt.args.history); got != tt.want {
				t.Errorf("getNextValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "part1example__want_114",
			args: args{
				input: "0   3   6   9  12  15\n" +
					"1   3   6  10  15  21\n" +
					"10  13  16  21  30  45",
			},
			want: 114,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPreviousValue(t *testing.T) {
	type args struct {
		history []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0_3_6_9_12_15__want_-3",
			args: args{
				history: []int{0, 3, 6, 9, 12, 15},
			},
			want: -3,
		}, {
			name: "1_3_6_10_15_21__want_0",
			args: args{
				history: []int{1, 3, 6, 10, 15, 21},
			},
			want: 0,
		}, {
			name: "10_13_16_21_30_45__want_5",
			args: args{
				history: []int{10, 13, 16, 21, 30, 45},
			},
			want: 5,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPreviousValue(tt.args.history); got != tt.want {
				t.Errorf("getPreviousValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

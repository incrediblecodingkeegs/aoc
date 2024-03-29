package main

import "testing"

func Test_getCalibrationValue1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy_path_1",
			args: args{
				value: "two1nine",
			},
			want: 29,
		}, {
			name: "happy_path_2",
			args: args{
				value: "eightwothree",
			},
			want: 83,
		}, {
			name: "happy_path_3",
			args: args{
				value: "abcone2threexyz",
			},
			want: 13,
		}, {
			name: "happy_path_4",
			args: args{
				value: "xtwone3four",
			},
			want: 24,
		}, {
			name: "happy_path_5",
			args: args{
				value: "4nineeightseven2",
			},
			want: 42,
		}, {
			name: "happy_path_6",
			args: args{
				value: "zoneight234",
			},
			want: 14,
		}, {
			name: "happy_path_7",
			args: args{
				value: "7pqrstsixteen",
			},
			want: 76,
		}, {
			name: "happy_path_8",
			args: args{
				value: "9three4nd1eightseven",
			},
			want: 97,
		}, {
			name: "happy_path_9",
			args: args{
				value: "6oneighthlf",
			},
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationValue(tt.args.value); got != tt.want {
				t.Errorf("getCalibrationValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

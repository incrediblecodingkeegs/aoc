package main

import "testing"

func Test_sumCalibrationValues(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy_path",
			args: args{
				values: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
			},
			want: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumCalibrationValues(tt.args.values); got != tt.want {
				t.Errorf("sumCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

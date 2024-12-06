package main

import (
	"testing"
)

func Test_performOperation(t *testing.T) {
	type args struct {
		operation string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "mul(11,8)__want_88",
			args: args{
				operation: "mul(11,8)",
			},
			want: 88,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := performOperation(tt.args.operation); got != tt.want {
				t.Errorf("performOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "testString__want_len(4)",
			args: args{
				file: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			wantLen: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFile(tt.args.file); len(got) != tt.wantLen {
				t.Errorf("parseFile() = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

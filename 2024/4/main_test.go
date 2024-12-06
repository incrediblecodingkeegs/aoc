package main

import (
	"reflect"
	"testing"
)

func Test_getCols(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple__want_true",
			args: args{
				rows: []string{
					"abc",
					"def",
					"ghi",
				},
			},
			want: []string{
				"adg",
				"beh",
				"cfi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCols(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDiags(t *testing.T) {
	type args struct {
		rows []string
		cols []string
	}
	tests := []struct {
		name          string
		args          args
		wantContains  []string
		wantContains1 []string
	}{
		{
			name: "simple__want_true",
			args: args{
				rows: []string{
					"abc",
					"def",
					"ghi",
				},
			},
			wantContains: []string{
				"aei",
				"dh",
				"bf",
			},
			wantContains1: []string{
				"ceg",
				"bd",
				"fh",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.cols = getCols(tt.args.rows)
			got, got1 := getDiags(tt.args.rows, tt.args.cols)
			if !contains(got, tt.wantContains) {
				t.Errorf("getDiags() got = %v, want %v", got, tt.wantContains)
			}
			if !contains(got1, tt.wantContains1) {
				t.Errorf("getDiags() got1 = %v, want %v", got1, tt.wantContains1)
			}
		})
	}
}

func contains(arr []string, targets []string) bool {
	matches := 0
	for _, target := range targets {
		for _, s := range arr {
			if s == target {
				matches++
				break
			}
		}
	}

	if matches == len(targets) {
		return true
	}
	return false
}

func Test_isXmas(t *testing.T) {
	type args struct {
		rows []string
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy_path__want_true",
			args: args{
				rows: []string{
					"SXS",
					"XAX",
					"MXM",
				},
				x: 1,
				y: 1,
			},
			want: true,
		},
		{
			name: "correct_corners_wrong_place__want_false",
			args: args{
				rows: []string{
					"MXS",
					"XAX",
					"SXM",
				},
				x: 1,
				y: 1,
			},
			want: false,
		},
		{
			name: "incorrect_corners__want_false",
			args: args{
				rows: []string{
					"XXS",
					"XAX",
					"SXM",
				},
				x: 1,
				y: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isXmas(tt.args.rows, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isXmas() = %v, want %v", got, tt.want)
			}
		})
	}
}

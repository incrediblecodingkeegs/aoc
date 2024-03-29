package main

import (
	"reflect"
	"testing"
)

func Test_isPartNumber(t *testing.T) {
	type args struct {
		schematic []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "corner_number_want_true",
			args: args{
				schematic: []string{
					"467.",
					"...*",
				},
			},
			want: true,
		}, {
			name: "middle_number_want_true",
			args: args{
				schematic: []string{
					"....+",
					".592.",
					".....",
				},
			},
			want: true,
		}, {
			name: "middle_number_want_false",
			args: args{
				schematic: []string{
					"....",
					".52.",
					"....",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPartNumber(tt.args.schematic); got != tt.want {
				t.Errorf("isPartNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPartNumbers(t *testing.T) {
	type args struct {
		schematic []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "all_are_part_numbers",
			args: args{
				schematic: []string{
					"......755.",
					"...$.*....",
					".664.598..",
				},
			},
			want: []int{
				755,
				664,
				598,
			},
		}, {
			name: "some_are_part_numbers",
			args: args{
				schematic: []string{
					"467..114.*",
					"...*....32",
					"..35..633.",
				},
			},
			want: []int{
				467,
				32,
				35,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPartNumbers(tt.args.schematic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

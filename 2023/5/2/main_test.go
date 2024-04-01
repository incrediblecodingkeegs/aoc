package main

import (
	"reflect"
	"testing"
)

func Test_convertValue(t *testing.T) {
	type args struct {
		value   int
		almanac []*Range
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "correct_values__want_success",
			args: args{
				value: 79,
				almanac: []*Range{
					getConversionRange("50 98 2"),
					getConversionRange("52 50 48"),
				},
			},
			want: 81,
		}, {
			name: "correct_values__want_success",
			args: args{
				value: 14,
				almanac: []*Range{
					getConversionRange("50 98 2"),
					getConversionRange("52 50 48"),
				},
			},
			want: 14,
		}, {
			name: "correct_values__want_success",
			args: args{
				value: 55,
				almanac: []*Range{
					getConversionRange("50 98 2"),
					getConversionRange("52 50 48"),
				},
			},
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertValue(tt.args.value, tt.args.almanac); got != tt.want {
				t.Errorf("convertValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConversionRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Range
	}{
		{
			name: "correct_values__want_true",
			args: args{
				s: "50 98 2",
			},
			want: &Range{
				start: 98,
				end:   99,
				delta: -48,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConversionRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConversionRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

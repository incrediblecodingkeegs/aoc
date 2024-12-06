package main

import (
	"reflect"
	"testing"
)

func Test_isCorrectOrder(t *testing.T) {

	exampleRuleList := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	rules := newRules(exampleRuleList)

	type args struct {
		order []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1__want_true",
			args: args{
				order: []int{75, 47, 61, 53, 29},
			},
			want: true,
		}, {
			name: "example2__want_true",
			args: args{
				order: []int{97, 61, 53, 29, 13},
			},
			want: true,
		}, {
			name: "example3__want_true",
			args: args{
				order: []int{75, 29, 13},
			},
			want: true,
		}, {
			name: "example4__want_false",
			args: args{
				order: []int{75, 97, 47, 61, 53},
			},
			want: false,
		}, {
			name: "example5__want_false",
			args: args{
				order: []int{61, 13, 29},
			},
			want: false,
		}, {
			name: "example6__want_false",
			args: args{
				order: []int{97, 13, 75, 29, 47},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrectOrder(rules, tt.args.order); got != tt.want {
				t.Errorf("isCorrectOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeCorrectOrder(t *testing.T) {
	exampleRuleList := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	rules := newRules(exampleRuleList)

	type args struct {
		order []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1__want_match",
			args: args{
				order: []int{75, 97, 47, 61, 53},
			},
			want: []int{97, 75, 47, 61, 53},
		}, {
			name: "example1__want_match",
			args: args{
				order: []int{61, 13, 29},
			},
			want: []int{61, 29, 13},
		}, {
			name: "example1__want_match",
			args: args{
				order: []int{97, 13, 75, 29, 47},
			},
			want: []int{97, 75, 47, 29, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeCorrectOrder(rules, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeCorrectOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

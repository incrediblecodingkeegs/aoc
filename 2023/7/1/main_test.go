package main

import (
	"reflect"
	"testing"
)

func TestNewHand(t *testing.T) {
	type args struct {
		handString string
	}
	tests := []struct {
		name string
		args args
		want *Hand
	}{
		{
			name: "correct_hand__want_success",
			args: args{
				handString: "32T3K 765",
			},
			want: &Hand{
				cards: "32T3K",
				bid:   765,
			},
		}, {
			name: "nil_hand__want_success",
			args: args{
				handString: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHand(tt.args.handString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareHandHighCard(t *testing.T) {
	type args struct {
		h1 *Hand
		h2 *Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "77888_>_77788__want_true",
			args: args{
				h1: NewHand("77888 0"),
				h2: NewHand("77788 0"),
			},
			want: true,
		}, {
			name: "33332_>_2AAAA__want_true",
			args: args{
				h1: NewHand("33332 0"),
				h2: NewHand("2AAAA 0"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareHandHighCard(tt.args.h1, tt.args.h2); got != tt.want {
				t.Errorf("compareHandHighCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_getHandStrength(t *testing.T) {
	type fields struct {
		cards string
		bid   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "five_of_kind__want_success",
			fields: fields{
				cards: "AAAAA",
				bid:   0,
			},
			want: 6,
		}, {
			name: "four_of_kind__want_success",
			fields: fields{
				cards: "AA8AA",
				bid:   0,
			},
			want: 5,
		}, {
			name: "full_house__want_success",
			fields: fields{
				cards: "23332",
				bid:   0,
			},
			want: 4,
		}, {
			name: "three_of_kind__want_success",
			fields: fields{
				cards: "TTT98",
				bid:   0,
			},
			want: 3,
		}, {
			name: "two_pair__want_success",
			fields: fields{
				cards: "23432",
				bid:   0,
			},
			want: 2,
		}, {
			name: "one_pair__want_success",
			fields: fields{
				cards: "A23A4",
				bid:   0,
			},
			want: 1,
		}, {
			name: "high_card__want_success",
			fields: fields{
				cards: "23456",
				bid:   0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hand{
				cards: tt.fields.cards,
				bid:   tt.fields.bid,
			}
			if got := h.getHandStrength(); got != tt.want {
				t.Errorf("getHandStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		handStrings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example__want_6440",
			args: struct {
				handStrings []string
			}{
				handStrings: []string{
					"32T3K 765",
					"T55J5 684",
					"KK677 28",
					"KTJJT 220",
					"QQQJA 483",
				},
			},
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.handStrings); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_getCardWinnings(t *testing.T) {
	type args struct {
		card *ScratchCard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "four_winning_numbers__expect_8",
			args: args{
				card: &ScratchCard{
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    map[int]bool{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
				},
			},
			want: 8,
		}, {
			name: "two_winning_numbers__expect_2",
			args: args{
				card: &ScratchCard{
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    map[int]bool{69: true, 82: true, 63: true, 72: true, 16: true, 21: true, 14: true, 1: true},
				},
			},
			want: 2,
		}, {
			name: "no_winning_numbers__expect_0",
			args: args{
				card: &ScratchCard{
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    map[int]bool{88: true, 30: true, 70: true, 12: true, 93: true, 22: true, 82: true, 36: true},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCardWinnings(tt.args.card); got != tt.want {
				t.Errorf("getCardWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createCard(t *testing.T) {
	type args struct {
		cardString string
	}
	tests := []struct {
		name string
		args args
		want *ScratchCard
	}{
		{
			name: "create_card__want_success",
			args: args{cardString: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"},
			want: &ScratchCard{
				id:             1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				gameNumbers:    map[int]bool{83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createCard(tt.args.cardString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

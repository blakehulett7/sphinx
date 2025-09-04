package main

import (
	"testing"
)

func FuzzBestFusion(f *testing.F) {
	max_card_id := 722

	f.Fuzz(func(t *testing.T, a, b, c, d, e int) {
		hand := []int{a, b, c, d, e}
		for _, card_id := range hand {
			if card_id > max_card_id {
				return
			}
		}

		app.BestFusion(hand, 0)
	})
}

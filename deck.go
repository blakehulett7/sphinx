package main

import (
	"fmt"
	"slices"
)

type Deck [40]int

func (d Deck) AllHandKeys() []HandKey {
	var (
		c1    = 0
		c2    = 1
		c3    = 2
		c4    = 3
		c5    = 4
		hands []HandKey
	)

	for c1 < len(d)-4 {
		for c2 < len(d)-3 {
			for c3 < len(d)-2 {
				for c4 < len(d)-1 {
					for c5 < len(d) {
						hand := HandKey{c1, c2, c3, c4, c5}
						hands = append(hands, hand)
						c5++
					}
					c4++
					c5 = c4 + 1
				}
				c3++
				c4 = c3 + 1
				c5 = c4 + 1
			}
			c2++
			c3 = c2 + 1
			c4 = c3 + 1
			c5 = c4 + 1
		}
		c1++
		c2 = c1 + 1
		c3 = c2 + 1
		c4 = c3 + 1
		c5 = c4 + 1
	}

	return hands
}

func (d Deck) DistinctHands() [][]int {
	var result [][]int
	return result
}

func (d Deck) GetHand(hand_key HandKey) []int {
	var hand []int
	for _, idx := range hand_key {
		hand = append(hand, d[idx])
	}
	return hand
}

func (d Deck) PrintHand(hand_key HandKey) {
	fmt.Printf("[%d %d %d %d %d]\n", d[hand_key[0]], d[hand_key[1]], d[hand_key[2]], d[hand_key[3]], d[hand_key[4]])
}

func ManualDistinctHands(hands []Hand) []Hand {
	var distinct_hands []Hand
	seen := make(map[[5]int]bool)
	for _, hand := range hands {
		slices.Sort(hand)
		var key [5]int

		for i, card_id := range hand {
			key[i] = card_id
		}

		if seen[key] {
			continue
		}

		seen[key] = true

		distinct_hands = append(distinct_hands, hand)
	}

	return distinct_hands
}

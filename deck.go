package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Deck [40]int

func AllHandKeys(d Deck) []HandKey {
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

func (app Bridge) EvaluateDeck(d Deck) {
	hand_keys := AllHandKeys(d)
	key_channel := make(chan HandKey)
	value_channel := make(chan MapEntry)
	done_channel := make(chan map[string]int)

	go KeyMapper(key_channel, value_channel, d, app)
	go SendKeys(key_channel, hand_keys)
	go WriteKey(value_channel, done_channel)

	hand_values := <-done_channel

	data, err := json.MarshalIndent(hand_values, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	filepath := "sample_deck.json"
	fmt.Println("Writing...")
	os.WriteFile(filepath, data, 0644)
	fmt.Println("Written")
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

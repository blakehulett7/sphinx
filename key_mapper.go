package main

import "fmt"

type MapEntry struct {
	Key   HandKey
	Value int
}

func KeyMapper(key_channel chan HandKey, value_channel chan MapEntry, deck Deck, app Bridge) {
	producer_done := make(chan bool)
	num_producer_cores := 10
	for range num_producer_cores {
		go MapKey(key_channel, value_channel, producer_done, deck, app)
	}

	for num_producer_cores > 0 {
		<-producer_done
		num_producer_cores--
	}

	close(value_channel)
}

func MapKey(key_channel chan HandKey, value_channel chan MapEntry, producer_done chan bool, deck Deck, app Bridge) {
	for {
		key, open := <-key_channel
		if !open {
			producer_done <- true
			return
		}

		initial_weight := 0
		hand := deck.GetHand(key)
		value := app.BestFusion(hand, initial_weight)

		result := MapEntry{key, value}
		value_channel <- result
	}
}

func WriteKey(value_channel chan MapEntry, done_channel chan map[HandKey]int) {
	result := make(map[HandKey]int)
	for {
		to_write, open := <-value_channel
		if !open {
			break
		}

		ColorPrint(Green, fmt.Sprintf("Saving key %v", to_write.Key))
		result[to_write.Key] = to_write.Value
	}

	done_channel <- result
}

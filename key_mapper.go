package main

type MapEntry struct {
	Key   HandKey
	Value int
}

func KeyMapper(key_channel chan HandKey, value_channel chan MapEntry, deck Deck, app Bridge) {
	num_producer_cores := 10
	for range num_producer_cores {
		go MapKey(key_channel, value_channel, deck, app)
	}
}

func MapKey(key_channel chan HandKey, value_channel chan MapEntry, deck Deck, app Bridge) {
	for {
		key := <-key_channel

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

		result[to_write.Key] = to_write.Value
	}

	done_channel <- result
}

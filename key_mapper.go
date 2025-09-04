package main

import "fmt"

type MapEntry struct {
	Key   HandKey
	Value int
}

func KeyMapper(key_channel chan HandKey, value_channel chan int, deck Deck, app Bridge) {
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

func MapKey(key_channel chan HandKey, value_channel chan int, producer_done chan bool, deck Deck, app Bridge) {
	for {
		key, open := <-key_channel
		if !open {
			producer_done <- true
			return
		}

		initial_weight := 0
		hand := deck.GetHand(key)
		value := app.BestFusion(hand, initial_weight)

		value_channel <- value
	}
}

func SendKeys(key_channel chan HandKey, keys []HandKey) {
	for _, key := range keys {
		key_channel <- key
	}
	close(key_channel)
}

func WriteKey(value_channel chan int, done_channel chan map[string]int) {
	result := make(map[string]int)
	for {
		to_write, open := <-value_channel
		if !open {
			break
		}

		_, exists := result[fmt.Sprint(to_write)]
		if !exists {
			result[fmt.Sprint(to_write)] = 1
			continue
		}

		result[fmt.Sprint(to_write)]++
	}

	done_channel <- result
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Dominus Iesus Christus")
	fmt.Println("----------------------")
	fmt.Println()

	app := Bridge{}

	ColorPrint(Blue, "Connecting to db...")
	db, err := gorm.Open(sqlite.Open("fmr.db"))
	if err != nil {
		ColorPrint(Red, fmt.Sprintf("- error connnecting to db: %v", err))
		ColorPrint(Red, "- returning early...")
	}
	app.Db = db
	ColorPrint(Blue, "Connection successful!")
	fmt.Println()

	r := rand.New(rand.NewSource(1))
	var deck Deck
	for i := range 40 {
		num := r.Intn(723)
		deck[i] = num
	}

	hand_keys := deck.AllHandKeys()

	key_channel := make(chan HandKey)
	value_channel := make(chan MapEntry)
	done_channel := make(chan map[HandKey]int)

	go KeyMapper(key_channel, value_channel, deck, app)
	go WriteKey(value_channel, done_channel)

	for i, key := range hand_keys {
		fmt.Println(i)
		key_channel <- key
	}

	close(key_channel)
	hand_values := <-done_channel

	data, err := json.MarshalIndent(hand_values, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath := "sample_deck.json"
	fmt.Println("Writing...")
	os.WriteFile(filepath, data, 0644)
	fmt.Println("Written")

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("Et Spiritus Sancti!")
}

package main

import (
	"fmt"
	"math/rand"

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

	hand_values := make(map[HandKey]int)
	initial_weight := 0
	for i, key := range hand_keys {
		fmt.Println(i)
		hand := deck.GetHand(key)
		value := app.BestFusion(hand, initial_weight)
		hand_values[key] = value
	}

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("Et Spiritus Sancti!")
}

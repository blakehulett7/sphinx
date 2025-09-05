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

	app.EvaluateDeck(deck)

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("Et Spiritus Sancti!")
}

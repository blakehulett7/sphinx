package main

import (
	"fmt"

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

	hand := []int{2, 9, 10, 4, 98}
	app.BestFusion(hand)
}

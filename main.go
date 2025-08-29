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

	ColorPrint(Blue, "Connecting to db...")
	db, err := gorm.Open(sqlite.Open("fmr.db"))
	if err != nil {
		ColorPrint(Red, fmt.Sprintf("- error connnecting to db: %v", err))
		ColorPrint(Red, "- returning early...")
	}

	key := [5]int{1, 2, 3, 4, 5}
}

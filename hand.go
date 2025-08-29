package main

import "fmt"

type Hand [5]int

func (app Bridge) GetBestFusion(hand Hand) int {
	res := []Fusion{}
	app.Db.Where("material1_id = ?", hand[0]).Find(&res)

	fmt.Println(res)

	return 1
}

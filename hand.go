package main

import "fmt"

type Hand [5]int

func (app Bridge) BestFusion(hand Hand) (int, error) {
	subject := hand[0]
	targets := hand[1:]

	app.PossibleFusions(subject, targets)

	subject = hand[1]
	targets = hand[2:]

	app.PossibleFusions(subject, targets)

	return 0, nil
}

func (app Bridge) PossibleFusions(subject int, targets []int) ([]int, error) {
	fmt.Println(subject, targets)

	fusions := []Fusion{}
	err := app.Db.Where("material1_id = ? AND material2_id IN ?", subject, targets).Find(&fusions).Error
	if err != nil {
		return []int{}, err
	}

	res := []int{}
	for _, fusion := range fusions {
		fmt.Println(fusion.ResultId)
		res = append(res, fusion.ResultId)
	}
	fmt.Println()

	return res, nil
}

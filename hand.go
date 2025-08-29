package main

import (
	"fmt"
	"slices"
)

type Hand [5]int

func (app Bridge) BestFusion(hand Hand) (int, error) {
	sorted := Sort(hand)

	possible_fusions := []int{}
	for i := 1; i < 5; i++ {
		subject := sorted[i-1]
		targets := sorted[i:]
		possible_fusions = append(possible_fusions, app.PossibleFusions(subject, targets)...)
	}
	fmt.Println(possible_fusions)

	return 0, nil
}

func (app Bridge) NestedFusions(fusion Fusion, targets []int) []int {
	nested_targets := []int{}
	for _, target := range targets {
		if fusion.Material2Id == target {
			continue
		}

		nested_targets = append(nested_targets, target)
	}

	return app.PossibleFusions(fusion.ResultId, nested_targets)
}

func (app Bridge) PossibleFusions(subject int, targets []int) []int {
	fmt.Println(subject, targets)

	fusions := []Fusion{}
	err := app.Db.Where("material1_id = ? AND material2_id IN ?", subject, targets).Find(&fusions).Error
	if err != nil {
		return []int{}
	}

	res := []int{}
	for _, fusion := range fusions {
		fmt.Println(fusion.ResultId)
		res = append(res, fusion.ResultId)
	}
	fmt.Println()

	return res
}

func Sort(hand Hand) [5]int {
	cards := hand[:]

	slices.Sort(cards)

	sorted := [5]int{}
	for i := range 5 {
		sorted[i] = cards[i]
	}

	return sorted
}

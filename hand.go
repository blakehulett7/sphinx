package main

import (
	"fmt"
	"slices"
)

type Hand [5]int

func (app Bridge) BestFusion(hand Hand) (int, error) {
	sorted := Sort(hand)

	possible_fusions := []Fusion{}
	for i := 1; i < 5; i++ {
		subject := sorted[i-1]
		targets := sorted[i:]
		possible_fusions = append(possible_fusions, app.PossibleFusions(subject, targets)...)
	}

	for _, fusion := range possible_fusions {
		app.NestedFusions(fusion, hand[:])
	}

	return 0, nil
}

func (app Bridge) NestedFusions(fusion Fusion, hand []int) []Fusion {
	subject := fusion.ResultId

	m1_found := false
	m2_found := false
	targets := []int{}
	for _, card := range hand {
		if !m1_found && fusion.Material1Id == card {
			m1_found = true
			continue
		}

		if !m2_found && fusion.Material2Id == card {
			m2_found = true
			continue
		}

		targets = append(targets, card)
	}

	app.PossibleFusions(subject, targets)
	return []Fusion{}
}

func (app Bridge) PossibleFusions(subject int, targets []int) []Fusion {
	fmt.Println(subject, targets)

	fusions := []Fusion{}
	err := app.Db.Where("material1_id = ? AND material2_id IN ?", subject, targets).Find(&fusions).Error
	if err != nil {
		return []Fusion{}
	}

	for _, fusion := range fusions {
		fmt.Println(fusion.ResultId)
	}
	fmt.Println()

	return fusions
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

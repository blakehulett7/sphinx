package main

import (
	"fmt"
	"slices"
)

type Hand []int

func (app Bridge) BestFusion(hand Hand, current_weight int) int {
	slices.Sort(hand)
	ColorPrint(Green, fmt.Sprintf("%v", hand))
	ColorPrint(Green, fmt.Sprintf("Current Weight: %v", current_weight))

	cards := []Card{}
	err := app.Db.Where("id IN ?", hand).Find(&cards).Error
	if err != nil {
		panic("invalid hand")
	}

	for _, card := range cards {
		fmt.Printf("Id: %d, Weight: %d\n", card.Id, card.Attack)
		if current_weight < card.Attack {
			current_weight = card.Attack
		}
	}
	fmt.Println()

	possible_fusions := []Fusion{}
	for i := 1; i < len(hand); i++ {
		subject := hand[i-1]
		targets := hand[i:]
		possible_fusions = append(possible_fusions, app.PossibleFusions(subject, targets, hand)...)
	}

	res := []int{}
	nested_hands := [][]int{}
	for _, fusion := range possible_fusions {
		// app.NestedFusions(fusion, hand)
		res = append(res, fusion.ResultId)
		nested_hands = append(nested_hands, CreateNestedTargets(fusion, hand))
	}
	fmt.Println(res)
	fmt.Println()

	var nested_weight int
	for _, nested := range nested_hands {
		nested_weight = app.BestFusion(nested, current_weight)
	}

	if nested_weight < current_weight {
		return current_weight
	}

	return nested_weight
}

func CreateNestedTargets(fusion Fusion, hand []int) []int {
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

	return append(targets, fusion.ResultId)
}

func (app Bridge) NestedFusions(fusion Fusion, hand []int) []Fusion {
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

	nested_hand := append(targets, fusion.ResultId)
	ColorPrint(Yellow, fmt.Sprintf("----Starting Nested Run %v----\n", nested_hand))
	//app.BestFusion(nested_hand)

	return []Fusion{}
}

func (app Bridge) PossibleFusions(subject int, targets []int, hand []int) []Fusion {
	fmt.Println(subject, targets)

	fusions := []Fusion{}
	err := app.Db.Where("material1_id = ? AND material2_id IN ?", subject, targets).Find(&fusions).Error
	if err != nil {
		return []Fusion{}
	}

	for _, fusion := range fusions {
		// nested_targets := CreateNestedTargets(fusion, hand)
		fmt.Println(fusion.Material1Id, "+", fusion.Material2Id, "=", fusion.ResultId)
		// fmt.Println(nested_targets)
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

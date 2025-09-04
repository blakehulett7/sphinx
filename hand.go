package main

import (
	"slices"
)

type Hand []int

type HandKey [5]int

func (h Hand) Key() [5]int {
	var key [5]int
	for i, id := range h {
		key[i] = id
	}
	return key
}

func (app Bridge) BestFusion(hand Hand, current_weight int) int {
	slices.Sort(hand)

	cards := []Card{}
	err := app.Db.Where("id IN ?", hand).Find(&cards).Error
	if err != nil {
		panic("invalid hand")
	}

	for _, card := range cards {
		if current_weight < card.Attack {
			current_weight = card.Attack
		}
	}

	possible_fusions := []Fusion{}
	for i := 1; i < len(hand); i++ {
		subject := hand[i-1]
		targets := hand[i:]
		possible_fusions = append(possible_fusions, app.PossibleFusions(subject, targets, hand)...)
	}

	nested_hands := [][]int{}
	for _, fusion := range possible_fusions {
		nested_hands = append(nested_hands, CreateNestedTargets(fusion, hand))
	}

	var nested_weight int
	for _, nested := range nested_hands {
		nested_weight = app.BestFusion(nested, current_weight)
	}

	if nested_weight > current_weight {
		return nested_weight
	}

	return current_weight
}

func CreateNestedTargets(fusion Fusion, hand []int) []int {
	m1_found := false
	m2_found := false
	targets := []int{}

	for _, card_id := range hand {
		if !m1_found && fusion.Material1Id == card_id {
			m1_found = true
			continue
		}

		if !m2_found && fusion.Material2Id == card_id {
			m2_found = true
			continue
		}

		targets = append(targets, card_id)
	}

	return append(targets, fusion.ResultId)
}

func (app Bridge) PossibleFusions(subject int, targets []int, hand []int) []Fusion {
	fusions := []Fusion{}
	err := app.Db.Where("material1_id = ? AND material2_id IN ?", subject, targets).Find(&fusions).Error
	if err != nil {
		return []Fusion{}
	}

	return fusions
}

package main

type Card struct {
	Id            int `json:"CardId"`
	CardName      string
	Description   string
	GuardianStar1 string
	GuardianStar2 string
	Level         int
	Type          string
	Attack        int
	Defense       int
	Attribute     string
	Password      string
	StarchipCost  int
}

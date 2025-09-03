package main

import "fmt"

type Fusion struct {
	Id          int `json:"FusionCardId"`
	Material1Id int
	Material1   Card `gorm:"foreignKey:Material1Id"`
	Material2Id int
	Material2   Card `gorm:"foreignKey:Material2Id"`
	ResultId    int
	Result      Card `gorm:"foreignKey:ResultId"`
}

func (f Fusion) String() string {
	return fmt.Sprintf("%d + %d = %d", f.Material1Id, f.Material2Id, f.ResultId)
}

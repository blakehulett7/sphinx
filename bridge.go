package main

import "gorm.io/gorm"

type Bridge struct {
	Db *gorm.DB
}

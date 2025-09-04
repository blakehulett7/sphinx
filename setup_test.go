package main

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var app Bridge

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("fmr.db"))
	if err != nil {
		return
	}
	app.Db = db

	os.Exit(m.Run())
}

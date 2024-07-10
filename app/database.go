package app

import (
	"Redikru-technical-test/helper"
	"database/sql"
	"fmt"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/redikru_technical_test")
	helper.PanicIfError(err)

	fmt.Println("TESTTTTT")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

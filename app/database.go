package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3307)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

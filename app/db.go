package app

import (
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/library_db?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

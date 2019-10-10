package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rudglazier/dbutils"
)

func main() {
	// Connect to DB

	db, err := sql.Open("sqlite3", "railapi.db")

	if err != nil {
		log.Println("Driver Creation failed!")
	}

	// Create Tables

	dbutils.Initialize(db)

}

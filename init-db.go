package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type DB struct {
	connection *sql.DB
}

func (db *DB) initDB() {
	tables := []string{`
		CREATE TABLE IF NOT EXISTS mem (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			time INTEGER,
			total INTEGER,
			available INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS cpu (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			time INTEGER,
			"index" INTEGER,
			percent INTEGER
		);
	`,
	}

	for _, table := range tables {
		_, err := db.connection.Exec(table)
		if err != nil {
			log.Fatal("couldn't create table", table, err)
		}
	}
}

func (db *DB) connect() {
	connection, err := sql.Open("sqlite", "file:vision.db")
	if err != nil {
		log.Fatal("couldn't open sqlite database", err)
	}
	db.connection = connection
	db.initDB()
}

func NewDB() *DB {
	db := new(DB)
	db.connect()
	return db
}

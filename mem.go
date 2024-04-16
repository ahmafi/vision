package main

import (
	"log"
	"time"
)

type Mem struct {
	total     uint64
	available uint64
}

func (db *DB) Mem(mem *Mem) {
	const sql string = `
		INSERT INTO mem (
			time,
			total,
			available
		) VALUES(?, ?, ?);
	`

	_, err := db.connection.Exec(sql, time.Now().Unix(), mem.total, mem.available)
	if err != nil {
		log.Println("couldn't insert to mem table", err)
		// TODO: send crash report
	}
}

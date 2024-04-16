package main

import (
	"log"
	"time"
)

type Cpu struct {
	time    time.Time
	index   int
	percent uint64
}

func (db *DB) Cpu(cpu *Cpu) {
	const sql string = `
		INSERT INTO cpu (
			time,
			"index",
			percent
		) VALUES(?, ?, ?);
	`

	_, err := db.connection.Exec(sql, time.Now().Unix(), cpu.index, cpu.percent)
	if err != nil {
		log.Println("couldn't insert to cpu table", err)
		// TODO: send crash report
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	db := NewDB()

	for {
		m, _ := mem.VirtualMemory()
		db.Mem(&Mem{
			total:     m.Total,
			available: m.Available,
		})

		c, _ := cpu.Percent(0, true)
		for i, usage := range c {
			fmt.Printf("cpu %v %f%%\n", i, usage)
			db.Cpu(&Cpu{
				index:   i,
				percent: uint64(usage * 1000),
			})
		}

		time.Sleep(time.Second * 3)
	}

}

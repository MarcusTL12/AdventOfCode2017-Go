package main

import (
	"aoc2017/days"
	"fmt"
	"os"
	"strconv"
	"time"
)

var funcs = [][2]func(){days.Day1}

func main() {
	day, day_err := strconv.ParseInt(os.Args[1], 10, 64)
	part, part_err := strconv.ParseInt(os.Args[2], 10, 64)

	if day_err != nil {
		panic(day_err)
	}

	if part_err != nil {
		panic(part_err)
	}

	t0 := time.Now()
	funcs[day-1][part-1]()
	t1 := time.Now()

	t := t1.Sub(t0)

	fmt.Printf("Took: %.6f s", t.Seconds())
}

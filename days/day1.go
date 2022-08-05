package days

import (
	"os"
)

var Day1 = [2]func(){part1, part2}

func part1() {
	inp, err := os.ReadFile("input/day1/input")
	if err != nil {
		panic(err)
	}

	ans := 0
	if inp[0] == inp[len(inp)-1] {
		ans = int(inp[0])
	}

	for i := 0; i < len(inp)-1; i++ {
		if inp[i] == inp[i+1] {
			ans += int(inp[i]) - '0'
		}
	}

	println(ans)
}

func part2() {
	inp, err := os.ReadFile("input/day1/input")
	if err != nil {
		panic(err)
	}

	ans := 0

	for i, j := 0, len(inp)/2; j < len(inp); i++ {
		if inp[i] == inp[j] {
			ans += 2 * (int(inp[i]) - '0')
		}
		j++
	}

	println(ans)
}

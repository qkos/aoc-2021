package main

import (
	"fmt"
	"github.com/qkos/aoc-2021/pkg"
	"strconv"
)

const inputPart1 = "part1.in"
const inputPart2 = "part1.in"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	lines, err := pkg.FileToLines(inputPart1)
	if err != nil {
		panic(err)
	}

	var last, c int

	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
		}

		if last == 0 {
			last = num
			continue
		}

		if num > last {
			c++
		}
		last = num
	}

	return c
}

func part2() int {
	lines, err := pkg.FileToLines(inputPart2)
	if err != nil {
		panic(err)
	}

	var last, c int

	for i := 0; i < len(lines)-2; i++ {
		var curr int
		for j := 0; j < 3; j++ {
			num, _ := strconv.Atoi(lines[i+j])
			curr += num
		}

		if last == 0 {
			last = curr
			continue
		}

		if curr > last {
			c++
		}
		last = curr
	}

	return c
}
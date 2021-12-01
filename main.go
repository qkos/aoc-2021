package main

import (
	"fmt"
	"github.com/qkos/aoc-2021/pkg"
	"strconv"
)

const inputPart1 = "test.in"
const inputPart2 = "part2.in"

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	lines, err := pkg.FileToLines(inputPart1)
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
		}

		// TODO: implement
		_ = num
	}

	return 0
}

func part2() int {
	lines, err := pkg.FileToLines(inputPart2)
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
		}

		// TODO: implement
		_ = num
	}

	return 0
}
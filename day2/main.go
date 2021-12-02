package main

import (
	"fmt"
	"github.com/qkos/aoc-2021/pkg"
	"strconv"
	"strings"
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

	var h, d int
	for i, line := range lines {
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
		}

		switch parts[0] {
		case "forward":
			h += num
			break
		case "down":
			d += num
			break
		case "up":
			d -= num
			break
		}

	}

	return h*d
}

func part2() int {
	lines, err := pkg.FileToLines(inputPart2)
	if err != nil {
		panic(err)
	}

	var a, h, d int
	for i, line := range lines {
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(fmt.Sprintf("line #%d is not a number [%s]", i, line))
		}

		switch parts[0] {
		case "forward":
			h += num
			d += a * num
			break
		case "down":
			a += num
			break
		case "up":
			a -= num
			break
		}

	}

	return h*d
}
package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"regexp"
	"strings"
)

var (
	mulRegex  = `mul\(\d{1,3},\d{1,3}\)`
	doRegex   = `do\(\)`
	dontRegex = `don't\(\)`
)

// 175700056
func part1(input string) int {
	r := _024.Must(regexp.Compile(mulRegex))

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			sum += calculateMul(match)
		}
	}
	return sum
}

// 71668682
func part2(input string) int {
	allRegex := mulRegex + "|" + doRegex + "|" + dontRegex
	r := _024.Must(regexp.Compile(allRegex))

	input = strings.ReplaceAll(input, "\n", "")
	sum := 0
	isDo := true

	matches := r.FindAllString(input, -1)
	for _, match := range matches {
		switch match {
		case "do()":
			isDo = true
		case "don't()":
			isDo = false
		default:
			if isDo {
				sum += calculateMul(match)
			}
		}
	}
	return sum
}

// ToDo ender try to extract numbers with regex just for fun
func calculateMul(s string) int {
	s, _ = strings.CutPrefix(s, "mul(")
	s, _ = strings.CutSuffix(s, ")")
	splitted := strings.Split(s, ",")
	return _024.MustParseInt(splitted[0]) * _024.MustParseInt(splitted[1])
}

func main() {
	filename := "inputs/day_03/input.txt"

	input := _024.ReadInput(filename)

	println(part1(input))
	println(part2(input))
}

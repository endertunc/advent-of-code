package main

import (
	"fmt"
	_024 "github.com/endertunc/advent-of-code/2024"
	"math"
	"slices"
	"strings"
)

// 21790168
func part2(left []int, right []int) int {
	occurrences := make(map[int]int)
	for _, n := range right {
		occurrences[n]++
	}
	sum := 0
	for _, n := range left {
		sum += n * occurrences[n]
	}
	return sum
}

// 1151792
func part1(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	return sum
}

func main() {
	filename := "inputs/day_01/input.txt"
	input := _024.ReadInput(filename)

	left, right := parseInput(input)

	println(part1(left, right))
	println(part2(left, right))
}

func parseInput(input string) ([]int, []int) {
	var left, right []int
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}
		left = append(left, _024.MustParseInt(fields[0]))
		right = append(right, _024.MustParseInt(fields[1]))
	}
	return left, right
}

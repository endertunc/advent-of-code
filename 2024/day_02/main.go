package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

func isSafe(report []int) bool {
	diff := report[0] - report[1]

	if diff == 0 {
		return false
	}

	if diff > 0 {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]
			if !(0 < diff && diff < 4) {
				return false
			}
		}
	} else {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]
			if !(-4 < diff && diff < 0) {
				return false
			}
		}
	}
	return true
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// 343
func part2(reports [][]int) int {
	safe := 0
	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				subReport := removeIndex(report, i)
				if isSafe(subReport) {
					safe++
					break
				}
			}
		}
	}
	return safe
}

// 279
func part1(reports [][]int) int {
	safe := 0
	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}
	return safe
}

func main() {
	filename := "inputs/day_02/input.txt"
	input := _024.ReadInput(filename)

	reports := parseInput(input)

	println(part1(reports))
	println(part2(reports))
}

func parseInput(input string) [][]int {
	var reports [][]int
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		report := make([]int, 0)
		for _, field := range fields {
			report = append(report, _024.MustParseInt(field))
		}
		reports = append(reports, report)
	}
	return reports
}

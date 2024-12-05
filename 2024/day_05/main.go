package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

// 5639
func part1(updates [][]string, orderingRules map[string]bool) int {
	sum := 0
	for _, update := range updates {
		ok := checkUpdate(update, orderingRules)
		if ok {
			sum += _024.MustParseInt(update[findMiddleIndex(update)])
		}
	}
	return sum
}

func checkUpdate(update []string, orderingRules map[string]bool) bool {
	isValid := true
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			_, ok := orderingRules[update[i]+"|"+update[j]]
			if !ok {
				isValid = false
				break
			}
		}
	}
	return isValid
}

// 5273
func part2(updates [][]string, orderingRules map[string]bool) int {
	sum := 0
	for _, update := range updates {
		updated, ok := correctUpdate(update, orderingRules)
		if ok {
			sum += _024.MustParseInt(updated[findMiddleIndex(updated)])
		}
	}
	return sum
}

func correctUpdate(update []string, orderingRules map[string]bool) ([]string, bool) {
	isCorrected := false
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			ij := update[i] + "|" + update[j]
			_, ok := orderingRules[ij]
			if !ok {
				isCorrected = true
				a := update[i]
				b := update[j]
				update[i] = b
				update[j] = a
			}
		}
	}
	return update, isCorrected
}

func findMiddleIndex(s []string) int {
	if len(s)%2 == 0 {
		return len(s)/2 - 1
	}
	return len(s) / 2
}

func main() {
	filename := "inputs/day_05/input.txt"
	input := _024.ReadInput(filename)

	updates, orderingRules := parseInput(input)

	println(part1(updates, orderingRules))
	println(part2(updates, orderingRules))
}

func parseInput(input string) ([][]string, map[string]bool) {
	updates := make([][]string, 0)
	orderingRules := make(map[string]bool)
	for _, line := range strings.Split(input, "\n") {
		if line == "" { // skip the separator line
			continue
		}
		if strings.Contains(line, "|") {
			orderingRules[line] = true
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	return updates, orderingRules
}

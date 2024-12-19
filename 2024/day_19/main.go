package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

// 330
func part1(patterns []string, designs []string) int {
	possibleDesigns := 0
	dpCache := make(map[string]bool)

	for _, design := range designs {
		if isDesignPossible(patterns, design, dpCache) {
			possibleDesigns++
		}
	}
	return possibleDesigns
}

func isDesignPossible(patterns []string, design string, dpCache map[string]bool) bool {
	possible, ok := dpCache[design]
	if ok {
		return possible
	}

	if design == "" {
		return true
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining, _ := strings.CutPrefix(design, pattern)
			if isDesignPossible(patterns, remaining, dpCache) {
				return true
			} else {
				// cache the impossible designs to avoid recalculating
				dpCache[design] = false
			}
		}
	}

	return false
}

// 950763269786650
func part2(patterns []string, designs []string) int {
	possibleWays := 0
	dpCache := make(map[string]int)
	for _, design := range designs {
		possibleWays += findPossibleWays(patterns, design, dpCache)
	}
	return possibleWays
}

// as always, we can use part 2 to solve the part 1, but I will leave it like this to show the difference and remind myself the progress
func findPossibleWays(patterns []string, design string, dpCache map[string]int) int {
	if count, ok := dpCache[design]; ok {
		return count
	}

	if design == "" {
		return 1
	}

	count := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining, _ := strings.CutPrefix(design, pattern)
			count += findPossibleWays(patterns, remaining, dpCache)
		}
	}
	dpCache[design] = count
	return count
}

func main() {
	filename := "inputs/day_19/input.txt"
	input := _024.ReadInput(filename)

	patterns, designs := parseInput(input)
	println(part1(patterns, designs))
	println(part2(patterns, designs))
}

func parseInput(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")
	patterns := strings.Split(lines[0], ", ")
	designs := lines[2:]

	return patterns, designs
}

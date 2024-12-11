package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"math"
	"strings"
)

// part1 - 183620
// part2 - 220377651399268
func solution(numbersMap map[int]int, times int) int {
	for i := 0; i < times; i++ {
		numbersMap = blink(numbersMap)
	}
	sum := 0
	for _, v := range numbersMap {
		sum += v
	}
	return sum
}

func blink(numbers map[int]int) map[int]int {
	newNumbers := make(map[int]int)
	for k, v := range numbers {
		if k == 0 {
			newNumbers[1] += v
			continue
		}
		digits := findNumberOfDigits(k)
		if digits%2 == 0 {
			// this could be cached in a map, but it's not worth it since it's very, very fast already
			left, right := splitIntInHalf(k, digits)
			newNumbers[left] += v
			newNumbers[right] += v
			continue
		}

		newNumbers[k*2024] += v
	}
	return newNumbers
}

func splitIntInHalf(n, digits int) (int, int) {
	p := int(math.Pow10(digits / 2))
	right := n % p
	left := n / p
	return left, right
}

func findNumberOfDigits(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func main() {
	filename := "inputs/day_11/input.txt"
	input := _024.ReadInput(filename)

	numbersMap := parseInput(input)

	println(solution(numbersMap, 25)) // part1
	println(solution(numbersMap, 75)) // part2
}

func parseInput(input string) map[int]int {
	numbersMap := make(map[int]int)
	for _, s := range strings.Fields(input) {
		numbersMap[_024.MustParseInt(s)] += 1
	}
	return numbersMap
}

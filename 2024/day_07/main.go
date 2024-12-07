package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"math"
	"strconv"
	"strings"
)

// 1260333054159
func part1(equations []Equation) int {
	totalSum := 0
	for _, eq := range equations {
		for _, s := range combinations(len(eq.Numbers), 2) {
			sum := calc2(eq.Numbers, s)
			if sum == eq.Target {
				totalSum += eq.Target
				break
			}
		}
	}
	return totalSum
}

// 162042343638683
func part2(equations []Equation) int {
	totalSum := 0
	for _, eq := range equations {
		for _, s := range combinations(len(eq.Numbers), 3) {
			sum := calc3(eq.Numbers, s)
			if sum == eq.Target {
				totalSum += eq.Target
				break
			}
		}
	}
	return totalSum
}

func calc2(numbers []int, binaryString string) int {
	res := numbers[0]
	for i, c := range binaryString {
		switch string(c) {
		case "0":
			res += numbers[i+1]
		case "1":
			res *= numbers[i+1]
		}
	}
	return res
}

func calc3(numbers []int, binaryString string) int {
	res := numbers[0]
	for i, c := range binaryString {
		switch string(c) {
		case "0":
			res += numbers[i+1]
		case "1":
			res *= numbers[i+1]
		case "2":
			next := numbers[i+1]
			pow10 := math.Pow10(len(strconv.Itoa(next)))
			res = (res * int(pow10)) + next

		}
	}
	return res
}

type Equation struct {
	Target  int
	Numbers []int
}

func main() {
	filename := "inputs/day_07/input.txt"
	input := _024.ReadInput(filename)

	equations := parseInput(input)

	println(part1(equations))
	println(part2(equations))
}

// combinations generate all possible combinations of a given base and length.
// this function could be made more efficient by accepting consumer function which each combination is passed to.
// that way, we can short-circuit generation once the desired goal (whatever it might be) is reached
func combinations(n int, base int) []string {
	total := int(math.Pow(float64(base), float64(n-1)))
	result := make([]string, total)

	for i := 0; i < total; i++ {
		binaryStr := strconv.FormatInt(int64(i), base)
		if len(binaryStr) < n-1 {
			binaryStr = strings.Repeat("0", n-1-len(binaryStr)) + binaryStr
		}
		result[i] = binaryStr
	}
	return result
}

func parseInput(input string) []Equation {
	equations := make([]Equation, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ":")
		numbers := make([]int, 0)
		for _, n := range strings.Fields(split[1]) {
			numbers = append(numbers, _024.MustParseInt(n))
		}
		equations = append(equations, Equation{Target: _024.MustParseInt(split[0]), Numbers: numbers})
	}
	return equations
}

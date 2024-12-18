package main

import (
	"fmt"
	_024 "github.com/endertunc/advent-of-code/2024"
	"log"
	"math"
	"strconv"
	"strings"
)

// 1,5,0,5,2,0,1,3,5
// 150520135
func part1(A, B, C int, program []int) string {
	i := 0
	output := ""
	for i < len(program) {
		opcode := program[i]
		operand := program[i+1]
		didJump := false

		switch opcode {
		case 0:
			A = A / denominator(A, B, C, operand)
		case 1:
			B = B ^ operand
		case 2:
			B = comboOperand(A, B, C, operand) % 8
		case 3:
			if A != 0 {
				i = operand
				didJump = true
			}
		case 4:
			B = B ^ C
		case 5:
			output += strconv.Itoa(comboOperand(A, B, C, operand) % 8)
		case 6:
			B = A / denominator(A, B, C, operand)
		case 7:
			C = A / denominator(A, B, C, operand)
		default:
			log.Fatalf("invalid opcode: %d\n", opcode)
		}

		if !didJump {
			i += 2
		}
	}
	return output
}

func denominator(A, B, C, operand int) int {
	return int(math.Pow(float64(2), float64(comboOperand(A, B, C, operand))))
}

func comboOperand(A, B, C, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		log.Fatalf("invalid operand: %d", operand)
		return -1
	}
}

func main() {
	filename := "inputs/day_17/input.txt"
	input := _024.ReadInput(filename)

	A, B, C, program := parseInput(input)
	println(part1(A, B, C, program))
}

func parseInput(input string) (int, int, int, []int) {
	lines := strings.Split(input, "\n")
	var A, B, C int
	var program string
	_024.Must(fmt.Sscanf(lines[0], "Register A: %d", &A))
	_024.Must(fmt.Sscanf(lines[1], "Register B: %d", &B))
	_024.Must(fmt.Sscanf(lines[2], "Register C: %d", &C))
	_024.Must(fmt.Sscanf(lines[4], "Program: %s", &program))

	instructions := make([]int, 0)
	for _, s := range strings.Split(program, ",") {
		instructions = append(instructions, _024.MustParseInt(s))
	}

	return A, B, C, instructions
}

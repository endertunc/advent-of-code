package main

import (
	"fmt"
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

type ClawMachine struct {
	A     struct{ X, Y int }
	B     struct{ X, Y int }
	Prize struct{ X, Y int }
}

const (
	PrizeCorrectionValue = 10000000000000
)

// 25629
func part1(clawMachines []ClawMachine) int {
	totalCost := 0
	for _, clawMachine := range clawMachines {
		cost := winPrize(clawMachine)
		if cost != -1 {
			totalCost += cost
		}
	}
	return totalCost
}

// 107487112929999
func part2(clawMachines []ClawMachine) int {
	totalCost := 0
	for _, cm := range clawMachines {
		cost := winPrizeWithCramersRule(cm)
		if cost != -1 {
			totalCost += cost
		}
	}
	return totalCost
}

// brute force and one could also optimize by finding min values for i and j,
// but as part2 points out, the ultimate solution is to use Cramer's rule
func winPrize(clawMachine ClawMachine) int {
	cost := -1
	for i := 100; i >= 0; i-- {
		for j := 0; j <= 100; j++ {
			x := clawMachine.A.X*i + clawMachine.B.X*j
			y := clawMachine.A.Y*i + clawMachine.B.Y*j
			if x == clawMachine.Prize.X && y == clawMachine.Prize.Y {
				// there is only one solution in input data,
				// so we don't look for a solution with lower cost
				return (i * 3) + j
			}
		}
	}
	return cost
}

/*
Cramer's rule

Given the system:
a1 * x + b1 * y = c1	94 * x + 22 * y = 8400
a2 * x + b2 * y = c2	34 * x + 67 * y = 5400

with
D = a1 * b2 - a2 * b1
Dx = c1 * b2 - c2 * b1
Dy = a1 * c2 - c1 * a2

Then the solution is:
x = Dx / D
y = Dy / D
*/
func winPrizeWithCramersRule(cm ClawMachine) int {
	cm.Prize.X += PrizeCorrectionValue
	cm.Prize.Y += PrizeCorrectionValue

	D := cm.A.X*cm.B.Y - cm.A.Y*cm.B.X
	Dx := cm.Prize.X*cm.B.Y - cm.Prize.Y*cm.B.X
	Dy := cm.A.X*cm.Prize.Y - cm.Prize.X*cm.A.Y
	x := Dx / D
	y := Dy / D

	// check if x and y are integers
	if D != 0 && Dx == x*D && Dy == y*D {
		return x*3 + y
	}
	return -1
}

func main() {
	filename := "inputs/day_13/input.txt"
	input := _024.ReadInput(filename)

	clawMachines := parseInput(input)

	println(part1(clawMachines))
	println(part2(clawMachines))
}

func parseInput(input string) []ClawMachine {
	clawMachine := make([]ClawMachine, 0)
	lines := strings.Split(input, "\n")

	var aX, aY, bX, bY, prizeX, prizeY int
	for i := 0; i < len(lines); i += 4 {
		_024.Must(fmt.Sscanf(lines[i+0], "Button A: X+%d, Y+%d", &aX, &aY))
		_024.Must(fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY))
		_024.Must(fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &prizeX, &prizeY))

		clawMachine = append(clawMachine, ClawMachine{
			A:     struct{ X, Y int }{X: aX, Y: aY},
			B:     struct{ X, Y int }{X: bX, Y: bY},
			Prize: struct{ X, Y int }{X: prizeX, Y: prizeY},
		})
	}

	return clawMachine
}

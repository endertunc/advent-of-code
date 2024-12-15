package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"log"
	"slices"
	"strings"
)

type Box struct {
	X _024.Point
	Y _024.Point
}

// 1487337
func part1(grid [][]string, instructions []int) int {
	robot := findRobot(grid)

	for _, instruction := range instructions {
		robot = movePartOne(robot, instruction, grid)
	}

	sum := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == "O" {
				sum += (100 * i) + j
			}
		}
	}

	return sum
}

// 1521952
func part2(grid [][]string, instructions []int) int {
	grid = transformGrid(grid)
	robot := findRobot(grid)

	for _, instruction := range instructions {
		robot = movePartTwo(robot, instruction, grid)
	}

	sum := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == "[" {
				sum += (100 * i) + j
			}
		}
	}

	return sum
}

func transformGrid(grid [][]string) [][]string {
	newGrid := make([][]string, 0)
	for _, row := range grid {
		newRow := make([]string, 0)
		for _, cell := range row {
			if cell == "#" {
				newRow = append(newRow, "#")
				newRow = append(newRow, "#")
			}
			if cell == "O" {
				newRow = append(newRow, "[")
				newRow = append(newRow, "]")
			}
			if cell == "." {
				newRow = append(newRow, ".")
				newRow = append(newRow, ".")
			}
			if cell == "@" {
				newRow = append(newRow, "@")
				newRow = append(newRow, ".")
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func findRobot(grid [][]string) _024.Point {
	robot := _024.Point{X: 0, Y: 0}
	for i, row := range grid {
		for j, cell := range row {
			if cell == "@" {
				return _024.Point{X: j, Y: i}
			}
		}
	}
	return robot
}

func movePartOne(robot _024.Point, direction int, grid [][]string) _024.Point {
	d, ok := _024.AllDirectionsMap[direction]
	if !ok {
		log.Fatalf("invalid direction: %d", direction)
	}
	boxes := make([]_024.Point, 0)
	p := robot
	for {
		temp := p.MoveDirection(direction)
		v := grid[temp.Y][temp.X]
		if v == "#" || v == "." {
			break
		}
		if v == "O" {
			boxes = append(boxes, temp)
		}
		p = temp
	}

	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		temp := box.Move(d)
		if grid[temp.Y][temp.X] == "#" {
			break
		} else {
			grid[box.Y][box.X] = "."
			grid[temp.Y][temp.X] = "O"
			box = temp
		}
	}

	return moveRobot(robot, d, grid)
}

func movePartTwo(robot _024.Point, direction int, grid [][]string) _024.Point {
	d, ok := _024.AllDirectionsMap[direction]
	if !ok {
		log.Fatalf("invalid direction: %d", direction)
	}

	boxes := findAllImpactedBoxes(robot, robot, direction, grid)

	// we need to move the boxes farthest from the robot first, and since we find the impacted boxes with DFS we need to order them here.
	// we could potentially use BFS instead of DFS
	if direction == _024.NORTH || direction == _024.SOUTH {
		slices.SortFunc(boxes, func(i, j Box) int {
			if direction == _024.NORTH {
				return j.Y.Y - i.Y.Y
			}
			if direction == _024.SOUTH {
				return i.Y.Y - j.Y.Y
			}
			return 0
		})
	}

	// move the boxes from the farthest to the closest
	for i := len(boxes) - 1; i >= 0; i-- {
		box := boxes[i]
		newX, newY := box.X.Move(d), box.Y.Move(d)
		if grid[newX.Y][newX.X] == "#" || grid[newY.Y][newY.X] == "#" {
			break
		} else {
			switchPointValues(grid, box.Y, newY)
			switchPointValues(grid, box.X, newX)
		}
	}

	// finally, move the robot
	return moveRobot(robot, d, grid)
}

func moveRobot(robot _024.Point, direction _024.Point, grid [][]string) _024.Point {
	temp := robot.Move(direction)
	if grid[temp.Y][temp.X] == "." {
		grid[robot.Y][robot.X] = "."
		grid[temp.Y][temp.X] = "@"
		robot = temp
	}
	return robot
}

// ToDo move this to grid.go
func switchPointValues(grid [][]string, old, new _024.Point) {
	v := grid[new.Y][new.X]
	grid[new.Y][new.X] = grid[old.Y][old.X]
	grid[old.Y][old.X] = v
}

var directionToToken = map[int]string{
	_024.WEST: "]",
	_024.EAST: "[",
}

func findAllImpactedBoxes(pX, pY _024.Point, direction int, grid [][]string) []Box {
	boxes := make([]Box, 0)

	switch direction {
	case _024.EAST, _024.WEST:
		for {
			next := pX.Move(_024.AllDirectionsMap[direction])
			if grid[next.Y][next.X] == "#" || grid[next.Y][next.X] == "." {
				return boxes
			}

			if grid[next.Y][next.X] == directionToToken[direction] {
				nextY := next.Move(_024.AllDirectionsMap[direction])
				boxes = append(boxes, Box{X: next, Y: nextY})
				pX = nextY
			}
		}
	case _024.NORTH, _024.SOUTH:
		seenBoxes := make(map[Box]bool)
		r, ok := findAllImpactedBoxesNorthOrSouth(pX, pY, direction, seenBoxes, grid)
		if ok {
			boxes = append(boxes, r...)
		}
	default:
		log.Fatalf("invalid direction: %d", direction)
	}

	return boxes
}

func findAllImpactedBoxesNorthOrSouth(pX, pY _024.Point, direction int, seenBoxes map[Box]bool, grid [][]string) ([]Box, bool) {

	if direction == _024.NORTH || direction == _024.SOUTH {
		boxes := make([]Box, 0)
		boxesX, canMoveX := findBoxInDirection(pX, direction, seenBoxes, grid)
		boxesY, canMoveY := findBoxInDirection(pY, direction, seenBoxes, grid)

		if !canMoveX || !canMoveY {
			return boxes, false
		}

		boxes = append(boxes, boxesX...)
		boxes = append(boxes, boxesY...)

		for _, box := range boxes {
			r, ok := findAllImpactedBoxesNorthOrSouth(box.X, box.Y, direction, seenBoxes, grid)
			if !ok {
				return []Box{}, false
			}
			boxes = append(boxes, r...)
		}
		return boxes, true
	} else {
		log.Fatalf("invalid direction: %d", direction)
		return []Box{}, true
	}
}

func findBoxInDirection(p _024.Point, direction int, seenBoxes map[Box]bool, grid [][]string) ([]Box, bool) {
	boxes := make([]Box, 0)
	next := p.MoveDirection(direction)
	v := grid[next.Y][next.X]

	if v == "#" {
		return boxes, false
	}

	if v == "[" {
		east := _024.AllDirectionsMap[_024.EAST]
		pair := next.Move(east)
		box := Box{next, pair}
		_, ok := seenBoxes[box]
		if !ok {
			boxes = append(boxes, box)
			seenBoxes[box] = true
		}
	}
	if v == "]" {
		west := _024.AllDirectionsMap[_024.WEST]
		pair := next.Move(west)
		box := Box{pair, next}
		_, ok := seenBoxes[box]
		if !ok {
			boxes = append(boxes, box)
			seenBoxes[box] = true
		}
	}
	return boxes, true
}

func main() {
	filename := "inputs/day_15/input.txt"
	input := _024.ReadInput(filename)

	grid1, instructions := parseInput(input)
	grid2, _ := parseInput(input)

	println(part1(grid1, instructions))
	println(part2(grid2, instructions))
}

func parseInput(input string) ([][]string, []int) {
	grid := make([][]string, 0)
	instructions := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			grid = append(grid, strings.Split(line, ""))
		}

		for _, c := range strings.Split(line, "") {
			switch c {
			case "^":
				instructions = append(instructions, _024.NORTH)
			case ">":
				instructions = append(instructions, _024.EAST)
			case "v":
				instructions = append(instructions, _024.SOUTH)
			case "<":
				instructions = append(instructions, _024.WEST)
			}
		}
	}
	return grid, instructions
}

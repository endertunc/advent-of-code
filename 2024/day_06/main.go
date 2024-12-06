package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"log"
	"strings"
)

// ToDo @ender I should move this common grid functions to grid.go
const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type visitedObstacle struct {
	x, y, direction int
}

func moveInDirection(p _024.Point, dir int) _024.Point {
	switch dir {
	case NORTH:
		return _024.Point{X: p.X, Y: p.Y - 1}

	case EAST:
		return _024.Point{X: p.X + 1, Y: p.Y}

	case SOUTH:
		return _024.Point{X: p.X, Y: p.Y + 1}

	case WEST:
		return _024.Point{X: p.X - 1, Y: p.Y}
	}
	return _024.Point{}
}

func turnRight(dir int) int {
	return (dir + 1) % 4
}

// 5551
func part1(grid [][]string, startingPoint _024.Point, startingDirection int) int {
	currentPoint := startingPoint
	currentDirection := startingDirection

	visited := make(map[_024.Point]bool)
	for {
		visited[currentPoint] = true
		nextPoint := moveInDirection(currentPoint, currentDirection)
		if nextPoint.X < 0 || nextPoint.X >= len(grid[0]) || nextPoint.Y < 0 || nextPoint.Y >= len(grid) {
			break
		}

		s := grid[nextPoint.Y][nextPoint.X]
		if s == "#" {
			currentDirection = turnRight(currentDirection)
			currentPoint = moveInDirection(currentPoint, currentDirection)
		} else {
			currentPoint = nextPoint
		}
		visited[currentPoint] = true

		// safety check that we should never be on an obstacle
		if grid[currentPoint.Y][currentPoint.X] == "#" {
			log.Fatalf("passing through obstacle at %v, %v", currentPoint.X, currentPoint.Y)
		}
	}
	return len(visited)
}

// 1939
func part2(grid [][]string, startingPoint _024.Point, startingDirection int) int {
	currentPoint := startingPoint
	currentDirection := startingDirection

	count := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == "#" || (currentPoint.X == j && currentPoint.Y == i) {
				continue
			}
			loop := checkLoop(grid, i, j, currentPoint, currentDirection)
			if loop {
				count++
			}
		}
	}

	return count
}

func checkLoop(grid [][]string, i, j int, currentPoint _024.Point, currentDirection int) bool {
	loop := false
	visitedObstacles := make(map[visitedObstacle]bool)
	for {
		if loop {
			break
		}
		nextPoint := moveInDirection(currentPoint, currentDirection)
		if nextPoint.X < 0 || nextPoint.X >= len(grid[0]) || nextPoint.Y < 0 || nextPoint.Y >= len(grid) {
			break
		}

		// it's possible that next point after a turn is also an obstacle, so we might need to turn multiple times
		for {
			maybeObstacle := grid[nextPoint.Y][nextPoint.X]
			isImaginaryObstacle := nextPoint.X == j && nextPoint.Y == i
			if maybeObstacle == "#" || isImaginaryObstacle {
				_, ok := visitedObstacles[visitedObstacle{nextPoint.X, nextPoint.Y, currentDirection}]
				if ok {
					//log.Printf("Loop detected at %v, %v, %v", currentPoint.x, currentPoint.y, currentDirection)
					loop = true
					break
				}
				visitedObstacles[visitedObstacle{nextPoint.X, nextPoint.Y, currentDirection}] = true
				currentDirection = turnRight(currentDirection)
				nextPoint = moveInDirection(currentPoint, currentDirection)
			} else {
				currentPoint = nextPoint
				break
			}
		}

		// safety check that we should never be on an obstacle
		if grid[currentPoint.Y][currentPoint.X] == "#" {
			log.Fatalf("passing through obstacle at %v, %v", currentPoint.X, currentPoint.Y)
		}
	}
	return loop
}

func main() {
	filename := "inputs/day_06/input.txt"
	input := _024.ReadInput(filename)

	grid, startingPoint, startingDirection := parseInput(input)

	println(part1(grid, startingPoint, startingDirection))
	println(part2(grid, startingPoint, startingDirection))
}

func parseInput(input string) ([][]string, _024.Point, int) {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, cell := range line {
			row = append(row, string(cell))
		}
		grid = append(grid, row)
	}

	var startingPoint _024.Point
	var startingDirection int
	for i, row := range grid {
		for j, cell := range row {

			switch cell {
			case "^":
				startingPoint = _024.Point{X: j, Y: i}
				startingDirection = NORTH
				break
			case ">":
				startingPoint = _024.Point{X: j, Y: i}
				startingDirection = EAST
				break
			case "v":
				startingPoint = _024.Point{X: j, Y: i}
				startingDirection = SOUTH
				break
			case "<":
				startingPoint = _024.Point{X: j, Y: i}
				startingDirection = WEST
				break
			}
		}
	}
	return grid, startingPoint, startingDirection
}

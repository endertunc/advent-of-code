package _024

import (
	"fmt"
	"log"
	"slices"
)

// Directions in clockwise
var Directions = [][]int{
	{0, -1},  // up
	{1, -1},  // up-right
	{1, 0},   // right
	{1, 1},   // down-right
	{0, 1},   // down
	{-1, 1},  // down-left
	{-1, 0},  // left
	{-1, -1}, // up-left
}

const (
	NORTH = iota
	NORTH_EAST
	EAST
	SOUTH_EAST
	SOUTH
	SOUTH_WEST
	WEST
	NORTH_WEST
)

var AllDirectionsMap = map[int]struct{ X, Y int }{
	NORTH:      {0, -1},
	NORTH_EAST: {1, -1},
	EAST:       {1, 0},
	SOUTH_EAST: {1, 1},
	SOUTH:      {0, 1},
	SOUTH_WEST: {-1, 1},
	WEST:       {-1, 0},
	NORTH_WEST: {-1, -1},
}

var NonDiagonallyDirectionsMap = map[int]struct{ X, Y int }{
	NORTH: {0, -1},
	EAST:  {1, 0},
	SOUTH: {0, 1},
	WEST:  {-1, 0},
}

type Point struct {
	X, Y int
}

func (p Point) IsValid(maxX, maxY int) bool {
	return p.X >= 0 && p.X < maxX && p.Y >= 0 && p.Y < maxY
}

func IsPointValid[T any](p Point, grid [][]T) bool {
	maxY, maxX := len(grid), len(grid[0])
	return p.IsValid(maxX, maxY)
}

func (p Point) DistanceTo(other Point) Point {
	return Point{p.X - other.X, +p.Y - other.Y}
}

func (p Point) Move(distance Point) Point {
	return Point{X: p.X + distance.X, Y: p.Y + distance.Y}
}

func (p Point) MoveN(distance Point, n int) Point {
	return Point{X: p.X + n*distance.X, Y: p.Y + n*distance.Y}
}

func (p Point) MoveDirection(direction int) Point {
	d, ok := AllDirectionsMap[direction]
	if !ok {
		log.Fatalf("invalid direction: %d", direction)
	}
	return Point{X: p.X + 1*d.X, Y: p.Y + 1*d.Y}
}

func FindNonDiagonalValidPoints[T any](point Point, grid [][]T) []Point {
	var neighbours []Point
	for direction, _ := range NonDiagonallyDirectionsMap {
		nextPoint := point.MoveDirection(direction)
		if nextPoint.IsValid(len(grid[0]), len(grid)) {
			neighbours = append(neighbours, nextPoint)
		}
	}
	return neighbours
}

// - - - - - - - - - - - - - - - Debug Helpers - - - - - - - - - - - - - - -
const (
	Reset = "\033[0m"
	Green = "\033[32m"
)

func prettyPrintGrid(grid [][]string, points []Point) {
	for i, row := range grid {
		for j, val := range row {
			point := Point{X: j, Y: i}
			if ok := slices.Contains(points, point); ok {
				fmt.Printf("%s%s%s ", Green, val, Reset)
			} else {
				fmt.Printf("%s ", val)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

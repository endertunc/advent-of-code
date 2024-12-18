package main

import (
	"cmp"
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	_024 "github.com/endertunc/advent-of-code/2024"
	"slices"
	"strings"
)

type Path struct {
	Point _024.Point
	Score int
}

func byPriority(a, b interface{}) int {
	return cmp.Compare(a.(Path).Score, b.(Path).Score)
}

// simplified version of day 16
// 372
func part1(grid [][]string, start _024.Point, end _024.Point) int {
	paths := pq.NewWith(byPriority) // also called OpenSet in A* algorithms
	paths.Enqueue(Path{
		Point: start,
		Score: 0,
	})
	visitedPaths := make(map[_024.Point]int) // also called ClosedSet in A* algorithms

	for paths.Size() > 0 {
		item, _ := paths.Dequeue()
		path := item.(Path)

		// since we are doing a BFS, the first path we find is the shortest
		if path.Point == end {
			return path.Score
		}

		// if we have already visited this point with a better score, skip
		_, ok := visitedPaths[path.Point]
		if ok {
			continue
		}
		visitedPaths[path.Point] = path.Score

		for direction, _ := range _024.OrthogonalDirectionsMap {
			nextPoint := path.Point.MoveDirection(direction)
			if _024.IsPointValid(nextPoint, grid) && grid[nextPoint.Y][nextPoint.X] != "#" {
				paths.Enqueue(Path{Point: nextPoint, Score: path.Score + 1})
			}
		}
	}
	return -1
}

// 25,6
func part2(grid [][]string, start _024.Point, end _024.Point, allCorruptedPoints []_024.Point, initialFallenBytes int) string {
	// we start i from initialFallenBytes since we know from part 1 that there is a path.
	for i := initialFallenBytes; i < len(allCorruptedPoints); i++ {
		nextCorruptedPoint := allCorruptedPoints[i]
		grid[nextCorruptedPoint.Y][nextCorruptedPoint.X] = "#"
		if part1(grid, start, end) == -1 {
			return fmt.Sprintf("%d,%d", allCorruptedPoints[i].X, allCorruptedPoints[i].Y)
		}
	}
	return ""
}

func main() {
	//filename, gridDimension, initialFallenBytes := "test.txt", 7, 12
	filename, gridDimension, initialFallenBytes := "input.txt", 71, 1024

	input := _024.ReadInput("inputs/day_18/" + filename)

	// I realized a bit late, but actually we don't need a grid.
	// all we need is the corrupted points, dimensions - that's all.
	grid, allCorruptedPoints := parseInput(input, gridDimension, initialFallenBytes)

	start := _024.Point{X: 0, Y: 0}
	end := _024.Point{X: gridDimension - 1, Y: gridDimension - 1}

	println(part1(grid, start, end))
	println(part2(grid, start, end, allCorruptedPoints, initialFallenBytes))
}

func parseInput(input string, gridDimension, initialFallenBytes int) ([][]string, []_024.Point) {
	grid := make([][]string, gridDimension)
	var X, Y int
	allCorruptedPoints := make([]_024.Point, 0)

	for _, line := range strings.Split(input, "\n") {
		_024.Must(fmt.Sscanf(line, "%d,%d", &X, &Y))
		allCorruptedPoints = append(allCorruptedPoints, _024.Point{X: X, Y: Y})
	}

	firstFallenBytes := allCorruptedPoints[:initialFallenBytes]
	for i := 0; i < gridDimension; i++ {
		grid[i] = make([]string, gridDimension)
		for j := 0; j < gridDimension; j++ {
			if slices.Contains(firstFallenBytes, _024.Point{X: j, Y: i}) {
				grid[i][j] = "#"
			} else {
				grid[i][j] = "."
			}
		}
	}
	return grid, allCorruptedPoints
}

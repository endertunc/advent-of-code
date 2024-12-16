package main

import (
	"cmp"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	_024 "github.com/endertunc/advent-of-code/2024"
	"slices"
	"strings"
)

type PathPoint struct {
	Point     _024.Point
	Direction int
}

type Path struct {
	Point      _024.Point
	Direction  int
	Score      int
	PathPoints []PathPoint // only used in part 2
}

func byPriority(a, b interface{}) int {
	return cmp.Compare(a.(Path).Score, b.(Path).Score)
}

// 105508
func part1(grid [][]string, start _024.Point, end _024.Point) int {
	paths := pq.NewWith(byPriority) // also called OpenSet in A* algorithms
	paths.Enqueue(Path{
		Point:     start,
		Direction: _024.EAST,
		Score:     0,
	})
	visitedPaths := make(map[PathPoint]int) // also called ClosedSet in A* algorithms

	for paths.Size() > 0 {
		item, _ := paths.Dequeue()
		path := item.(Path)

		// since we are doing a BFS, the first path we find is the shortest
		if path.Point == end {
			return path.Score
		}

		pathPoint := PathPoint{
			Point:     path.Point,
			Direction: path.Direction,
		}

		// if we have already visited this point with a better score, skip
		score, ok := visitedPaths[pathPoint]
		if ok && score < path.Score {
			continue
		}
		visitedPaths[pathPoint] = path.Score

		directionCost := map[int]int{
			path.Direction: 1, // moving on the same direction cost 1
			_024.TurnClockwiseOrthogonal(path.Direction):        1001, // turn and move cost 1000 + 1
			_024.TurnCounterClockwiseOrthogonal(path.Direction): 1001, // turn and move cost 1000 + 1
		}

		for direction, cost := range directionCost {
			nextPoint := path.Point.MoveDirection(path.Direction)
			if grid[nextPoint.Y][nextPoint.X] != "#" {
				paths.Enqueue(Path{Point: nextPoint, Direction: direction, Score: path.Score + cost})
			}
		}
	}
	return -1
}

// part1 and part2 can very well be combined into a single function, but I wanted to keep them separate for clarity.
// for example, in part1 we don't need to keep track of the path points, but in part2 we do, etc.
// 548
func part2(grid [][]string, start _024.Point, end _024.Point, bestPathScore int) int {
	paths := pq.NewWith(byPriority) // also called OpenSet in A* algorithms
	initialPathPoint := PathPoint{Point: start, Direction: _024.EAST}
	paths.Enqueue(Path{
		Point:      start,
		Direction:  _024.EAST,
		Score:      0,
		PathPoints: []PathPoint{initialPathPoint},
	})

	visitedPaths := make(map[PathPoint]int) // also called ClosedSet in A* algorithms
	bestPaths := make([]Path, 0)

	for paths.Size() > 0 {
		item, _ := paths.Dequeue()
		path := item.(Path)

		if path.Point == end && path.Score == bestPathScore {
			bestPaths = append(bestPaths, path)
			continue
		}

		pathPoint := PathPoint{
			Point:     path.Point,
			Direction: path.Direction,
		}

		// if we have already visited this point with a better score, skip
		score, ok := visitedPaths[pathPoint]
		if ok && score < path.Score {
			continue
		}
		visitedPaths[pathPoint] = path.Score

		directionCost := map[int]int{
			path.Direction: 1, // moving on the same direction cost 1
			_024.TurnClockwiseOrthogonal(path.Direction):        1001, // turn and move cost 1000 + 1
			_024.TurnCounterClockwiseOrthogonal(path.Direction): 1001, // turn and move cost 1000 + 1
		}
		for direction, cost := range directionCost {
			nextPoint := path.Point.MoveDirection(direction)
			if grid[nextPoint.Y][nextPoint.X] != "#" {
				paths.Enqueue(Path{
					Point:     nextPoint,
					Direction: direction,
					Score:     path.Score + cost,
					PathPoints: append(slices.Clone(path.PathPoints), PathPoint{
						Point:     nextPoint,
						Direction: direction,
					}),
				})
			}
		}
	}

	uniquePoints := make(map[_024.Point]bool)
	for _, bestPath := range bestPaths {
		for _, pathPoint := range bestPath.PathPoints {
			uniquePoints[pathPoint.Point] = true
		}
	}

	return len(uniquePoints)
}

func main() {
	filename := "inputs/day_16/input.txt"
	input := _024.ReadInput(filename)

	start, end, grid := parseInput(input)

	bestPathScore := part1(grid, start, end)
	println(bestPathScore)
	println(part2(grid, start, end, bestPathScore))
}

func parseInput(input string) (_024.Point, _024.Point, [][]string) {
	grid := make([][]string, 0)
	start, end := _024.Point{}, _024.Point{}
	for i, line := range strings.Split(input, "\n") {
		row := strings.Split(line, "")
		grid = append(grid, row)
		for j, v := range row {
			if v == "S" {
				start = _024.Point{X: j, Y: i}
			}
			if v == "E" {
				end = _024.Point{X: j, Y: i}
			}
		}
	}
	return start, end, grid
}

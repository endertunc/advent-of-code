package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

// 1396562
func part1(grid [][]string) int {
	foundAreas := findAllAreas(grid)
	perimeters := calculatePerimeter(grid)

	total := 0
	for _, area := range foundAreas {
		sum := 0
		for _, p := range area {
			sum += perimeters[p.Y][p.X]
		}
		sum *= len(area)
		total += sum
	}

	return total
}

// 844132
func part2(grid [][]string) int {
	foundAreas := findAllAreas(grid)
	total := 0
	for _, area := range foundAreas {
		total += len(area) * countCorners(area, grid)
	}
	return total
}

// To reduce if else statements, I created a map to hold the orthogonal directions and their pairs
var orthogonalDirectionsPair = map[int][]int{
	_024.SOUTH_EAST: {_024.EAST, _024.SOUTH},
	_024.SOUTH_WEST: {_024.WEST, _024.SOUTH},
	_024.NORTH_EAST: {_024.EAST, _024.NORTH},
	_024.NORTH_WEST: {_024.WEST, _024.NORTH},
}

func countCorners(points []_024.Point, grid [][]string) int {
	corners := 0
	for _, p := range points {
		for direction, sides := range orthogonalDirectionsPair {
			orthogonalRef := p.Move(_024.AllDirectionsMap[direction])
			side1 := p.Move(_024.AllDirectionsMap[sides[0]])
			side2 := p.Move(_024.AllDirectionsMap[sides[1]])

			side1Match := isMatch(p, side1, grid)
			side2Match := isMatch(p, side2, grid)
			// if both sides are not matching, then the p is an outer corner

			if !side1Match && !side2Match {
				corners++
			}

			orthogonalMatch := isMatch(p, orthogonalRef, grid)
			// if both sides are matching and the orthogonalRef direction is not matching,
			// then the p is an inner corner
			if side1Match && side2Match && !orthogonalMatch {
				corners++
			}
		}
	}
	return corners
}

func isMatch(p1, p2 _024.Point, grid [][]string) bool {
	if _024.IsPointValid(p1, grid) && _024.IsPointValid(p2, grid) {
		return grid[p1.Y][p1.X] == grid[p2.Y][p2.X]
	}
	return false
}

func findAllAreas(grid [][]string) [][]_024.Point {
	seen := make(map[_024.Point]bool)
	foundAreas := make([][]_024.Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			p := _024.Point{X: j, Y: i}
			_, ok := seen[p]
			if !ok {
				seen[p] = true
				area := make(map[_024.Point]bool)
				findArea(p, grid, area)
				points := make([]_024.Point, 0)
				for point, _ := range area {
					seen[point] = true
					points = append(points, point)
				}
				foundAreas = append(foundAreas, points)
			}
		}
	}
	return foundAreas
}

func findArea(p _024.Point, grid [][]string, areas map[_024.Point]bool) {
	neighbors := _024.FindOrthogonalValidPoints(p, grid)
	areas[p] = true
	for _, neighbor := range neighbors {
		neighborValue := grid[neighbor.Y][neighbor.X]
		currentValue := grid[p.Y][p.X]
		_, ok := areas[neighbor]
		if !ok {
			if neighborValue == currentValue {
				areas[neighbor] = true
				findArea(neighbor, grid, areas)
			}
		}
	}
}

func calculatePerimeter(grid [][]string) [][]int {
	perimeters := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		perimeters[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			p := _024.Point{X: j, Y: i}
			v := grid[i][j]
			d := 0
			neighbors := _024.FindOrthogonalValidPoints(p, grid)
			d += 4 - len(neighbors)
			for _, n := range neighbors {
				if grid[n.Y][n.X] != v {
					d++
				}
			}
			perimeters[i][j] = d
		}
	}
	return perimeters
}

func main() {
	filename := "inputs/day_12/input.txt"
	input := _024.ReadInput(filename)

	grid := parseInput(input)

	println(part1(grid))
	println(part2(grid))
}

func parseInput(input string) [][]string {
	grid := make([][]string, 0)
	for _, s := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(s, ""))
	}
	return grid
}

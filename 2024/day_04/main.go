package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

// (i,j), (i,j+1), (i,j+2), (i,j+3),
// (i,j), (i,j-1), (i,j-2), (i,j-3),
// (i,j), (i+1,j), (i+2,j), (i+3,j)
// (i,j), (i-1,j), (i-2,j), (i-3,j)
// (i,j), (i+1,j+1), (i+3,j+2), (i+3,j+3)
// (i,j), (i-1,j-1), (i-2,j-2), (i-3,j-3)
// (i,j), (i+1,j-1), (i+2,j-2), (i+3,j-3)
// (i,j), (i-1,j+1), (i-2,j+2), (i-3,j+3)
func checkNeighborsForXmas(grid [][]string, x, y, maxX, maxY int) int {
	if grid[x][y] != "X" {
		return 0
	}
	words := make([]string, 0)
	for _, direction := range _024.Directions {
		dx, dy := direction[0], direction[1]
		w := ""
		for j := 0; j < 4; j++ {
			if (x+dx*j) > maxX || (y+dy*j) > maxY || (x+dx*j) < 0 || (y+dy*j) < 0 {
				break
			}
			w += grid[x+dx*j][y+dy*j]
		}
		words = append(words, w)
	}

	xmasCount := 0
	for _, w := range words {
		if w == "XMAS" {
			xmasCount++
		}
	}
	return xmasCount

}

// 2297
func part1(grid [][]string) int {
	totalXmas := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			c := checkNeighborsForXmas(grid, i, j, len(grid)-1, len(grid[i])-1)
			totalXmas += c
		}
	}
	return totalXmas
}

var startingDirections = [][]int{
	{-1, -1}, // up-left
	{1, -1},  // up-right
}

// 1745
func checkNeighborsForRealXmas(grid [][]string, x, y, maxX, maxY int) bool {
	if grid[x][y] != "A" {
		return false
	}
	words := make([]string, 0)

	for _, direction := range startingDirections {
		dx, dy := direction[0], direction[1]

		i1 := x + dx
		j1 := y + dy
		// opposite direction
		i2 := x - dx
		j2 := y - dy
		if i1 > maxX || i1 < 0 || j1 > maxY || j1 < 0 || i2 > maxX || i2 < 0 || j2 > maxY || j2 < 0 {
			break
		}

		w := grid[i1][j1] + grid[x][y] + grid[i2][j2]
		words = append(words, w)
	}

	xmasCount := 0
	for _, w := range words {
		if w == "MAS" || w == "SAM" {
			xmasCount++
		}

	}
	return xmasCount == 2
}

func part2(grid [][]string) int {
	totalXmas := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			found := checkNeighborsForRealXmas(grid, i, j, len(grid)-1, len(grid[i])-1)
			if found {
				totalXmas++
			}
		}
	}
	return totalXmas
}

func main() {
	filename := "inputs/day_04/input.txt"
	input := _024.ReadInput(filename)

	grid := parseInput(input)

	println(part1(grid))
	println(part2(grid))
}

func parseInput(input string) [][]string {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, cell := range line {
			row = append(row, string(cell))
		}
		grid = append(grid, row)
	}
	return grid
}

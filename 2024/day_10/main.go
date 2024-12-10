package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

// 552
func part1(grid [][]int, startingPoints []_024.Point) int {
	sum := 0
	for _, point := range startingPoints {
		peaks := make(map[_024.Point]int)
		findPeaks(grid, point, peaks, 0)
		sum += len(peaks)
	}
	return sum
}

// 1225
func part2(grid [][]int, startingPoints []_024.Point) int {
	sum := 0
	for _, point := range startingPoints {
		peaks := make(map[_024.Point]int)
		findPeaks(grid, point, peaks, 0)
		for _, count := range peaks {
			sum += count
		}
	}
	return sum
}

func findPeaks(grid [][]int, point _024.Point, peaks map[_024.Point]int, number int) {
	if grid[point.Y][point.X] == number { // do nothing (implicitly finish the recursion chain) if the number in the current point is not the one we are looking for
		if grid[point.Y][point.X] == 9 { // if the number is 9, it means we reached to a peak
			peaks[point] += 1
			return
		}
		// here we are at the desired number but not at the peak yet.
		// next, we look for the next number (number + 1) around the current point, recursively.
		// this approach is very similar to Depth-First Search (DFS)
		for _, nextPoint := range _024.FindNonDiagonalValidPoints(point, grid) {
			findPeaks(grid, nextPoint, peaks, number+1)
		}
		return
	}
}

func main() {
	filename := "inputs/day_10/input.txt"
	input := _024.ReadInput(filename)

	grid, startingPoints := parseInput(input)

	println(part1(grid, startingPoints))
	println(part2(grid, startingPoints))
}

func parseInput(input string) ([][]int, []_024.Point) {
	grid := make([][]int, 0)

	for _, s := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, cell := range strings.Split(s, "") {
			row = append(row, _024.MustParseInt(cell))
		}
		grid = append(grid, row)
	}
	startingPoints := make([]_024.Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				startingPoints = append(startingPoints, _024.Point{X: j, Y: i})
			}
		}
	}
	return grid, startingPoints
}

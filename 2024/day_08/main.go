package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"math"
	"strings"
)

// 394
func part1(antennas map[string][]_024.Point, maxX, maxY int) int {
	return len(findAntinodes(antennas, 1, maxX, maxY))
}

// 1277
func part2(antennas map[string][]_024.Point, depth, maxX, maxY int) int {
	antinodes := findAntinodes(antennas, depth, maxX, maxY)
	// not completely sure about this, but it seems to work
	// add all antennas with more than 1 point as antinodes
	for _, points := range antennas {
		if len(points) > 1 { // antennas with only one point can't have antinodes
			for _, point := range points {
				antinodes[point] = true
			}
		}
	}
	return len(antinodes)
}

func findAntinodes(antennas map[string][]_024.Point, depth, maxX, maxY int) map[_024.Point]bool {
	antinodes := make(map[_024.Point]bool)
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				pi := points[i]
				pj := points[j]

				// distance = pi - pj
				distance := pi.DistanceTo(pj)
				n := 1
				for {
					if n > depth {
						break
					}
					antinode := pi.MoveN(distance, n) // pi's antinode: move one distance away from pi
					if antinode.IsValid(maxX, maxY) {
						antinodes[antinode] = true
						n++
						continue
					}
					break
				}
				n = 1
				for {
					if n > depth {
						break
					}
					antinode := pj.MoveN(distance, -n) // pj's antinode: move one distance away from pj
					if antinode.IsValid(maxX, maxY) {
						antinodes[antinode] = true
						n++
						continue
					}
					break
				}
			}
		}
	}
	return antinodes
}

func main() {
	filename := "inputs/day_08/input.txt"
	input := _024.ReadInput(filename)

	antennas, maxX, maxY := parseInput(input)

	println(part1(antennas, maxX, maxY))
	maxDepth := int(math.Max(float64(maxX), float64(maxY))) // float...  ̄\_(ツ)_/ ̄
	println(part2(antennas, maxDepth, maxX, maxY))
}

func parseInput(input string) (map[string][]_024.Point, int, int) {
	antennas := make(map[string][]_024.Point)
	split := strings.Split(input, "\n")
	maxY := len(split)
	maxX := len(split[0])
	for i, line := range split {
		for j, c := range line {
			frequency := string(c)
			if frequency != "." {
				antennas[frequency] = append(antennas[frequency], _024.Point{X: i, Y: j})
			}
		}
	}
	return antennas, maxX, maxY
}

package main

import (
	"fmt"
	_024 "github.com/endertunc/advent-of-code/2024"
	"strings"
)

type Robot struct {
	Current  _024.Point
	Velocity _024.Point
}

// 229839456
func part1(robots []Robot, MaxX, MaxY int) int {
	midPoint := _024.Point{X: MaxX / 2, Y: MaxY / 2}
	quadrants := make(map[int]int)

	for _, robot := range robots {
		r := moveNTimesWithMath(robot, MaxX, MaxY, 100)
		if r.Current.X < midPoint.X && r.Current.Y < midPoint.Y {
			quadrants[_024.NORTH_WEST] += 1
		}
		if r.Current.X > midPoint.X && r.Current.Y < midPoint.Y {
			quadrants[_024.NORTH_EAST] += 1
		}
		if r.Current.X > midPoint.X && r.Current.Y > midPoint.Y {
			quadrants[_024.SOUTH_EAST] += 1
		}
		if r.Current.X < midPoint.X && r.Current.Y > midPoint.Y {
			quadrants[_024.SOUTH_WEST] += 1
		}
	}

	sum := 1
	for _, nOfRobots := range quadrants {
		sum *= nOfRobots
	}
	return sum
}

func moveNTimesWithIteration(robot Robot, MaxX, MaxY, n int) Robot {
	for i := 0; i < n; i++ {
		robot.Current = robot.Current.Move(robot.Velocity)
		if !robot.Current.IsValid(MaxX, MaxY) {
			if robot.Current.Y < 0 {
				robot.Current.Y = MaxY + robot.Current.Y + 1
			}
			if robot.Current.Y > MaxY {
				robot.Current.Y = robot.Current.Y - MaxY - 1
			}

			if robot.Current.X < 0 {
				robot.Current.X = MaxX + robot.Current.X + 1
			}
			if robot.Current.X > MaxX {
				robot.Current.X = robot.Current.X - MaxX - 1
			}
		}
	}
	return robot
}

func moveNTimesWithMath(robot Robot, MaxX, MaxY, n int) Robot {
	p := robot.Current
	step := _024.Point{
		X: robot.Velocity.X * n,
		Y: robot.Velocity.Y * n,
	}
	p = p.Move(step)
	robot.Current = _024.Point{X: ((p.X % MaxX) + MaxX) % MaxX, Y: ((p.Y % MaxY) + MaxY) % MaxY}
	return robot
}

func main() {
	filename := "inputs/day_14/input.txt"
	input := _024.ReadInput(filename)

	MaxX, MaxY := 101, 103
	//MaxX, MaxY := 11, 7

	robots := parseInput(input)

	println(part1(robots, MaxX, MaxY))
}

func parseInput(input string) []Robot {
	robots := make([]Robot, 0)
	var pX, pY, vX, vY int
	for _, line := range strings.Split(input, "\n") {
		_024.Must(fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY))
		robots = append(robots, Robot{
			Current:  _024.Point{X: pX, Y: pY},
			Velocity: _024.Point{X: vX, Y: vY},
		})
	}

	return robots
}

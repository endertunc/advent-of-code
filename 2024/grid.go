package _024

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

type Point struct {
	X, Y int
}

func (p Point) IsValid(maxX, maxY int) bool {
	return p.X >= 0 && p.X < maxX && p.Y >= 0 && p.Y < maxY
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

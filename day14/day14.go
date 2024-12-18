package day14

import (
	"fmt"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

var maxX int = 101
var maxY int = 103

type robot struct {
	position, vector util.Point
}

func (r *robot) move() {
	x := r.position.X + r.vector.X
	y := r.position.Y + r.vector.Y
	if x < 0 {
		x = x + maxX
	}
	if y < 0 {
		y = y + maxY
	}
	r.position = util.Point{X: x % maxX, Y: y % maxY}
}

func init() {
	aoc.Register(2024, 14, solve2024Day14Part1, solve2024Day14Part2)
}

func solve2024Day14Part1(lines []string, test bool) interface{} {
	robots := make([]*robot, 0)
	for _, line := range lines {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, &robot{util.Point{X: px, Y: py}, util.Point{X: vx, Y: vy}})
	}

	quadrant1Count, quadrant2Count, quadrant3Count, quadrant4Count := 0, 0, 0, 0
	for _, r := range robots {
		for range 100 {
			r.move()
		}
		if r.position.X < maxX/2 && r.position.Y < maxY/2 {
			quadrant1Count++
		} else if r.position.X < maxX/2 && r.position.Y > maxY/2 {
			quadrant2Count++
		} else if r.position.X > maxX/2 && r.position.Y < maxY/2 {
			quadrant3Count++
		} else if r.position.X > maxX/2 && r.position.Y > maxY/2 {
			quadrant4Count++
		}
	}
	return quadrant1Count * quadrant2Count * quadrant3Count * quadrant4Count
}

func solve2024Day14Part2(lines []string, test bool) interface{} {
	robots := make([]*robot, 0)
	for _, line := range lines {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, &robot{util.Point{X: px, Y: py}, util.Point{X: vx, Y: vy}})
	}

	seconds := 0
	for {
		seconds++
		currentPositions := make(map[util.Point]int)
		for _, r := range robots {
			r.move()
			currentPositions[r.position] = currentPositions[r.position] + 1
		}
		if len(currentPositions) == len(robots) {
			break
		}
	}
	print(robots)
	return seconds
}

func print(robots []*robot) {
	grid := make(map[util.Point]int)
	for _, r := range robots {
		grid[r.position] = grid[r.position] + 1
	}

	for y := range maxY {
		for x := range maxX {
			if v, ok := grid[util.Point{X: x, Y: y}]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

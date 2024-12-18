package day6

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 6, solve2024Day6Part1, solve2024Day6Part2)
}

func solve2024Day6Part1(lines []string, test bool) interface{} {
	layout := util.LinesToPointMap(lines)
	position := util.Point{}
	for point, char := range layout {
		if char == "^" {
			position = point
			break
		}
	}

	directions := []util.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}
	currentDirection := 0
	visited := make(map[util.Point]map[int]bool)
	for {
		newPosition := position.Add(directions[currentDirection])
		if layout[newPosition] == "" {
			break
		}
		if layout[newPosition] == "#" {
			currentDirection = (currentDirection + 1) % 4
		} else {
			if _, ok := visited[position]; !ok {
				visited[position] = make(map[int]bool)
			}
			visited[position][currentDirection] = true
			position = newPosition
		}
	}
	if _, ok := visited[position]; !ok {
		visited[position] = make(map[int]bool)
	}
	visited[position][currentDirection] = true

	return len(visited)
}

func solve2024Day6Part2(lines []string, test bool) interface{} {
	layout := util.LinesToPointMap(lines)
	initialStart := util.Point{}
	for point, char := range layout {
		if char == "^" {
			initialStart = point
			break
		}
	}

	obstacles := make([]util.Point, 0)
	directions := []util.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}
	for point, char := range layout {
		if char == "." {
			position := initialStart
			currentDirection := 0
			visited := make(map[util.Point]map[int]bool)
			for {
				newPosition := position.Add(directions[currentDirection])
				if layout[newPosition] == "" {
					break
				}
				if c := layout[newPosition]; c == "#" || newPosition == point {
					currentDirection = (currentDirection + 1) % 4
				} else {
					if visited[position][currentDirection] {
						obstacles = append(obstacles, point)
						break
					}
					if _, ok := visited[position]; !ok {
						visited[position] = make(map[int]bool)
					}
					visited[position][currentDirection] = true
					position = newPosition
				}
			}
		}
	}

	return len(obstacles)
}

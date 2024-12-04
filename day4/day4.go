package day4

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 4, solve2024Day4Part1, solve2024Day4Part2)
}

func solve2024Day4Part1(lines []string) interface{} {
	input := util.LinesToPointMap(lines)
	count := 0
	for point, letter := range input {
		if letter == "X" {
			for _, neighbour := range []util.Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: -1}, {X: -1, Y: -1}, {X: -1, Y: 1}} {
				if input[point.Add(neighbour)] == "M" && input[point.Add(neighbour).Add(neighbour)] == "A" && input[point.Add(neighbour).Add(neighbour).Add(neighbour)] == "S" {
					count++
				}
			}
		}

	}
	return count
}

func solve2024Day4Part2(lines []string) interface{} {
	input := util.LinesToPointMap(lines)
	count := 0
	for point, letter := range input {
		if letter == "A" {
			if ((input[point.Add(util.Point{X: -1, Y: -1})] == "M" && input[point.Add(util.Point{X: 1, Y: 1})] == "S") ||
				(input[point.Add(util.Point{X: -1, Y: -1})] == "S" && input[point.Add(util.Point{X: 1, Y: 1})] == "M")) &&
				((input[point.Add(util.Point{X: -1, Y: 1})] == "M" && input[point.Add(util.Point{X: 1, Y: -1})] == "S") ||
					(input[point.Add(util.Point{X: -1, Y: 1})] == "S" && input[point.Add(util.Point{X: 1, Y: -1})] == "M")) {
				count++
			}
		}

	}
	return count
}

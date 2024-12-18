package day18

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 18, solve2024Day18Part1, solve2024Day18Part2)
}

func solve2024Day18Part1(lines []string, test bool) interface{} {
	maxX, maxY, numberOfBytes := 70, 70, 1024
	if test {
		maxX, maxY, numberOfBytes = 6, 6, 12
	}

	grid := util.CreatePointMapOfBasicNodes(maxX+1, maxY+1, ".", filterNeighbour)

	for i := range numberOfBytes {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].Value = "#"
	}

	start := grid[util.Point{X: 0, Y: 0}]
	end := grid[util.Point{X: maxX, Y: maxY}]

	_, distance, _ := graph.FindShortestPath(start, end)

	return distance
}

func solve2024Day18Part2(lines []string, test bool) interface{} {
	maxX, maxY, numberOfBytes := 70, 70, 1024
	if test {
		maxX, maxY, numberOfBytes = 6, 6, 12
	}

	grid := util.CreatePointMapOfBasicNodes(maxX+1, maxY+1, ".", filterNeighbour)

	for i := range len(lines) {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].Value = "#"
	}

	start := grid[util.Point{X: 0, Y: 0}]
	end := grid[util.Point{X: maxX, Y: maxY}]

	for i := len(lines) - 1; i >= numberOfBytes; i-- {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].Value = "."

		_, _, found := graph.FindShortestPath(start, end)
		if found {
			return lines[i]
		}
	}

	return "Path always open"
}

func filterNeighbour(neighbour *graph.BasicNode) bool {
	return neighbour.Value != "#"
}

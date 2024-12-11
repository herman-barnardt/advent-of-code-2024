package day10

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 10, solve2024Day10Part1, solve2024Day10Part2)
}

func solve2024Day10Part1(lines []string) interface{} {
	grid := util.LinesToPointMapOfInts(lines)
	sum := 0
	for p, v := range grid {
		if v == 0 {
			heads, _ := findPaths(grid, p, make(map[util.Point]bool), 0)
			sum += len(heads)
		}
	}
	return sum
}

func solve2024Day10Part2(lines []string) interface{} {
	grid := util.LinesToPointMapOfInts(lines)
	sum := 0
	for p, v := range grid {
		if v == 0 {
			_, numberOfPaths := findPaths(grid, p, make(map[util.Point]bool), 0)
			sum += numberOfPaths
		}
	}
	return sum
}

func getValidNeighbours(grid map[util.Point]int, point util.Point) []util.Point {
	neighbours := make([]util.Point, 0)
	for _, n := range point.GetAdjacent() {
		if v, ok := grid[n]; ok && v == grid[point]+1 {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func findPaths(grid map[util.Point]int, start util.Point, heads map[util.Point]bool, numberOfPaths int) (map[util.Point]bool, int) {
	if grid[start] == 9 {
		heads[start] = true
		return heads, numberOfPaths + 1
	}
	neighbours := getValidNeighbours(grid, start)
	if len(neighbours) == 0 {
		return heads, numberOfPaths
	}

	for _, n := range neighbours {
		heads, numberOfPaths = findPaths(grid, n, heads, numberOfPaths)
	}

	return heads, numberOfPaths
}

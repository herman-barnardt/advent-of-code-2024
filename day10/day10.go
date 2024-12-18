package day10

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 10, solve2024Day10Part1, solve2024Day10Part2)
}

func solve2024Day10Part1(lines []string, test bool) interface{} {
	grid := util.LinesToPointMapOfInts(lines)
	sum := 0
	for p, v := range grid {
		if v == 0 {
			toVist := []util.Point{p}
			trailHeads := make(map[util.Point]bool)
			for len(toVist) > 0 {
				current := toVist[0]
				toVist = toVist[1:]

				if grid[current] == 9 {
					trailHeads[current] = true
				}
				for _, n := range current.GetAdjacent() {
					if v, ok := grid[n]; ok && v == grid[current]+1 {
						toVist = append(toVist, n)
					}
				}
			}
			sum += len(trailHeads)
		}
	}
	return sum
}

func solve2024Day10Part2(lines []string, test bool) interface{} {
	grid := util.LinesToPointMapOfInts(lines)
	sum := 0
	for p, v := range grid {
		if v == 0 {
			toVist := []util.Point{p}
			trailHeads := make(map[util.Point]int)
			for len(toVist) > 0 {
				current := toVist[0]
				toVist = toVist[1:]

				if grid[current] == 9 {
					trailHeads[current] = trailHeads[current] + 1
				}
				for _, n := range current.GetAdjacent() {
					if v, ok := grid[n]; ok && v == grid[current]+1 {
						toVist = append(toVist, n)
					}
				}
			}
			for _, v := range trailHeads {
				sum += v
			}
		}
	}
	return sum
}

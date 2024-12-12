package day12

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 12, solve2024Day12Part1, solve2024Day12Part2)
}

func solve2024Day12Part1(lines []string) interface{} {
	plots := util.LinesToPointMap(lines)
	visited := map[util.Point]bool{}
	sum := 0
	for p := range plots {
		if visited[p] {
			continue
		}
		visited[p] = true
		regionFences := 0
		regionArea := 1
		toVisit := []util.Point{p}

		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]

			for _, neighbour := range current.GetAdjacent() {
				if plots[neighbour] != plots[current] {
					regionFences++
				} else if !visited[neighbour] {
					toVisit = append(toVisit, neighbour)
					visited[neighbour] = true
					regionArea++
				}
			}
		}
		sum += regionArea * regionFences
	}
	return sum
}

func solve2024Day12Part2(lines []string) interface{} {
	plots := util.LinesToPointMap(lines)
	visited := map[util.Point]bool{}
	sum := 0
	for p := range plots {
		if visited[p] {
			continue
		}
		visited[p] = true

		regionArea := 1
		regionSides := 0
		toVisit := []util.Point{p}
		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]

			for _, d := range []util.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}} {
				if neighbour := current.Add(d); plots[neighbour] != plots[current] {
					neighbourDiagonal := current.Add(util.Point{X: -d.Y, Y: d.X})
					if plots[neighbourDiagonal] != plots[current] || plots[neighbourDiagonal.Add(d)] == plots[current] {
						regionSides++
					}
				} else if !visited[neighbour] {
					visited[neighbour] = true
					toVisit = append(toVisit, neighbour)
					regionArea++
				}
			}
		}
		sum += regionArea * regionSides
	}
	return sum
}

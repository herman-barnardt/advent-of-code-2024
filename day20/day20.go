package day20

import (
	"slices"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 20, solve2024Day20Part1, solve2024Day20Part2)
}

func solve2024Day20Part1(lines []string, test bool) interface{} {
	maze := util.LinesToPointMapOfBasicNodes(lines, func(neighbour *graph.BasicNode) bool {
		return neighbour.Value != "#"
	})
	var start, end *graph.BasicNode

	for _, n := range maze {
		if n.Value == "S" {
			start = n
		}
		if n.Value == "E" {
			end = n
		}
		if start != nil && end != nil {
			break
		}
	}

	nodePath, floatDistance, _ := graph.FindShortestPath(start, end)

	distance := int(floatDistance)

	path := make([]*graph.BasicNode, 0)

	for _, n := range nodePath {
		path = append(path, n.(*graph.BasicNode))
	}

	distanceFromStart := make(map[util.Point]int)

	for i, n := range path {
		p := util.Point{X: n.X, Y: n.Y}
		distanceFromStart[p] = distance - i
	}

	count := 0

	for _, s := range path {
		startPoint := util.Point{X: s.X, Y: s.Y}
		for x := -2; x <= 2; x++ {
			maxY := 2 - util.IntAbs(x)
			for y := -maxY; y <= maxY; y++ {
				p := util.Point{X: s.X + x, Y: s.Y + y}
				if n, ok := maze[p]; ok && slices.Contains(path, n) && !s.Equal(n) {
					distanceBetween := util.DistanceBetween(&startPoint, &p)
					if distanceFromStart[p]-distanceBetween-distanceFromStart[startPoint] >= 100 {
						count++
					}
				}
			}
		}
	}

	return count
}

func solve2024Day20Part2(lines []string, test bool) interface{} {
	maze := util.LinesToPointMapOfBasicNodes(lines, func(neighbour *graph.BasicNode) bool {
		return neighbour.Value != "#"
	})
	var start, end *graph.BasicNode

	for _, n := range maze {
		if n.Value == "S" {
			start = n
		}
		if n.Value == "E" {
			end = n
		}
		if start != nil && end != nil {
			break
		}
	}

	nodePath, floatDistance, _ := graph.FindShortestPath(start, end)

	distance := int(floatDistance)

	path := make([]*graph.BasicNode, 0)

	for _, n := range nodePath {
		path = append(path, n.(*graph.BasicNode))
	}

	distanceFromStart := make(map[util.Point]int)

	for i, n := range path {
		p := util.Point{X: n.X, Y: n.Y}
		distanceFromStart[p] = distance - i
	}

	count := 0

	for _, s := range path {
		startPoint := util.Point{X: s.X, Y: s.Y}
		for x := -20; x <= 20; x++ {
			maxY := 20 - util.IntAbs(x)
			for y := -maxY; y <= maxY; y++ {
				p := util.Point{X: s.X + x, Y: s.Y + y}
				if n, ok := maze[p]; ok && slices.Contains(path, n) && !s.Equal(n) {
					distanceBetween := util.DistanceBetween(&startPoint, &p)
					if distanceFromStart[p]-distanceBetween-distanceFromStart[startPoint] >= 100 {
						count++
					}
				}
			}
		}
	}

	return count
}

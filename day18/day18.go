package day18

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

type pointNode struct {
	position   util.Point
	value      string
	neighbours []*pointNode
}

func (p *pointNode) GetNeighbours() []graph.Node {
	nodeNeighbours := make([]graph.Node, 0)
	for _, neighbour := range p.neighbours {
		if neighbour.value != "#" {
			nodeNeighbours = append(nodeNeighbours, neighbour)
		}
	}
	return nodeNeighbours
}
func (p *pointNode) GetCost(to graph.Node) float64 {
	return 1
}

func (p *pointNode) GetHeuristicCost(to graph.Node) float64 {
	return 1
}

func (p *pointNode) Equal(to graph.Node) bool {
	toPoint := to.(*pointNode)
	return p.position.X == toPoint.position.X && p.position.Y == toPoint.position.Y
}

func init() {
	aoc.Register(2024, 18, solve2024Day18Part1, solve2024Day18Part2)
}

func solve2024Day18Part1(lines []string) interface{} {
	maxX, maxY, numberOfBytes := 70, 70, 1024

	grid := make(map[util.Point]*pointNode)

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			grid[util.Point{X: x, Y: y}] = &pointNode{util.Point{X: x, Y: y}, ".", make([]*pointNode, 0)}
		}
	}

	for i := range numberOfBytes {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].value = "#"
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			node := grid[util.Point{X: x, Y: y}]
			for _, n := range node.position.GetAdjacent() {
				if v, ok := grid[n]; ok && v.value != "#" {
					node.neighbours = append(node.neighbours, v)
				}
			}
		}
	}

	start := grid[util.Point{X: 0, Y: 0}]
	end := grid[util.Point{X: maxX, Y: maxY}]

	_, distance, _ := graph.FindShortestPath(start, end)

	return distance
}

func solve2024Day18Part2(lines []string) interface{} {
	maxX, maxY, numberOfBytes := 70, 70, 1024

	grid := make(map[util.Point]*pointNode)

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			grid[util.Point{X: x, Y: y}] = &pointNode{util.Point{X: x, Y: y}, ".", make([]*pointNode, 0)}
		}
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			node := grid[util.Point{X: x, Y: y}]
			for _, n := range node.position.GetAdjacent() {
				if v, ok := grid[n]; ok && v.value != "#" {
					node.neighbours = append(node.neighbours, v)
				}
			}
		}
	}

	for i := range len(lines) {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].value = "#"
	}

	start := grid[util.Point{X: 0, Y: 0}]
	end := grid[util.Point{X: maxX, Y: maxY}]

	for i := len(lines) - 1; i >= numberOfBytes; i-- {
		coords := util.StringToIntSlice(lines[i], ",")
		grid[util.Point{X: coords[0], Y: coords[1]}].value = "."

		_, _, found := graph.FindShortestPath(start, end)
		if found {
			return lines[i]
		}
	}

	return "Path always open"
}

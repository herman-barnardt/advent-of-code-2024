package day16

import (
	"fmt"
	"slices"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/graph"
	"github.com/herman-barnardt/aoc/util"
)

type Maze struct {
	grid  map[util.Point]string
	start util.Point
	end   util.Point
}

type QueueItem struct {
	position  util.Point
	direction int
	score     int
	path      []util.Point
}

func (p QueueItem) key() string {
	return fmt.Sprintf("%d,%d,%d", p.position.X, p.position.Y, p.direction)
}

func (i QueueItem) Less(j graph.QItem) bool {
	return i.score < j.(QueueItem).score
}

func (m Maze) isValid(p util.Point) bool {
	v, ok := m.grid[p]
	return ok && v != "#"
}

func (m Maze) isEnd(p util.Point) bool {
	return p == m.end
}

func init() {
	aoc.Register(2024, 16, solve2024Day16Part1, solve2024Day16Part2)
}

func solve2024Day16Part1(lines []string, test bool) interface{} {
	grid := util.LinesToPointMap(lines)
	var start, end util.Point

	for p, v := range grid {
		if v == "S" {
			start = p
		}
		if v == "E" {
			end = p
		}
	}

	lowestScore := findLowestScore(start, end, grid)
	return lowestScore
}

func solve2024Day16Part2(lines []string, test bool) interface{} {
	grid := util.LinesToPointMap(lines)
	var start, end util.Point

	for p, v := range grid {
		if v == "S" {
			start = p
		}
		if v == "E" {
			end = p
		}
	}

	maze := Maze{grid, start, end}
	lowestScore := findLowestScore(start, end, grid)
	paths := findAllOptimalPaths(maze, lowestScore)

	unique := make(map[util.Point]bool)
	for _, path := range paths {
		for _, p := range path {
			unique[p] = true
		}
	}
	return len(unique)
}

func findLowestScore(start util.Point, end util.Point, grid map[util.Point]string) int {
	queue := graph.NewPriorityQueue()
	queue.Push(QueueItem{start, 1, 0, nil})
	visited := make(map[string]bool)

	for queue.Length() > 0 {

		current := queue.Pop().(QueueItem)

		if end == current.position {
			return current.score
		}

		key := current.key()
		if visited[key] {
			continue
		}
		visited[key] = true

		nextPos := current.position.Add(util.Directions[current.direction])

		if v, ok := grid[nextPos]; ok && v != "#" {
			queue.Push(QueueItem{
				nextPos,
				current.direction,
				current.score + 1,
				nil,
			})
		}

		queue.Push(QueueItem{current.position, (current.direction + 1) % 4, current.score + 1000, nil})
		queue.Push(QueueItem{current.position, (current.direction + 3) % 4, current.score + 1000, nil})
	}

	return -1
}

func findAllOptimalPaths(m Maze, targetScore int) [][]util.Point {
	queue := graph.NewPriorityQueue()
	queue.Push(QueueItem{m.start, 1, 0, []util.Point{m.start}})
	visited := make(map[string]int)
	var paths [][]util.Point

	for queue.Length() > 0 {
		current := queue.Pop().(QueueItem)

		if current.score > targetScore {
			continue
		}

		key := current.key()
		if score, exists := visited[key]; exists && score < current.score {
			continue
		}
		visited[key] = current.score

		if m.isEnd(current.position) && current.score == targetScore {
			paths = append(paths, current.path)
			continue
		}

		nextPos := current.position.Add(util.Directions[current.direction])
		if m.isValid(nextPos) {
			newPath := slices.Clone(current.path)
			queue.Push(QueueItem{
				nextPos,
				current.direction,
				current.score + 1,
				append(newPath, nextPos),
			})
		}

		for _, newDir := range []int{(current.direction + 1) % 4, (current.direction + 3) % 4} {
			queue.Push(QueueItem{
				current.position,
				newDir,
				current.score + 1000,
				current.path,
			})
		}
	}

	return paths
}

package day15

import (
	"maps"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

var directions map[string]util.Point = map[string]util.Point{
	"^": {X: 0, Y: -1},
	"v": {X: 0, Y: 1},
	">": {X: 1, Y: 0},
	"<": {X: -1, Y: 0},
}

func init() {
	aoc.Register(2024, 15, solve2024Day15Part1, solve2024Day15Part2)
}

func solve2024Day15Part1(lines []string) interface{} {
	parts := util.SplitLines(lines, "")
	grid, moves := util.LinesToPointMap(parts[0]), strings.Split(strings.Join(parts[1], ""), "")
	robot := util.Point{}
	for p, v := range grid {
		if v == "@" {
			robot = p
			break
		}
	}
	for _, m := range moves {
		_, robot, grid = canMove(robot, directions[m], grid)
	}

	sum := 0
	for p, v := range grid {
		if v == "O" {
			sum += 100*p.Y + p.X
		}
	}
	return sum
}

func solve2024Day15Part2(lines []string) interface{} {
	parts := util.SplitLines(lines, "")
	for i := range len(parts[0]) {
		parts[0][i] = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(parts[0][i], "#", "##"), "O", "[]"), ".", ".."), "@", "@.")
	}
	grid, moves := util.LinesToPointMap(parts[0]), strings.Split(strings.Join(parts[1], ""), "")
	robot := util.Point{}
	for p, v := range grid {
		if v == "@" {
			robot = p
			break
		}
	}
	for _, m := range moves {
		_, robot, grid = canMove(robot, directions[m], grid)
	}
	sum := 0
	for p, v := range grid {
		if v == "[" {
			sum += 100*p.Y + p.X
		}
	}
	return sum
}

func canMove(position util.Point, direction util.Point, grid map[util.Point]string) (bool, util.Point, map[util.Point]string) {
	saveState := maps.Clone(grid)
	newPosition := position.Add(direction)
	if grid[newPosition] == "." {
		grid[newPosition] = grid[position]
		grid[position] = "."
		return true, newPosition, grid
	}
	if grid[newPosition] == "O" {
		moved := false
		moved, _, grid = canMove(newPosition, direction, grid)
		if moved {
			grid[newPosition] = grid[position]
			grid[position] = "."
			return true, newPosition, grid
		}
	}
	if grid[newPosition] == "[" || grid[newPosition] == "]" {
		if direction == directions["<"] || direction == directions[">"] {
			moved := false
			moved, _, grid = canMove(newPosition, direction, grid)
			if moved {
				grid[newPosition] = grid[position]
				grid[position] = "."
				return true, newPosition, grid
			}
		} else {
			char := "["
			partnerChar := "]"
			partner := newPosition.Add(directions[">"])
			if grid[newPosition] == "]" {
				char = "]"
				partnerChar = "["
				partner = newPosition.Add(directions["<"])
			}
			boxCanMove, _, grid := canMove(newPosition, direction, grid)
			partnerCanMove, _, grid := canMove(partner, direction, grid)
			if boxCanMove && partnerCanMove {
				grid[newPosition] = grid[position]
				grid[position] = "."
				grid[newPosition.Add(direction)] = char
				grid[partner.Add(direction)] = partnerChar
				grid[partner] = "."
				return true, newPosition, grid
			}
		}
	}
	return false, position, saveState
}

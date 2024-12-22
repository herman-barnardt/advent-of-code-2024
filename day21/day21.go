package day21

import (
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

var numPad = map[string]util.Point{
	"A": {X: 2, Y: 0},
	"0": {X: 1, Y: 0},
	"1": {X: 0, Y: 1},
	"2": {X: 1, Y: 1},
	"3": {X: 2, Y: 1},
	"4": {X: 0, Y: 2},
	"5": {X: 1, Y: 2},
	"6": {X: 2, Y: 2},
	"7": {X: 0, Y: 3},
	"8": {X: 1, Y: 3},
	"9": {X: 2, Y: 3},
}

var arrowPad = map[string]util.Point{
	"A": {X: 2, Y: 1},
	"^": {X: 1, Y: 1},
	"<": {X: 0, Y: 0},
	"v": {X: 1, Y: 0},
	">": {X: 2, Y: 0},
}

func init() {
	aoc.Register(2024, 21, solve2024Day21Part1, solve2024Day21Part2)
}

func solve2024Day21Part1(lines []string, test bool) interface{} {
	sum := 0
	cache := make(map[string][]int)

	for _, line := range lines {
		moves := generateNumPadKeystrokes(line, "A")
		length := getKeystrokesLength(moves, 2, 1, cache)
		num, _ := strconv.Atoi(line[:len(line)-1])
		sum += num * length
	}

	return sum
}

func solve2024Day21Part2(lines []string, test bool) interface{} {
	sum := 0
	cache := make(map[string][]int)

	for _, line := range lines {
		moves := generateNumPadKeystrokes(line, "A")
		length := getKeystrokesLength(moves, 25, 1, cache)
		num, _ := strconv.Atoi(line[:len(line)-1])
		sum += num * length
	}

	return sum
}

func generateNumPadKeystrokes(input string, start string) string {
	curr := numPad[start]
	seq := ""

	for _, char := range strings.Split(input, "") {
		dest := numPad[char]
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := "", ""

		// Build horizontal moves
		if dx >= 0 {
			horiz = strings.Repeat(">", dx)
		} else {
			horiz = strings.Repeat("<", -dx)
		}

		// Build vertical moves
		if dy >= 0 {
			vert = strings.Repeat("^", dy)
		} else {
			vert = strings.Repeat("v", -dy)
		}

		// Order moves based on position
		if curr.Y == 0 && dest.X == 0 {
			seq += vert + horiz
		} else if curr.X == 0 && dest.Y == 0 {
			seq += horiz + vert
		} else if dx < 0 {
			seq += horiz + vert
		} else {
			seq += vert + horiz
		}

		curr = dest
		seq += "A"
	}
	return seq
}

func generateArrowPadKeyStrokes(input string, start string) string {
	curr := arrowPad[start]
	seq := ""

	for _, char := range strings.Split(input, "") {
		dest := arrowPad[char]
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := "", ""

		if dx >= 0 {
			horiz = strings.Repeat(">", dx)
		} else {
			horiz = strings.Repeat("<", -dx)
		}

		// Build vertical moves
		if dy >= 0 {
			vert = strings.Repeat("^", dy)
		} else {
			vert = strings.Repeat("v", -dy)
		}

		if curr.X == 0 && dest.Y == 1 {
			seq += horiz + vert
		} else if curr.Y == 1 && dest.X == 0 {
			seq += vert + horiz
		} else if dx < 0 {
			seq += horiz + vert
		} else {
			seq += vert + horiz
		}

		curr = dest
		seq += "A"
	}
	return seq
}

func getKeystrokesLength(input string, maxRobots, robot int, cache map[string][]int) int {
	if val, ok := cache[input]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	if _, ok := cache[input]; !ok {
		cache[input] = make([]int, maxRobots)
	}

	seq := generateArrowPadKeyStrokes(input, "A")
	if robot == maxRobots {
		return len(seq)
	}

	steps := getMovesFromKeystrokes(seq)
	count := 0
	for _, step := range steps {
		c := getKeystrokesLength(step, maxRobots, robot+1, cache)
		count += c
	}

	cache[input][robot-1] = count
	return count
}

func getMovesFromKeystrokes(input string) []string {
	var result []string
	var current string

	for _, char := range strings.Split(input, "") {
		current += char
		if char == "A" {
			result = append(result, current)
			current = ""
		}
	}
	return result
}

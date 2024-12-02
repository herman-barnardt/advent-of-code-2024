package day2

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 2, solve2024Day2Part1, solve2024Day2Part2)
}

func solve2024Day2Part1(lines []string) interface{} {
	count := 0
	for _, line := range lines {
		levels := make([]int, 0)
		parts := strings.Split(line, " ")
		for _, i := range parts {
			num, _ := strconv.Atoi(i)
			levels = append(levels, num)
		}
		if checkLevels(levels) {
			count++
		}
	}
	return count
}

func solve2024Day2Part2(lines []string) interface{} {
	count := 0
	for _, line := range lines {
		levels := make([]int, 0)
		parts := strings.Split(line, " ")
		for _, i := range parts {
			num, _ := strconv.Atoi(i)
			levels = append(levels, num)
		}
		for i := range levels {
			newRow := slices.Clone(levels)
			newRow = append(newRow[:i], newRow[i+1:]...)
			if checkLevels(newRow) {
				count++
				break
			}
		}
	}
	return count
}

func checkLevels(row []int) bool {
	order := row[0] > row[1]
	for i := 0; i < len(row)-1; i++ {
		diff := util.IntAbs(row[i] - row[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
		currentOrder := row[i] > row[i+1]
		if currentOrder != order {
			return false
		}
	}
	return true
}

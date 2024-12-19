package day19

import (
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 19, solve2024Day19Part1, solve2024Day19Part2)
}

func solve2024Day19Part1(lines []string, test bool) interface{} {
	parts := util.SplitLines(lines, "")
	towels := strings.Split(parts[0][0], ", ")
	patterns := parts[1]

	count := 0

	for _, pattern := range patterns {
		if possibleCount(towels, pattern, make(map[string]int)) > 0 {
			count++
		}
	}

	return count
}

func solve2024Day19Part2(lines []string, test bool) interface{} {
	parts := util.SplitLines(lines, "")
	towels := strings.Split(parts[0][0], ", ")
	patterns := parts[1]

	sum := 0

	for _, pattern := range patterns {
		sum += possibleCount(towels, pattern, make(map[string]int))
	}

	return sum
}

func possibleCount(towels []string, pattern string, cache map[string]int) int {
	if _, ok := cache[pattern]; !ok {
		if len(pattern) == 0 {
			cache[pattern] = 1
			return 1
		}

		sum := 0
		for _, towel := range towels {
			if strings.HasPrefix(pattern, towel) {
				sum += possibleCount(towels, pattern[len(towel):], cache)
			}
		}
		cache[pattern] = sum
	}
	return cache[pattern]
}

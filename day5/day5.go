package day5

import (
	"slices"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 5, solve2024Day5Part1, solve2024Day5Part2)
}

func solve2024Day5Part1(lines []string) interface{} {
	rules := make(map[int][]int)
	parts := util.SplitLines(lines, "")
	sum := 0
	for _, line := range parts[0] {
		pages := util.StringToIntSlice(line, "|")
		if _, ok := rules[pages[1]]; !ok {
			rules[pages[1]] = make([]int, 0)
		}
		rules[pages[1]] = append(rules[pages[1]], pages[0])
	}
	for _, line := range parts[1] {
		pages := util.StringToIntSlice(line, ",")
		if isInOrder(pages, rules) {
			sum += pages[len(pages)/2]
		}
	}
	return sum
}

func solve2024Day5Part2(lines []string) interface{} {
	rules := make(map[int][]int)
	parts := util.SplitLines(lines, "")
	sum := 0
	for _, line := range parts[0] {
		pages := util.StringToIntSlice(line, "|")
		if _, ok := rules[pages[1]]; !ok {
			rules[pages[1]] = make([]int, 0)
		}
		rules[pages[1]] = append(rules[pages[1]], pages[0])
	}
	for _, line := range parts[1] {
		pages := util.StringToIntSlice(line, ",")
		if !isInOrder(pages, rules) {
			list := slices.Clone(pages)
			for !isInOrder(list, rules) {
				printed := make([]int, 0)
				for _, page := range list {
					for _, required := range rules[page] {
						if slices.Contains(pages, required) && !slices.Contains(printed, required) {
							printed = append(printed, required)
						}
					}
					if !slices.Contains(printed, page) {
						printed = append(printed, page)
					}
				}
				list = slices.Clone(printed)
			}
			sum += list[len(list)/2]
		}
	}
	return sum
}

func isInOrder(pages []int, rules map[int][]int) bool {
	printed := make([]int, 0)
	correct := true
	for _, page := range pages {
		for _, required := range rules[page] {
			correct = (!slices.Contains(pages, required) || slices.Contains(printed, required)) && correct
		}
		printed = append(printed, page)
	}
	return correct
}

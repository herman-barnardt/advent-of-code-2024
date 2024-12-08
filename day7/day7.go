package day7

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 7, solve2024Day7Part1, solve2024Day7Part2)
}

func solve2024Day7Part1(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := util.StringToIntSlice(strings.Replace(line, ":", "", -1), " ")
		ans := parts[0]
		values := parts[1:]
		totals := []int{values[0]}
		for _, i := range values[util.IntMin(1, len(values)-1):] {
			newTotals := make([]int, 0)
			for _, t := range totals {
				newTotals = append(newTotals, t+i)
				newTotals = append(newTotals, t*i)
			}
			totals = newTotals
		}
		if slices.Contains(totals, ans) {
			sum += ans
		}
	}
	return sum
}

func solve2024Day7Part2(lines []string) interface{} {
	sum := 0
	for _, line := range lines {
		parts := util.StringToIntSlice(strings.Replace(line, ":", "", -1), " ")
		ans := parts[0]
		values := parts[1:]
		totals := []int{values[0]}
		for _, i := range values[util.IntMin(1, len(values)-1):] {
			newTotals := make([]int, 0)
			for _, t := range totals {
				newTotals = append(newTotals, t+i)
				newTotals = append(newTotals, t*i)
				concatInt, _ := strconv.Atoi(strconv.Itoa(t) + strconv.Itoa(i))
				newTotals = append(newTotals, concatInt)
			}
			totals = newTotals
		}
		if slices.Contains(totals, ans) {
			sum += ans
		}
	}
	return sum
}

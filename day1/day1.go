package day1

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 1, solve2024Day1Part1, solve2024Day1Part2)
}

func solve2024Day1Part1(lines []string) interface{} {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0

	for i := range list1 {
		sum += util.IntAbs(list1[i] - list2[i])
	}

	return sum
}

func solve2024Day1Part2(lines []string) interface{} {
	list := make([]int, 0)
	countMap := make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list = append(list, num1)
		countMap[num2] += 1
	}

	sum := 0

	for _, i := range list {
		sum += i * countMap[i]
	}

	return sum
}

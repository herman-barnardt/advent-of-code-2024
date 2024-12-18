package day3

import (
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2024, 3, solve2024Day3Part1, solve2024Day3Part2)
}

func solve2024Day3Part1(lines []string, test bool) interface{} {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	sum := 0
	for _, line := range lines {
		muls := regex.FindAllString(line, -1)
		for _, mul := range muls {
			nums := strings.Split(mul[4:len(mul)-1], ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			sum += num1 * num2
		}
	}
	return sum
}
func solve2024Day3Part2(lines []string, test bool) interface{} {
	regex := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	sum := 0
	enabled := true
	for _, line := range lines {
		for regex.MatchString(line) {
			loc := regex.FindStringIndex(line)
			cmd := line[loc[0]:loc[1]]
			line = line[loc[1]:]
			if cmd == "do()" {
				enabled = true
			} else if cmd == "don't()" {
				enabled = false
			} else if enabled {
				nums := strings.Split(cmd[4:len(cmd)-1], ",")
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				sum += num1 * num2
			}
		}
	}
	return sum
}

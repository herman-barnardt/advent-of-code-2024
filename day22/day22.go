package day22

import (
	"math"
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 22, solve2024Day22Part1, solve2024Day22Part2)
}

func solve2024Day22Part1(lines []string, test bool) interface{} {
	sum := 0
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		for range 2000 {
			num = generateSecretNumber(num)
		}
		sum += num
	}
	return sum
}

func solve2024Day22Part2(lines []string, test bool) interface{} {
	sequences := make(map[[4]int]int)
	for _, line := range lines {
		sequence := make([]int, 0)
		num, _ := strconv.Atoi(line)
		price := num % 10
		seen := make(map[[4]int]struct{})
		for range 2000 {
			num = generateSecretNumber(num)
			newPrice := num % 10
			sequence = append(sequence, newPrice-price)
			price = newPrice
			if len(sequence) > 3 {
				key := [4]int(sequence[len(sequence)-4:])
				if _, ok := seen[key]; !ok {
					sequences[key] = sequences[key] + price
				}
				seen[key] = struct{}{}
			}
		}
	}
	max := math.MinInt
	for _, price := range sequences {
		max = util.IntMax(max, price)
	}
	return max
}

func generateSecretNumber(num int) int {
	num = ((num * 64) ^ num) % 16777216
	num = ((num / 32) ^ num) % 16777216
	num = ((num * 2048) ^ num) % 16777216
	return num
}

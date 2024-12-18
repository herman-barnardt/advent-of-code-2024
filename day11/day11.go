package day11

import (
	"slices"
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 11, solve2024Day11Part1, solve2024Day11Part2)
}

func solve2024Day11Part1(lines []string, test bool) interface{} {
	stones := util.StringToIntSlice(lines[0], " ")
	for c := 0; c < 25; c++ {
		newStones := make([]int, 0)
		for i, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				stoneString := strconv.Itoa(stone)
				newStone, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
				newStone2, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
				newStones = append(newStones, newStone)
				newStones = append(newStones, newStone2)
			} else {
				newStones = append(newStones, stones[i]*2024)
			}
		}
		stones = slices.Clone(newStones)
	}
	return len(stones)
}

func solve2024Day11Part2(lines []string, test bool) interface{} {
	stones := util.StringToIntSlice(lines[0], " ")
	sum := 0
	cache := make(map[int]map[int]int)
	for i := 1; i <= 75; i++ {
		cache[i] = make(map[int]int)
	}
	for _, stone := range stones {
		sum += countStones(stone, cache, 75)
	}
	return sum
}

func countStones(stone int, cache map[int]map[int]int, repitions int) int {
	if repitions == 0 {
		return 1
	} else {
		if v, ok := cache[repitions][stone]; ok {
			return v
		} else {
			if stone == 0 {
				count := countStones(1, cache, repitions-1)
				cache[repitions][stone] = count
				return count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				stoneString := strconv.Itoa(stone)
				newStone, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
				newStone2, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
				count := countStones(newStone, cache, repitions-1) + countStones(newStone2, cache, repitions-1)
				cache[repitions][stone] = count
				return count
			} else {
				count := countStones(stone*2024, cache, repitions-1)
				cache[repitions][stone] = count
				return count
			}
		}
	}
}

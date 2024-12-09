package day9

import (
	"slices"
	"strconv"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

type file struct {
	id     int
	length int
}

func init() {
	aoc.Register(2024, 9, solve2024Day9Part1, solve2024Day9Part2)
}

func solve2024Day9Part1(lines []string) interface{} {
	id := 0
	diskMap := make(map[int]string)
	count := 0
	for i, char := range strings.Split(lines[0], "") {
		val, _ := strconv.Atoi(char)
		if i%2 == 0 {
			for j := 0; j < val; j++ {
				diskMap[count] = strconv.Itoa(id)
				count++
			}
			id++
		} else {
			for j := 0; j < val; j++ {
				diskMap[count] = "."
				count++
			}
		}
	}
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] != "." {
			emptyIndex := -1
			for k := 0; k < count; k++ {
				if diskMap[k] == "." {
					emptyIndex = k
					break
				}
			}
			if emptyIndex > i || emptyIndex == -1 {
				break
			}
			diskMap[emptyIndex] = diskMap[i]
			diskMap[i] = "."
		}
	}
	sum := 0
	for i, v := range diskMap {
		num, _ := strconv.Atoi(v)
		sum += i * num
	}
	return sum
}

func solve2024Day9Part2(lines []string) interface{} {
	id := 0
	files := make([]file, 0)
	for i, char := range strings.Split(lines[0], "") {
		val, _ := strconv.Atoi(char)
		if i%2 == 0 {
			files = append(files, file{id, val})
			id++
		} else {
			files = append(files, file{-1, val})

		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		if files[i].id != -1 {
			for j := 0; j < i; j++ {
				if files[j].id == -1 && files[j].length >= files[i].length {
					filesToInsert := []file{files[i]}
					if files[i].length < files[j].length {
						filesToInsert = append(filesToInsert, file{-1, files[j].length - files[i].length})
					}
					files = append(slices.Clone(files[:j]), append(append(filesToInsert, append(files[j+1:i], file{-1, files[i].length})...), files[i+1:]...)...)
					break
				}
			}
		}
	}

	sum := 0
	count := 0
	for _, v := range files {
		if v.id != -1 {
			for j := 0; j < v.length; j++ {
				sum += count * v.id
				count++
			}
		} else {
			count += v.length
		}
	}
	return sum
}
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

func solve2024Day9Part1(lines []string, test bool) interface{} {
	id := 0
	files := make([]file, 0)
	for i, char := range strings.Split(lines[0], "") {
		val, _ := strconv.Atoi(char)
		if i%2 == 0 {
			for j := 0; j < val; j++ {
				files = append(files, file{id, 1})
			}
			id++
		} else {
			for j := 0; j < val; j++ {
				files = append(files, file{-1, 1})
			}
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		if files[i].id != -1 {
			for j := 0; j < i; j++ {
				if files[j].id == -1 {
					filesToInsert := []file{files[i]}
					files[i].id = -1
					files = slices.Delete(files, j, j+1)
					files = slices.Insert(files, j, filesToInsert...)
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

func solve2024Day9Part2(lines []string, test bool) interface{} {
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
					files[i].id = -1
					files = slices.Delete(files, j, j+1)
					files = slices.Insert(files, j, filesToInsert...)
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

package day8

import (
	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 8, solve2024Day8Part1, solve2024Day8Part2)
}

func solve2024Day8Part1(lines []string, test bool) interface{} {
	antennaMap := make(map[string][]util.Point)
	cityMap := util.LinesToPointMap(lines)
	for point, char := range cityMap {
		if char != "." {
			antennaMap[char] = append(antennaMap[char], point)
		}
	}

	antinodes := make(map[util.Point]bool)

	for _, antennas := range antennaMap {
		for _, point := range antennas {
			for _, point2 := range antennas {
				if point != point2 {
					diff := util.Point{X: (point.X - point2.X) * 2, Y: (point.Y - point2.Y) * 2}
					antinode := point2.Add(diff)
					if _, ok := cityMap[antinode]; ok {
						antinodes[antinode] = true
					}
				}
			}
		}
	}
	return len(antinodes)
}

func solve2024Day8Part2(lines []string, test bool) interface{} {
	antennaMap := make(map[string][]util.Point)
	cityMap := util.LinesToPointMap(lines)
	for point, char := range cityMap {
		if char != "." {
			antennaMap[char] = append(antennaMap[char], point)
		}
	}

	antinodes := make(map[util.Point]bool)

	for _, antennas := range antennaMap {
		for _, point := range antennas {
			for _, point2 := range antennas {
				if point != point2 {
					diff := util.Point{X: (point.X - point2.X), Y: (point.Y - point2.Y)}
					antinode := point2.Add(diff)
					for {
						if _, ok := cityMap[antinode]; ok {
							antinodes[antinode] = true
							antinode = antinode.Add(diff)
						} else {
							break
						}
					}
				}
			}
		}
	}
	return len(antinodes)
}

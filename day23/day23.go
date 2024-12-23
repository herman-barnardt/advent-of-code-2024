package day23

import (
	"math"
	"slices"
	"strings"

	aoc "github.com/herman-barnardt/aoc"
)

func init() {
	aoc.Register(2024, 23, solve2024Day23Part1, solve2024Day23Part2)
}

func solve2024Day23Part1(lines []string, test bool) interface{} {
	computers := make(map[string]struct{})
	connectedMap := make(map[string][]string)
	for _, line := range lines {
		parts := [2]string(strings.Split(line, "-"))
		computers[parts[0]] = struct{}{}
		computers[parts[1]] = struct{}{}
		connectedMap[parts[0]] = append(connectedMap[parts[0]], parts[1])
		connectedMap[parts[1]] = append(connectedMap[parts[1]], parts[0])
	}

	threeWayConnected := map[[3]string]struct{}{}

	for computer := range computers {
		for _, neighbour := range connectedMap[computer] {
			intersection := intersection(connectedMap[computer], connectedMap[neighbour])
			for _, i := range intersection {
				if strings.HasPrefix(computer, "t") || strings.HasPrefix(neighbour, "t") || strings.HasPrefix(i, "t") {
					if slices.Contains(connectedMap[i], computer) && slices.Contains(connectedMap[i], neighbour) {
						temp := []string{computer, neighbour, i}
						slices.Sort(temp)
						threeWayConnected[[3]string(temp)] = struct{}{}
					}
				}
			}
		}
	}

	return len(threeWayConnected)
}

func solve2024Day23Part2(lines []string, test bool) interface{} {
	computers := make(map[string]struct{})
	connectedMap := make(map[string][]string)
	for _, line := range lines {
		parts := [2]string(strings.Split(line, "-"))
		computers[parts[0]] = struct{}{}
		computers[parts[1]] = struct{}{}
		connectedMap[parts[0]] = append(connectedMap[parts[0]], parts[1])
		connectedMap[parts[1]] = append(connectedMap[parts[1]], parts[0])
	}

	networks := map[string]int{}

	for computer := range computers {
		for _, neighbour := range connectedMap[computer] {
			temp := []string{computer, neighbour}
			intersection := intersection(connectedMap[computer], connectedMap[neighbour])
			for _, i := range intersection {
				connected := true
				for _, t := range temp {
					connected = connected && slices.Contains(connectedMap[i], t)
				}
				if connected {
					temp = append(temp, i)
				}
			}
			if len(temp) > 2 {
				slices.Sort(temp)
				networks[strings.Join(temp, ",")] = len(temp)
			}
		}
	}

	max := math.MinInt
	password := ""
	for k, v := range networks {
		if v > max {
			max = v
			password = k
		}
	}

	return password
}

func intersection(list1, list2 []string) []string {
	retVal := make([]string, 0)
	hashMap := make(map[string]struct{})
	for _, v := range list1 {
		hashMap[v] = struct{}{}
	}

	for _, v := range list2 {
		if _, ok := hashMap[v]; ok {
			retVal = append(retVal, v)
		}
	}
	return retVal
}

package day13

import (
	"fmt"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

func init() {
	aoc.Register(2024, 13, solve2024Day13Part1, solve2024Day13Part2)
}

func solve2024Day13Part1(lines []string, test bool) interface{} {
	sum := 0
	for _, machine := range util.SplitLines(lines, "") {
		xa, ya, xb, yb, xp, yp := 0, 0, 0, 0, 0, 0
		fmt.Sscanf(machine[0], "Button A: X+%d, Y+%d", &xa, &ya)
		fmt.Sscanf(machine[1], "Button B: X+%d, Y+%d", &xb, &yb)
		fmt.Sscanf(machine[2], "Prize: X=%d, Y=%d", &xp, &yp)
		b := ((yp * xa) - (ya * xp)) / (-(xb * ya) + (yb * xa))
		a := (xp - xb*b) / xa
		if a < 100 && b < 100 && xa*a+xb*b == xp && ya*a+yb*b == yp {
			sum += a*3 + b
		}
	}
	return sum
}

func solve2024Day13Part2(lines []string, test bool) interface{} {
	sum := 0
	for _, machine := range util.SplitLines(lines, "") {
		xa, ya, xb, yb, xp, yp := 0, 0, 0, 0, 0, 0
		fmt.Sscanf(machine[0], "Button A: X+%d, Y+%d", &xa, &ya)
		fmt.Sscanf(machine[1], "Button B: X+%d, Y+%d", &xb, &yb)
		fmt.Sscanf(machine[2], "Prize: X=%d, Y=%d", &xp, &yp)
		xp, yp = xp+10000000000000, yp+10000000000000
		b := ((yp * xa) - (ya * xp)) / (-(xb * ya) + (yb * xa))
		a := (xp - xb*b) / xa
		if xa*a+xb*b == xp && ya*a+yb*b == yp {
			sum += a*3 + b
		}
	}
	return sum
}

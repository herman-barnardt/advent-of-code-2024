package day17

import (
	"fmt"
	"slices"
	"strconv"

	aoc "github.com/herman-barnardt/aoc"
	"github.com/herman-barnardt/aoc/util"
)

type state struct {
	a, b, c, instruction int
	out                  string
}

func init() {
	aoc.Register(2024, 17, solve2024Day17Part1, solve2024Day17Part2)
}

func solve2024Day17Part1(lines []string, test bool) interface{} {

	combo := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "A",
		5: "B",
		6: "C",
	}

	literal := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
	}

	a := 0

	fmt.Sscanf(lines[0], "Register A: %d", &a)

	instructions := util.StringToIntSlice(lines[4][9:], ",")

	aString := "a"
	bString := "b"
	cString := "c"
	outString := ""

	for i := 0; i < len(instructions); i += 2 {
		combo[4] = aString
		combo[5] = bString
		combo[6] = cString
		switch instructions[i] {
		case 0:
			aString = "(" + aString + " >> " + combo[instructions[i+1]] + ")"
		case 1:
			bString = "(" + bString + " ^ " + literal[instructions[i+1]] + ")"
			continue
		case 2:
			bString = "(" + combo[instructions[i+1]] + " % 8)"
			continue
		case 3:
			continue
		case 4:
			bString = "(" + bString + " ^ " + cString + ")"
			continue
		case 5:
			outString = "out += strconv.Itoa(" + combo[instructions[i+1]] + "%8)"
			continue
		case 6:
			bString = "(" + aString + " >> " + combo[instructions[i+1]] + ")"
			continue
		case 7:
			cString = "(" + aString + " >> " + combo[instructions[i+1]] + ")"
			continue
		}
	}

	//Use these lines in the method below to calculate the answer
	fmt.Println(outString)
	fmt.Println("a =", aString)

	out := ""

	for a > 0 {
		if len(out) > 0 {
			out += ","
		}
		//These lines should be replaced by the output of the first run
		out += strconv.Itoa(((((a % 8) ^ 2) ^ (a >> ((a % 8) ^ 2))) ^ 7) % 8)
		a = (a >> 3)
	}
	return out
}

func solve2024Day17Part2(lines []string, test bool) interface{} {
	instructions := util.StringToIntSlice(lines[4][9:], ",")
	slices.Reverse(instructions)

	quines := []int{0}
	for _, n := range instructions {
		newQuines := []int{}
		for _, current := range quines {
			for i := range 9 {
				new := (current << 3) + i
				if ((((new%8)^2)^(new>>((new%8)^2)))^7)%8 == n {
					newQuines = append(newQuines, new)
				}
			}
		}
		quines = newQuines
	}

	slices.Sort(quines)

	return quines[0]
}

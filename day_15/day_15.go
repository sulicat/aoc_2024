package main

import (
	"fmt"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

type Pos [2]int

func (a Pos) Add(b Pos) Pos {
	return Pos{
		a[0] + b[0],
		a[1] + b[1]}
}

func (a Pos) Sub(b Pos) Pos {
	return Pos{
		a[0] - b[0],
		a[1] - b[1]}
}

var width int
var height int
var obstacles map[Pos]int
var boxes map[Pos]int
var robot Pos
var commands []Pos

func printmap() {
	for y := range height {
		for x := range width {
			p := Pos{x, y}
			if obstacles[p] > 0 {
				fmt.Printf("#")
			} else if boxes[p] > 0 {
				fmt.Printf("O")
			} else if robot == p {
				fmt.Printf("@")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Printf("\n")
	}
}

func move_box(from Pos, move_by Pos) bool {
	new_pos := from.Add(move_by)

	if obstacles[new_pos] > 0 {
		return false
	}

	if boxes[new_pos] > 0 {
		can_move := move_box(new_pos, move_by)

		if can_move {
			boxes[from] -= 1
			boxes[new_pos] += 1
		}
		return can_move
	}

	boxes[from] -= 1
	boxes[new_pos] += 1
	return true
}

func move_bot(move_by Pos) {
	new_pos := robot.Add(move_by)

	if obstacles[new_pos] <= 0 && boxes[new_pos] <= 0 {
		robot = new_pos

	} else if boxes[new_pos] > 0 && obstacles[new_pos] <= 0 {
		moved := move_box(new_pos, move_by)
		if moved {
			robot = new_pos
		}
	}
}

func p1() {
	printmap()

	for _, command := range commands {
		move_bot(command)
		// fmt.Printf(col.Blue+"%v  \n"+col.Reset, command)
		// printmap()
		// fmt.Printf("-----\n")
	}

	count := 0
	for key, val := range boxes {
		if val > 0 {
			count += 100 * (key[1])
			count += (key[0])
		}
	}
	fmt.Printf("Count: %d\n", count)

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 15" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	obstacles = map[Pos]int{}
	boxes = map[Pos]int{}

	width = len(lines[0])

	i := 0
	for {

		if i >= len(lines) {
			break
		}
		l := lines[i]
		if l == "" {
			break
		}

		for col, char := range l {
			switch string(char) {
			case "#":
				obstacles[Pos{col, i}] = 1
			case "O":
				boxes[Pos{col, i}] = 1
			case "@":
				robot = Pos{col, i}
			}
		}
		i += 1
	}

	height = i

	i += 1
	for {
		if i >= len(lines) {
			break
		}

		l := lines[i]
		for _, c := range l {
			switch c {
			case '<':
				commands = append(commands, Pos{-1, 0})
			case '>':
				commands = append(commands, Pos{1, 0})
			case '^':
				commands = append(commands, Pos{0, -1})
			case 'v':
				commands = append(commands, Pos{0, 1})
			}
		}
		i += 1
	}

	p1()

	//printmap()
	//fmt.Print(commands)
}

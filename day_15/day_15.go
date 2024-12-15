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
var boxes_l map[Pos]int
var boxes_r map[Pos]int
var robot Pos
var commands []Pos

func printmap() {
	for y := range height {
		for x := range width {
			p := Pos{x, y}
			if obstacles[p] > 0 {
				fmt.Printf("#")
			} else if boxes_l[p] > 0 {
				fmt.Printf("[")
			} else if boxes_r[p] > 0 {
				fmt.Printf("]")
			} else if robot == p {
				fmt.Printf("@")
			} else if boxes[p] > 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Printf("\n")
	}
}

func make_move(from Pos, to Pos) {
	boxes[from] -= 1
	boxes[to] += 1

	if boxes_l[from] > 0 {
		boxes_l[from] -= 1
		boxes_l[to] += 1
	}

	if boxes_r[from] > 0 {
		boxes_r[from] -= 1
		boxes_r[to] += 1
	}
}

func move_box(from Pos, move_by Pos, should_move bool) bool {
	new_pos := from.Add(move_by)

	// fmt.Printf("Checking: %v by %v\n", from, move_by)

	if obstacles[new_pos] > 0 {
		return false
	}

	// if we are moving left and right, the obstacles just push each other like before
	if move_by[1] == 0 {
		if boxes[new_pos] > 0 {
			can_move := move_box(new_pos, move_by, true)

			if can_move {

				make_move(from, new_pos)
			}
			return can_move
		}

		make_move(from, new_pos)

		return true

	} else {

		// I need to move the left or right side
		var other_move_from Pos
		var other_move_to Pos

		if boxes_l[from] > 0 {
			other_move_from = from.Add(Pos{1, 0})
		} else {
			other_move_from = from.Add(Pos{-1, 0})
		}
		other_move_to = other_move_from.Add(move_by)

		if obstacles[other_move_to] > 0 || obstacles[new_pos] > 0 {
			return false
		}

		if boxes[other_move_to] <= 0 && boxes[new_pos] <= 0 {
			if should_move {
				make_move(from, new_pos)
				make_move(other_move_from, other_move_to)
			}
			return true
		}

		can_move_l := move_box(new_pos, move_by, false)
		can_move_r := move_box(other_move_to, move_by, false)

		if can_move_l && can_move_r {
			move_box(new_pos, move_by, true)
			move_box(other_move_to, move_by, true)

			make_move(from, new_pos)
			make_move(other_move_from, other_move_to)
			return true
		}
	}

	return false
}

func move_bot(move_by Pos) {
	new_pos := robot.Add(move_by)

	if obstacles[new_pos] <= 0 && boxes[new_pos] <= 0 {
		robot = new_pos

	} else if boxes[new_pos] > 0 && obstacles[new_pos] <= 0 {
		moved := move_box(new_pos, move_by, true)
		if moved {
			robot = new_pos
		}
	}
}

func p1() {
	printmap()

	for _, command := range commands {
		move_bot(command)
		fmt.Printf(col.Blue+"%v  \n"+col.Reset, command)
		printmap()
		fmt.Printf("-----\n")
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
	boxes_l = map[Pos]int{}
	boxes_r = map[Pos]int{}

	width = len(lines[0]) * 2

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

			new_col := col * 2

			switch string(char) {
			case "#":
				obstacles[Pos{new_col, i}] = 1
				obstacles[Pos{new_col + 1, i}] = 1
			case "O":
				boxes[Pos{new_col, i}] = 1
				boxes[Pos{new_col + 1, i}] = 1
				boxes_l[Pos{new_col, i}] = 1
				boxes_r[Pos{new_col + 1, i}] = 1
			case "@":
				robot = Pos{new_col, i}
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

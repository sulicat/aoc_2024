package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

func radToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

type Obst [2]int
type Pos [2]int

var obstacles []Obst
var guard_pos Pos

var width int
var height int

var visited map[Pos]int
var visited_dirs map[Pos]int
var added_new_obst map[Pos]int

var lines []string

var dir int // in degrees, 0 is up, increasing clockwise

func check_obstacle(px int, py int) bool {
	for _, obst := range obstacles {
		if obst[0] == px && obst[1] == py {
			return true
		}
	}
	return false
}

func move_guard(guard_pos Pos, dir int) Pos {
	new_pos := guard_pos

	switch dir {
	case 0:
		new_pos[1] -= 1
	case 90:
		new_pos[0] += 1
	case 180:
		new_pos[1] += 1
	case 270:
		new_pos[0] -= 1
	case 360:
		new_pos[1] -= 1
	}

	return new_pos
}

func p1() {

	dir := 0
	count := 0
	visited = make(map[Pos]int)

	for {
		// break out if we leave
		if guard_pos[0] < 0 ||
			guard_pos[0] >= width ||
			guard_pos[1] < 0 ||
			guard_pos[1] >= height {
			break
		}

		// store the guard position
		_, has_visited := visited[guard_pos]
		if !has_visited {
			count += 1
		}
		visited[guard_pos] = 1

		// future guard
		future_guard_pos := move_guard(guard_pos, dir)
		// fmt.Printf("gaurd pos: %d\n", guard_pos)
		// fmt.Printf("future guard: %d\n", future_guard_pos)
		// fmt.Printf("dir: %d\n", dir)

		if check_obstacle(future_guard_pos[0], future_guard_pos[1]) {
			dir += 90
			dir = dir % 360
		} else {
			guard_pos = move_guard(guard_pos, dir)
		}

	}

	fmt.Printf("Count: %d\n", count)
}

func join_in_future(guard_pos Pos, dir int) bool {

	for {
		// break out if we leave
		if guard_pos[0] < 0 ||
			guard_pos[0] >= width ||
			guard_pos[1] < 0 ||
			guard_pos[1] >= height {
			break
		}

		if check_obstacle(guard_pos[0], guard_pos[1]) {
			return false
		}

		seen_dir, ok := visited_dirs[guard_pos]
		if ok && seen_dir == dir {
			return true
		}

		// direction is not correct, check if its 90 degree rotated and has an obstacle infron of it
		// this is to handle tight corners

		if ok && seen_dir == (dir+90)%360 {
			future_would_be := move_guard(guard_pos, (dir)%360)
			future_obst := check_obstacle(future_would_be[0], future_would_be[1])
			return future_obst
		}

		guard_pos = move_guard(guard_pos, dir)

	}

	return false
}

func print_map() {
	for r, line := range lines {
		for c := range line {

			_, is_visited := visited[Pos{c, r}]
			_, is_added := added_new_obst[Pos{c, r}]
			if is_added {
				fmt.Printf(col.BgBlue + "O" + col.Reset)
			} else if is_visited {
				fmt.Printf(col.Yellow + "X" + col.Reset)
			} else {
				fmt.Printf("%s", string(line[c]))
			}
		}
		fmt.Printf("\n")
	}
}

func p2() {

	dir := 0
	count := 0
	new_obst_count := 0
	visited = make(map[Pos]int)
	visited_dirs = make(map[Pos]int)
	added_new_obst = make(map[Pos]int)

	for {
		// break out if we leave
		if guard_pos[0] < 0 ||
			guard_pos[0] >= width ||
			guard_pos[1] < 0 ||
			guard_pos[1] >= height {
			break
		}

		// store the guard position
		_, has_visited := visited[guard_pos]
		if !has_visited {
			count += 1
		}
		visited[guard_pos] = 1
		visited_dirs[guard_pos] = dir

		print_map()

		// future guard
		future_guard_pos := move_guard(guard_pos, dir)
		// fmt.Printf("gaurd pos: %d\n", guard_pos)
		// fmt.Printf("future guard: %d\n", future_guard_pos)
		// fmt.Printf("dir: %d\n", dir)

		// if we placed an obstacle in front of us
		// do a walk and check for loops

		if check_obstacle(future_guard_pos[0], future_guard_pos[1]) {
			dir += 90
			dir = dir % 360
		} else {
			guard_pos = move_guard(guard_pos, dir)
		}

	}

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("New Obst count: %d\n", new_obst_count)

}

func main() {

	fmt.Printf(col.BgBlue + "Day 6" + col.Reset + "\n\n" + col.Reset)

	input_file, _ := os.ReadFile("./input.txt")
	//fmt.Printf("%s\n\n", input_file)
	lines = strings.Split(string(input_file), "\n")

	width = len(lines)
	height = len(lines[0])

	obstacles = make([]Obst, 0)
	guard_pos = Pos{0, 0}

	for r, line := range lines {
		for c := range line {
			if string(lines[r][c]) == "#" {
				obstacles = append(obstacles, Obst{c, r})
			}

			if string(lines[r][c]) == "^" {
				guard_pos = Pos{c, r}
			}
		}
	}

	//fmt.Printf("Obstacles: %d\n", obstacles)
	//fmt.Printf("Guard Pos: %d\n", guard_pos)

	//p1()
	p2()
}

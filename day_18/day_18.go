package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

const width = 7
const height = 7

type Pos struct {
	x    int
	y    int
	cost int
}

var obstacles map[Pos]int
var obstacle_positions []Pos
var start Pos
var goal Pos

func draw_map() {
	for y := range height {
		for x := range width {
			if obstacles[Pos{x, y}] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}

func p1() {
	cells := PosHeap{}
	heap.Push(cells, start)

	heap.Init(options)
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 16" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	obstacle_positions = []Pos{}
	obstacles = map[Pos]int{}
	start = Pos{0, 0}
	goal = Pos{6, 6}

	for i, l := range lines {
		options := strings.Split(l, ",")
		x, _ := strconv.Atoi(options[0])
		y, _ := strconv.Atoi(options[1])

		p := Pos{x, y}
		fmt.Printf("%v\n", p)
		obstacles[p] = 1
		obstacle_positions = append(obstacle_positions, p)

		if i >= 11 {
			break
		}
	}

	fmt.Printf("%v\n", obstacles)

	draw_map()

	p1()
}

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

func (a Pos) Same(b Pos) bool {
	return a[0] == b[0] && a[1] == b[1]
}

type PosDir struct {
	Pos Pos
	Dir int // 0 deg is east, +90 is south
}

// -------------------------------------------------------------------------------------------------------

var world map[Pos]int // > 0 == obstacle

var player PosDir
var goal PosDir
var width int
var height int

func isWall(x, y int) bool {
	return world[Pos{x, y}] > 0
}

func printmap() {
	for y := range height {
		for x := range width {
			if player.Pos.Same(Pos{x, y}) {
				fmt.Printf("S")
			} else if goal.Pos.Same(Pos{x, y}) {
				fmt.Printf("E")
			} else if isWall(x, y) {
				fmt.Printf("o")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 16" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")
	width = len(lines[0])
	height = len(lines)
	world = map[Pos]int{}

	for r, line := range lines {
		for c, char := range line {
			if string(char) == "#" {
				world[Pos{c, r}] = 1
			} else {
				world[Pos{c, r}] = 0
			}

			if string(char) == "S" {
				player = PosDir{Pos{c, r}, 0}
			}

			if string(char) == "E" {
				goal = PosDir{Pos{c, r}, 0}
			}

		}
	}

	printmap()
}

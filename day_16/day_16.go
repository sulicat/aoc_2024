package main

import (
	"fmt"
	"math"
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

type Cell struct {
	is_obstacle bool
	visited     bool
	score       int
}

// -------------------------------------------------------------------------------------------------------

var world map[Pos]Cell // > 0 == obstacle
var empty_nodes []Pos

var player PosDir
var goal PosDir
var width int
var height int

func isWall(x, y int) bool {
	return world[Pos{x, y}].is_obstacle
}

func getScore(p Pos) int {
	v, ok := world[p]
	if ok {
		return v.score
	}
	return math.MaxInt32
}

func setScore(p Pos, u int) {
	v, ok := world[p]
	if ok {
		v.score = u
		world[p] = v
	}
}

func setVisited(p Pos, u bool) {
	v, ok := world[p]
	if ok {
		v.visited = u
		world[p] = v
	}
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

func get_free_neighbors(p Pos) []Pos {
	out := []Pos{}

	p1_p := Pos{p[0] + 1, p[1]}
	p1, exists := world[p1_p]
	if exists && !p1.is_obstacle {
		out = append(out, p1_p)
	}

	p2_p := Pos{p[0] - 1, p[1]}
	p2, exists := world[p2_p]
	if exists && !p2.is_obstacle {
		out = append(out, p2_p)
	}

	p3_p := Pos{p[0], p[1] + 1}
	p3, exists := world[p3_p]
	if exists && !p3.is_obstacle {
		out = append(out, p3_p)
	}

	p4_p := Pos{p[0], p[1] - 1}
	p4, exists := world[p4_p]
	if exists && !p4.is_obstacle {
		out = append(out, p4_p)
	}

	return out
}

func min_score(sol []PosDir, visited map[Pos]int) {
	min := math.MaxInt32
	for _, s := range sol {
		if visited[s.Pos] <= 0 {

		}
	}
}

func p1() {
	unvisted := empty_nodes
	for _, u := range unvisted {
		setScore(u, math.MaxInt32)
		setVisited(u, false)
	}

	solution := []PosDir{}
	solution = append(solution, player)
	visited := map[Pos]int{}

	setScore(current, 0)
	for {

		i := min_score(solution, visited)

		break
	}

	fmt.Printf(" %v -> %v\n", current, get_free_neighbors(current))

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 16" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")
	width = len(lines[0])
	height = len(lines)
	world = map[Pos]Cell{}
	empty_nodes = []Pos{}

	for r, line := range lines {
		for c, char := range line {
			if string(char) == "#" {
				world[Pos{c, r}] = Cell{is_obstacle: true, visited: false, score: math.MaxInt32}
			} else {
				world[Pos{c, r}] = Cell{is_obstacle: false, visited: false, score: math.MaxInt32}
				empty_nodes = append(empty_nodes, Pos{c, r})
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

	fmt.Printf("\n\n")

	p1()
}

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
	from        Pos
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

}

func pop_min_score(arr *[]Pos) Pos {
	min_s := math.MaxInt32
	min_i := 0
	for i, p := range *arr {
		if getScore(p) < min_s {
			min_s = getScore(p)
			min_i = i
		}
	}

	out := (*arr)[min_i]

	*arr = append((*arr)[:min_i], (*arr)[min_i+1:]...)
	return out
}

func calculate_score(src Pos, dst Pos) int {
	return 1
}

func p1() {
	unvisted := empty_nodes
	for _, u := range unvisted {
		setScore(u, math.MaxInt32)
		setVisited(u, false)
	}

	// visited := map[Pos]int{}
	look_at := []Pos{}
	look_at = append(look_at, player.Pos)
	setScore(player.Pos, 0)

	for {
		if len(look_at) <= 0 {
			break
		}

		min_pos := pop_min_score(&look_at)
		fmt.Printf("%v\n", min_pos)

		adj := get_free_neighbors(min_pos)
		for _, adj_p := range adj {
			if !world[adj_p].visited {
				setScore(adj_p, getScore(min_pos)+calculate_score(min_pos, adj_p))
				look_at = append(look_at, adj_p)
				fmt.Printf("score: %d\n", getScore(adj_p))
			}
		}

		setVisited(min_pos, true)

	}

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

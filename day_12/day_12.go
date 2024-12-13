package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

var data [][]string
var groups []Group
var group_maps []map[Pos]int
var group_fences []map[Pos]Fence

var visited map[Pos]int

type Fence struct {
	Top    bool
	Right  bool
	Bottom bool
	Left   bool
	P      Pos
}

type Pos struct {
	R int
	C int
}

type Group []Pos

func get(r, c int) (Pos, error) {

	if r < 0 || c < 0 || r >= len(data) || c >= len(data[0]) {
		return Pos{0, 0}, errors.New("no cell here")
	}

	return Pos{r, c}, nil
}

func get_adj(p Pos) []Pos {
	out := make([]Pos, 0, 4)

	neighbor, err := get(p.R-1, p.C)
	if err == nil {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R+1, p.C)
	if err == nil {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R, p.C-1)
	if err == nil {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R, p.C+1)
	if err == nil {
		out = append(out, neighbor)
	}

	return out
}

func get_same_adj(p Pos) []Pos {
	out := make([]Pos, 0, 4)

	letter := data[p.R][p.C]

	neighbor, err := get(p.R-1, p.C)
	if err == nil && data[p.R-1][p.C] == letter {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R+1, p.C)
	if err == nil && data[p.R+1][p.C] == letter {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R, p.C-1)
	if err == nil && data[p.R][p.C-1] == letter {
		out = append(out, neighbor)
	}

	neighbor, err = get(p.R, p.C+1)
	if err == nil && data[p.R][p.C+1] == letter {
		out = append(out, neighbor)
	}

	return out
}

func recurse_links(p Pos) []Pos {
	out := make([]Pos, 0)

	visited[p] = 1
	out = append(out, p)

	adj := get_same_adj(p)
	for _, a := range adj {
		_, has_visited := visited[a]
		if !has_visited {
			out = append(out, recurse_links(a)...)
		}
	}

	return out
}

func count_fences(p Pos, groupmap map[Pos]int) int {
	out := 4
	dirs := []Pos{
		{p.R - 1, p.C},
		{p.R + 1, p.C},
		{p.R, p.C - 1},
		{p.R, p.C + 1},
	}

	for _, d := range dirs {
		_, inmap := groupmap[d]
		if inmap {
			out -= 1
		}
	}
	return out
}

func update_fences(p Pos, groupmap map[Pos]int, fencemap *map[Pos]Fence) int {
	out := 4
	dirs := []Pos{
		{p.R - 1, p.C},
		{p.R + 1, p.C},
		{p.R, p.C - 1},
		{p.R, p.C + 1},
	}

	new_fence := Fence{}
	new_fence.P = p

	for di, d := range dirs {
		_, inmap := groupmap[d]
		if inmap {
			out -= 1
		} else {

			switch di {
			case 0:
				new_fence.Top = true
			case 1:
				new_fence.Bottom = true
			case 2:
				new_fence.Left = true
			case 3:
				new_fence.Right = true
			}

		}

	}

	(*fencemap)[p] = new_fence

	return out
}

func count_island(in map[Pos]bool) int {
	out := 0

	rvis := map[Pos]int{}

	var dfs func(p Pos)
	dfs = func(p Pos) {

		rvis[p] = 1
		adj := get_adj(p)
		for _, a := range adj {

			_, already_visited := rvis[a]
			_, exists := in[a]
			if exists &&
				!already_visited && in[a] {

				dfs(a)
			}
		}
	}

	for pos := range in {
		_, iv := rvis[pos]
		if !iv && in[pos] {
			dfs(pos)
			out += 1
		}

	}

	return out
}

func get_edges(gi int) int {
	// we will dfs to find all bot, top, left and right fences

	// dfs_visited := map[Pos]int{}

	top_fences := map[Pos]bool{}
	bot_fences := map[Pos]bool{}
	left_fences := map[Pos]bool{}
	right_fences := map[Pos]bool{}

	for _, p := range group_fences[gi] {
		top_fences[p.P] = p.Top
		bot_fences[p.P] = p.Bottom
		left_fences[p.P] = p.Left
		right_fences[p.P] = p.Right
	}

	top_count := count_island(top_fences)
	left_count := count_island(left_fences)
	right_count := count_island(right_fences)
	bot_count := count_island(bot_fences)

	return top_count + left_count + right_count + bot_count
}

func p1() {

	out := 0

	for r := range data {
		for c := range data[r] {

			_, is_visited := visited[Pos{r, c}]
			if !is_visited {
				cells := recurse_links(Pos{r, c})

				groups = append(groups, cells)
				group_map := make(map[Pos]int)
				for _, cell := range cells {
					group_map[Pos{cell.R, cell.C}] = 1
				}
				group_maps = append(group_maps, group_map)
			}
		}
	}

	for gi, g := range groups {
		fmt.Printf("Group: %d len: %d\n", gi, len(g))
		fmt.Printf("map: %v\n ", group_maps[gi])

		sum_fences := 0
		for _, p := range g {
			// fmt.Printf("%s ", data[p.R][p.C])
			sum_fences += count_fences(p, group_maps[gi])
		}
		out += sum_fences * len(g)
		fmt.Printf("%d * %d = %d", len(g), sum_fences, sum_fences*len(g))
		fmt.Printf("\n")
	}

	fmt.Printf("TOTAL p1: %d\n", out)
}

func p2() {

	out := 0

	for r := range data {
		for c := range data[r] {

			_, is_visited := visited[Pos{r, c}]
			if !is_visited {
				cells := recurse_links(Pos{r, c})

				groups = append(groups, cells)
				group_map := make(map[Pos]int)
				for _, cell := range cells {
					group_map[Pos{cell.R, cell.C}] = 1
				}
				group_maps = append(group_maps, group_map)

				group_fence := make(map[Pos]Fence)
				group_fences = append(group_fences, group_fence)
			}
		}
	}

	for gi, g := range groups {
		fmt.Printf("Group: %d %s len: %d\n", gi, data[g[0].R][g[0].C], len(g))
		fmt.Printf("map: %v\n ", group_maps[gi])

		sum_fences := 0
		for _, p := range g {
			sum_fences += update_fences(p, group_maps[gi], &group_fences[gi])
		}
		// out += sum_fences * len(g)
		// fmt.Printf("%d * %d = %d", len(g), sum_fences, sum_fences*len(g))
		// fmt.Printf("\n")

		edges := get_edges(gi)
		fmt.Printf("EDGES %d => %d\n", edges, edges*len(g))
		fmt.Printf("\n")
		out += edges * len(g)
	}

	fmt.Printf("total p2: %d\n", out)

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 10" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	data = make([][]string, len(lines))

	for r, l := range lines {
		data[r] = make([]string, len(l))
		for c, char := range l {
			data[r][c] = string(char)
		}
	}

	// fmt.Printf("%v\n", data)

	visited = make(map[Pos]int)

	group_maps = make([]map[Pos]int, 0)
	group_fences = make([]map[Pos]Fence, 0)

	// p1()
	p2()
}

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
var visited map[Pos]int

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

func p1() {

	// groups = append(groups, Group{})
	// current_group := &groups[len(groups)-1]
	// current_group.Members = make([]Pos, 0)

	for r := range data {
		for c := range data[r] {

			_, is_visited := visited[Pos{r, c}]
			if !is_visited {

				cells := recurse_links(Pos{r, c})
				groups = append(groups, cells)
				fmt.Printf("%v\n\n", cells)
			}
		}
	}

	for gi, g := range groups {
		fmt.Printf("Group: %d\n", gi)

		for _, p := range g {
			fmt.Printf("%s ", data[p.R][p.C])
		}
		fmt.Printf("\n")
	}

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

	fmt.Printf("%v\n", data)

	visited = make(map[Pos]int)

	p1()

}

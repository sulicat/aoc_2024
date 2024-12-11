package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

type Pos [2]int

var data [][]int
var trailheads []Pos

var num_rows int
var num_cols int

var dirs = [...]Pos{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p1 *Pos) Add(p2 Pos) Pos {
	return Pos{p1[0] + p2[0], p1[1] + p2[1]}
}

func (p1 *Pos) Inside() bool {
	return p1[0] >= 0 &&
		p1[0] < num_cols &&
		p1[1] >= 0 &&
		p1[1] < num_rows
}

var p1_count int
var counted map[Pos]int

func walk(p Pos) int {

	current_alt := data[p[0]][p[1]]

	count := 0
	for _, d := range dirs {

		new := p.Add(d)
		if new.Inside() {

			new_alt := data[new[0]][new[1]]

			if new_alt != current_alt+1 {
				continue
			}

			fmt.Printf("%v:%d -> %v%d\n", p, current_alt, new, new_alt)

			if new_alt == 9 {

				_, already_counted := counted[new]
				if !already_counted {
					p1_count += 1
				}
				counted[new] = 1

			} else {
				walk(new)
			}

		}
	}
	return count
}

func walk2(p Pos) int {

	current_alt := data[p[0]][p[1]]

	count := 0
	for _, d := range dirs {

		new := p.Add(d)
		if new.Inside() {

			new_alt := data[new[0]][new[1]]

			if new_alt != current_alt+1 {
				continue
			}

			fmt.Printf("%v:%d -> %v%d\n", p, current_alt, new, new_alt)

			if new_alt == 9 {
				p1_count += 1

			} else {
				walk2(new)
			}

		}
	}
	return count
}

func count_trail(start Pos) {
	walk2(start)
}

func p1() {
	total := 0

	// for every trailhead start walking
	for _, th := range trailheads {
		// counted = make(map[Pos]int)
		p1_count = 0
		count_trail(th)
		total += p1_count
	}

	fmt.Printf("Trails p1: %d\n", total)
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 10" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	num_rows = len(lines)
	num_cols = len(lines[0])
	data = make([][]int, num_rows)

	for r := range lines {
		data[r] = make([]int, num_cols)
		for c := range lines[r] {

			n, _ := strconv.Atoi(string(lines[r][c]))
			data[r][c] = n
			if n == 0 {
				trailheads = append(trailheads, Pos{r, c})
			}
		}
	}

	p1()
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

// const width = 11
// const height = 7

const width = 101
const height = 103

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

type Robot struct {
	Pos Pos
	Vel Pos
}

var robots []Robot
var botCount map[Pos]int

func printmap() {

	for y := range height {
		for x := range width {
			num_bots := botCount[Pos{x, y}]
			if num_bots > 0 {
				fmt.Printf("%d", num_bots)

			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func step() {
	for i := range robots {
		r := &robots[i]

		botCount[r.Pos] -= 1
		r.Pos = r.Pos.Add(r.Vel)

		// if the position is negative add width/height till its positive
		for {
			if r.Pos[0] >= 0 {
				break
			}
			r.Pos[0] += width
		}
		for {
			if r.Pos[1] >= 0 {
				break
			}
			r.Pos[1] += height
		}

		r.Pos[0] = r.Pos[0] % width
		r.Pos[1] = r.Pos[1] % height
		botCount[r.Pos] += 1

	}
}

func p1_count() int {
	mid_x := width / 2
	mid_y := height / 2

	q1_c := 0
	q2_c := 0
	q3_c := 0
	q4_c := 0

	for x := 0; x < mid_x; x++ {
		for y := 0; y < mid_y; y++ {
			bc := botCount[Pos{x, y}]
			if bc > 0 {
				// fmt.Printf("add: q1 %d %d   %d\n", x, y, bc)
				q1_c += bc
			}
		}
	}

	for x := mid_x + 1; x < width; x++ {
		for y := 0; y < mid_y; y++ {
			bc := botCount[Pos{x, y}]
			if bc > 0 {
				// fmt.Printf("add: q2 %d %d   %d\n", x, y, bc)
				q2_c += bc
			}
		}
	}

	for x := mid_x + 1; x < width; x++ {
		for y := mid_y + 1; y < height; y++ {
			bc := botCount[Pos{x, y}]
			if bc > 0 {
				// fmt.Printf("add: q3 %d %d   %d\n", x, y, bc)
				q3_c += bc
			}
		}
	}

	for x := 0; x < mid_x; x++ {
		for y := mid_y + 1; y < height; y++ {
			bc := botCount[Pos{x, y}]
			if bc > 0 {
				// fmt.Printf("add: q4 %d %d   %d\n", x, y, bc)
				q4_c += bc
			}
		}
	}

	return q1_c * q2_c * q3_c * q4_c
}

/*func p1() {

	for i := 0; i < 2000; i++ {
		fmt.Printf("STEP: %d\n", i)
		step()
		if i > 999 {
			printmap()
		}
		fmt.Printf("\n")
	}

	printmap()
	out := p1_count()
	fmt.Printf("p1 count; %d\n", out)
}*/

func check_tree() bool {
	// I will assume the tree has a trunc
	// I will check if I see a line of N dimension bots

	good := false

	trunc_height := 10

	for x := 0; x < width; x++ {
		for y := 0; y < height-trunc_height; y++ {

			found := true
			for off := 0; off <= trunc_height; off += 1 {
				if botCount[Pos{x, y + off}] <= 0 {
					found = false
				}
			}
			if found {
				good = true
				break
			}

		}
	}

	return good
}

func p1() {

	i := 0
	for {
		step()
		good := check_tree()
		if good {
			break
		}
		i += 1
	}

	fmt.Printf("STEP: %d\n", i)
	printmap()
	fmt.Printf("\n")
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 14" + col.Reset + "\n")

	robots = []Robot{}
	botCount = map[Pos]int{}

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	for _, l := range lines {
		sp := strings.Split(l, " ")
		p_str := strings.Split(sp[0], "p=")[1]
		v_str := strings.Split(sp[1], "v=")[1]

		p_str_sp := strings.Split(p_str, ",")
		p_x, _ := strconv.Atoi(p_str_sp[0])
		p_y, _ := strconv.Atoi(p_str_sp[1])

		v_str_sp := strings.Split(v_str, ",")
		v_x, _ := strconv.Atoi(v_str_sp[0])
		v_y, _ := strconv.Atoi(v_str_sp[1])

		var r = Robot{}
		r.Pos = Pos{p_x, p_y}
		r.Vel = Pos{v_x, v_y}
		robots = append(robots, r)

		botCount[r.Pos] += 1
	}

	fmt.Printf("robots:\n%v\n", robots)

	p1()
}

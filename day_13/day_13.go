package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

type Pos [2]int

type Machine struct {
	A     Pos
	B     Pos
	Prize Pos
}

func (a Pos) Add(b Pos) Pos {
	return Pos{
		a[0] + b[0],
		a[1] + b[1]}
}

var machines []Machine

func recurse(start_pos Pos, m Machine) (int, bool) {

	// fmt.Printf("Pos: %v\n", start_pos)

	if start_pos == m.Prize {
		// we are at the goal, no button to press
		// fmt.Printf("Goal\n")
		return 0, true

	} else if start_pos[0] > m.Prize[0] || start_pos[1] > m.Prize[1] {

		// fmt.Printf("FAIL\n")
		return 0, false

	} else {

		// we need to press a button
		a_pos := start_pos.Add(m.A)
		a_cost, a_winnable := recurse(a_pos, m)
		a_cost += 3 // we pressed a

		b_pos := start_pos.Add(m.B)
		b_cost, b_winnable := recurse(b_pos, m)
		b_cost += 1 // we pressed a

		if !a_winnable && !b_winnable {
			return -1, false
		}

		if a_winnable && !b_winnable {
			return a_cost, true
		}

		if !a_winnable && b_winnable {
			return b_cost, true
		}

		return int(math.Min(
				float64(a_cost),
				float64(b_cost))),
			true

	}
}

func compute(m Machine) (int, bool) {
	current_pos := Pos{0, 0}
	return recurse(current_pos, m)
}

func p1() {

	for _, m := range machines {
		fmt.Printf("Doing math for machine: %v\n", m)

		cost, winnable := compute(m)
		fmt.Printf("Cost: %v   winnable: %v\n", cost, winnable)

	}

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 13" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	m := Machine{}

	for r, l := range lines {
		iteration := r % 4

		if iteration == 0 {
			sa := strings.Split(l, "Button A: X+")[1]
			s_vals := strings.Split(sa, ", Y+")
			v_x, _ := strconv.Atoi(s_vals[0])
			v_y, _ := strconv.Atoi(s_vals[1])

			m.A[0] = v_x
			m.A[1] = v_y

		} else if iteration == 1 {
			sa := strings.Split(l, "Button B: X+")[1]
			s_vals := strings.Split(sa, ", Y+")
			v_x, _ := strconv.Atoi(s_vals[0])
			v_y, _ := strconv.Atoi(s_vals[1])

			m.B[0] = v_x
			m.B[1] = v_y

		} else if iteration == 2 {
			sa := strings.Split(l, "Prize: X=")[1]
			s_vals := strings.Split(sa, ", Y=")
			v_x, _ := strconv.Atoi(s_vals[0])
			v_y, _ := strconv.Atoi(s_vals[1])

			m.Prize[0] = v_x
			m.Prize[1] = v_y

			machines = append(machines, m)
		}
	}

	fmt.Printf("Machines: %v\n\n", machines)

	p1()

}

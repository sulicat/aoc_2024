package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"gonum.org/v1/gonum/mat"
)

type Pos [2]int

type Machine struct {
	A     Pos
	B     Pos
	Prize Pos
}

type CacheVal struct {
	Cost     int
	Winnable bool
}

func (a Pos) Add(b Pos) Pos {
	return Pos{
		a[0] + b[0],
		a[1] + b[1]}
}

var machines []Machine
var cache map[Pos]CacheVal // cache of pos to int

func recurse(start_pos Pos, m Machine) (int, bool) {

	// fmt.Printf("Pos: %v\n", start_pos)

	cached, in_cache := cache[start_pos]
	if in_cache {
		return cached.Cost, cached.Winnable
	}

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

		if !a_winnable {

		}

		b_pos := start_pos.Add(m.B)
		b_cost, b_winnable := recurse(b_pos, m)
		b_cost += 1 // we pressed a

		if !a_winnable && !b_winnable {

			cache[start_pos] = CacheVal{-1, false}
			return -1, false
		}

		if a_winnable && !b_winnable {
			cache[start_pos] = CacheVal{a_cost, true}
			return a_cost, true
		}

		if !a_winnable && b_winnable {
			cache[start_pos] = CacheVal{b_cost, true}
			return b_cost, true
		}

		min := int(math.Min(
			float64(a_cost),
			float64(b_cost)))

		cache[start_pos] = CacheVal{min, true}
		return min, true

	}
}

func compute(m Machine) (int, bool) {
	// current_pos := Pos{0, 0}
	// cache = map[Pos]CacheVal{}
	// return recurse(current_pos, m)

	a := mat.NewDense(2, 2, []float64{
		float64(m.A[0]), float64(m.B[0]),
		float64(m.A[1]), float64(m.B[1]),
	})

	var inv mat.Dense
	inv.Inverse(a)

	b := mat.NewDense(2, 1, []float64{
		float64(m.Prize[0]),
		float64(m.Prize[1]),
	})
	b.Mul(&inv, b)

	// fmt.Println(mat.Formatted(a, mat.Prefix("    ")))
	// fmt.Println(mat.Formatted(&inv, mat.Prefix("    ")))
	// fmt.Println(mat.Formatted(b, mat.Prefix("    ")))

	ac_f := b.At(0, 0)
	bc_f := b.At(1, 0)

	diff_a := math.Abs(math.Round(ac_f) - ac_f)
	diff_b := math.Abs(math.Round(bc_f) - bc_f)

	winnable := diff_a < 0.0001 && diff_b < 0.0001

	cost := 3*math.Round(b.At(0, 0)) + math.Round(b.At(1, 0))

	return int(cost), winnable

}

func p1() {

	out := 0
	for _, m := range machines {
		fmt.Printf("Doing math for machine: %v\n", m)

		cost, winnable := compute(m)
		fmt.Printf("Cost: %v   winnable: %v\n", cost, winnable)
		if winnable {
			out += cost
		}
	}

	fmt.Printf("\nTOTAL COST: %d\n", out)

}

func p2() {

	// a := mat.NewDense(2, 2, []float64{26, 66, 67, 21})
	// var inv mat.Dense
	// inv.Inverse(a)

	// b := mat.NewDense(2, 1, []float64{12748, 12176})
	// b.Mul(&inv, b)

	// fmt.Println(mat.Formatted(a, mat.Prefix("    ")))
	// fmt.Println(mat.Formatted(&inv, mat.Prefix("    ")))
	// fmt.Println(mat.Formatted(b, mat.Prefix("    ")))

	// fmt.Printf("---------------------------\n")

	out := 0
	for _, m := range machines {
		fmt.Printf("Doing math for machine: %v\n", m)

		cost, winnable := compute(m)
		fmt.Printf("Cost: %v   winnable: %v\n", cost, winnable)
		if winnable {
			out += cost
		}
	}

	fmt.Printf("\nTOTAL COST: %d\n", out)

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

			m.Prize[0] = v_x + 10000000000000
			m.Prize[1] = v_y + 10000000000000

			machines = append(machines, m)
		}
	}

	fmt.Printf("Machines: %v\n\n", machines)

	// p1()
	p2()
}

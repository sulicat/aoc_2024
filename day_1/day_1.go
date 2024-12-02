package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	utils "github.com/sulicat/goboi/goboi_utils"
)

func main_p1() {
	fmt.Printf(col.Red + "Day 1" + col.Reset + "\n")

	input_file, err := os.ReadFile("input.txt")
	utils.PanicOnErr(err)

	lines := strings.Split(string(input_file), "\n")

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	for i := 0; i < len(lines); i++ {
		l := lines[i]
		words := strings.Split(l, "   ")
		if len(words) > 1 {
			v1, _ := strconv.Atoi(words[0])
			v2, _ := strconv.Atoi(words[1])

			a1 = append(a1, v1)
			a2 = append(a2, v2)
		}
	}

	sort.Ints(a1)
	sort.Ints(a2)

	total := 0
	for i := 0; i < len(a1); i++ {
		fmt.Printf("%d  ~~ %d\n", a1[i], a2[i])
		total += int(math.Abs(float64(a1[i] - a2[i])))
	}

	fmt.Printf("TOTAL: %d\n", total)

}

func main_p2() {

	fmt.Printf(col.Red + "Day 1" + col.Reset + "\n")

	input_file, err := os.ReadFile("input.txt")
	utils.PanicOnErr(err)

	lines := strings.Split(string(input_file), "\n")

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	for i := 0; i < len(lines); i++ {
		l := lines[i]
		words := strings.Split(l, "   ")
		if len(words) > 1 {
			v1, _ := strconv.Atoi(words[0])
			v2, _ := strconv.Atoi(words[1])

			a1 = append(a1, v1)
			a2 = append(a2, v2)
		}
	}

	sort.Ints(a1)
	sort.Ints(a2)

	// first build a map of apearances
	a2_count := make(map[int]int)
	for i := 0; i < len(a2); i++ {
		a2_count[a2[i]] += 1
	}

	total := 0
	for i := 0; i < len(a1); i++ {

		v1 := a1[i]
		v2 := a2_count[v1]

		fmt.Printf("%d  ~~ %d\n", a1[i], a2_count[a1[i]])
		total += int(v1 * v2)
	}

	fmt.Printf("TOTAL: %d\n", total)

}

func main() {
	main_p2()
}

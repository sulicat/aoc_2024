package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
	//utils "github.com/sulicat/goboi/utils"
)

func p1() {
	fmt.Printf("%s\n", col.BrightBlue+"Day 3"+col.Reset)

	file, _ := os.ReadFile("./input.txt")
	//lines := strings.Split(string(file), "\n")
	//input := make([][]int, 0)

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}?\)`)
	matches := r.FindAllString(string(file), -1)

	sum := 0
	for _, m := range matches {
		fmt.Printf("match: %s\n", m)

		num_1 := m[4:]
		num_1 = num_1[:len(num_1)-1]

		vals := strings.Split(num_1, ",")
		sum += utils.First(strconv.Atoi(vals[0])) * utils.First(strconv.Atoi(vals[1]))

	}

	fmt.Printf("sum: %d\n", sum)

}

func p2() {

	fmt.Printf("%s\n", col.BrightBlue+"Day 3"+col.Reset)

	file, _ := os.ReadFile("./input.txt")
	//lines := strings.Split(string(file), "\n")
	//input := make([][]int, 0)

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}?\)|do.*?\(?\)`)
	matches := r.FindAllString(string(file), -1)

	state := 1
	sum := 0
	for _, m := range matches {
		fmt.Printf("match: %s\n", m)

		if m == "do()" {
			state = 1
		} else if m == "don't()" {
			state = 0
		} else if state == 1 {

			num_1 := m[4:]
			num_1 = num_1[:len(num_1)-1]

			vals := strings.Split(num_1, ",")
			sum += utils.First(strconv.Atoi(vals[0])) * utils.First(strconv.Atoi(vals[1]))
		}
	}

	fmt.Printf("sum: %d\n", sum)

}

func main() {
	p2()
}

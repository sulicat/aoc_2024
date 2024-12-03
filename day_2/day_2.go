package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

func is_safe(row []int) int {
	if row[1] == row[0] {
		return 0
	}

	is_increasing := row[1]-row[0] > 0

	good := true
	for i := 1; i < len(row); i++ {
		if is_increasing && row[i]-row[i-1] <= 0 {
			good = false
		}

		if !is_increasing && row[i]-row[i-1] >= 0 {
			good = false
		}

		diff := math.Abs(float64(row[i] - row[i-1]))
		if diff < 1 || diff > 3 {
			good = false
		}
	}

	if good {
		return 1
	}
	return 0

}

func main_p1() {
	fmt.Printf(col.Red + "hello world\n" + col.Reset)

	file, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	input := make([][]int, 0)

	for line_i := 0; line_i < len(lines); line_i++ {
		line := lines[line_i]

		arr := strings.Split(line, " ")

		int_arr := make([]int, 0)
		for i := 0; i < len(arr); i++ {
			n, _ := strconv.Atoi(arr[i])
			int_arr = append(int_arr, n)
		}
		input = append(input, int_arr)
	}

	total := 0
	for row_i := range input {

		row := input[row_i]
		total += is_safe(row)
	}
	fmt.Printf(col.Green+"Total: %d\n"+col.Reset, total)

}

func main_p2() {
	file, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")
	input := make([][]int, 0)

	for line_i := 0; line_i < len(lines); line_i++ {
		line := lines[line_i]

		arr := strings.Split(line, " ")

		int_arr := make([]int, 0)
		for i := 0; i < len(arr); i++ {
			n, _ := strconv.Atoi(arr[i])
			int_arr = append(int_arr, n)
		}
		input = append(input, int_arr)
	}

	total := 0
	for row_i := range input {
		row := input[row_i]
		safe := is_safe(row)

		if safe == 0 {

			for i := range row {
				copy_row := make([]int, len(row))
				copy(copy_row, row)
				new_row := append(copy_row[:i], copy_row[i+1:]...)
				safe = is_safe(new_row)

				if safe == 1 {
					break
				}
			}
		}

		fmt.Print("\n\n")

		total += safe
	}

	fmt.Printf(col.Green+"Total: %d\n"+col.Reset, total)

}

func main() {
	main_p2()
}

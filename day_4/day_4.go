package main

import (
	"fmt"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

var input_file []byte
var height int
var width int
var input_data [][]string
var X_positions [][2]int
var A_positions [][2]int

func word_from_pos(vals ...[2]int) string {
	out := ""
	for _, v := range vals {
		out += input_data[v[0]][v[1]]
	}
	return out
}

func count_adjacent(r int, c int, compare_str string) int {
	wordlen := len(compare_str) - 1
	total := 0
	if r-wordlen >= 0 {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r - 1, c},
			[2]int{r - 2, c},
			[2]int{r - 3, c})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	if r+wordlen < height {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r + 1, c},
			[2]int{r + 2, c},
			[2]int{r + 3, c})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	if c+wordlen < width {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r, c + 1},
			[2]int{r, c + 2},
			[2]int{r, c + 3})

		if check == compare_str {
			total += 1
			fmt.Printf(col.Green+"%d %d\n"+col.Reset, r, c)
		}
	}

	if c-wordlen >= 0 {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r, c - 1},
			[2]int{r, c - 2},
			[2]int{r, c - 3})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	return total
}

func count_diag(r int, c int, compare_str string) int {
	wordlen := len(compare_str) - 1
	total := 0
	if r-wordlen >= 0 && c-wordlen >= 0 {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r - 1, c - 1},
			[2]int{r - 2, c - 2},
			[2]int{r - 3, c - 3})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	if r+wordlen < height && c+wordlen < width {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r + 1, c + 1},
			[2]int{r + 2, c + 2},
			[2]int{r + 3, c + 3})

		if check == compare_str {
			total += 1
			fmt.Printf(col.Red+"%d %d\n"+col.Reset, r, c)
		}
	}

	if r-wordlen >= 0 && c+wordlen < width {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r - 1, c + 1},
			[2]int{r - 2, c + 2},
			[2]int{r - 3, c + 3})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	if r+wordlen < height && c-wordlen >= 0 {
		check := word_from_pos(
			[2]int{r, c},
			[2]int{r + 1, c - 1},
			[2]int{r + 2, c - 2},
			[2]int{r + 3, c - 3})

		if check == compare_str {
			total += 1
			fmt.Printf("%d %d\n", r, c)
		}
	}

	return total
}

func p1() {

	total := 0

	// find all X's then go through the iterations
	for _, pos := range X_positions {
		r := pos[0]
		c := pos[1]

		total += count_adjacent(r, c, "XMAS")
		total += count_diag(r, c, "XMAS")

	}

	fmt.Printf(col.Green+"TOTAL: %d\n"+col.Reset, total)
}

func count_mas(r, c int) int {
	out := 0

	if r >= 1 && r < height-1 {
		if c >= 1 && c < width-1 {

			w1 := word_from_pos(
				[2]int{r - 1, c - 1},
				[2]int{r, c},
				[2]int{r + 1, c + 1})

			w2 := word_from_pos(
				[2]int{r - 1, c + 1},
				[2]int{r, c},
				[2]int{r + 1, c - 1})

			if w1 == "MAS" || w1 == "SAM" {
				if w2 == "MAS" || w2 == "SAM" {
					out += 1
				}
			}

		}
	}

	return out
}

func p2() {
	total := 0

	// find all X's then go through the iterations
	for _, pos := range A_positions {
		r := pos[0]
		c := pos[1]

		total += count_mas(r, c)
	}

	fmt.Printf(col.Green+"TOTAL: %d\n"+col.Reset, total)
}

func main() {
	input_file, _ = os.ReadFile("./input.txt")
	fmt.Printf("%s\n\n", input_file)

	lines := strings.Split(string(input_file), "\n")
	height = len(lines)
	width = len(lines[0])

	input_data = make([][]string, height)
	for r := range height {
		input_data[r] = make([]string, width)
		for c := range width {
			input_data[r][c] = string(lines[r][c])

			if input_data[r][c] == "X" {
				X_positions = append(X_positions, [2]int{r, c})
			}

			if input_data[r][c] == "A" {
				A_positions = append(A_positions, [2]int{r, c})
			}

		}
	}

	fmt.Printf("size: %dx%d\n", width, height)

	//p1()
	p2()
}

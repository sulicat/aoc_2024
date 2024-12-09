package main

import (
	"fmt"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

type Pos [2]int

var rows int
var cols int
var antennas map[string][]Pos
var lines []string
var found_spots map[Pos]int

func p1() {

	found_spots := make(map[Pos]int)
	count := 0
	for key, v := range antennas {
		fmt.Printf("key: %v -> %v\n", key, v)

		mychar := key

		for i1 := range v {
			for i2 := 0; i2 < len(v); i2 += 1 {
				//fmt.Printf("comp: %v %v\n", v[i1], v[i2])

				new_spot := Pos{
					v[i2][0] - v[i1][0],
					v[i2][1] - v[i1][1]}

				new_spot = Pos{
					v[i1][0] - new_spot[0],
					v[i1][1] - new_spot[1]}

				if new_spot[0] >= 0 && new_spot[0] < rows && new_spot[1] >= 0 && new_spot[1] < cols {

					if string(lines[new_spot[0]][new_spot[1]]) != mychar {
						fmt.Printf("New sPot: %v\n", new_spot)
						_, found_already := found_spots[new_spot]
						if !found_already {
							found_spots[new_spot] = 1
							count += 1
						}
					}

				}
			}
		}
	}

	fmt.Printf("p1 count: %d\n", count)
}

func p2() {

	found_spots = make(map[Pos]int)
	count := 0
	for key, v := range antennas {
		//fmt.Printf("key: %v -> %v\n", key, v)

		mychar := key

		for i1 := range v {
			for i2 := 0; i2 < len(v); i2 += 1 {
				if i1 == i2 {
					continue
				}

				diff := Pos{
					v[i2][0] - v[i1][0],
					v[i2][1] - v[i1][1]}

				diff_i := 0
				for {
					new_spot := Pos{
						v[i1][0] - diff[0]*diff_i,
						v[i1][1] - diff[1]*diff_i}
					diff_i += 1

					if new_spot[0] >= 0 && new_spot[0] < rows && new_spot[1] >= 0 && new_spot[1] < cols {

						if true || string(lines[new_spot[0]][new_spot[1]]) != mychar {
							fmt.Printf("New sPot: %v\n", new_spot)
							_, found_already := found_spots[new_spot]
							if !found_already {
								found_spots[new_spot] = 1
								count += 1
							}
						}
					} else {
						break
					}
				}

			}
		}
	}

	fmt.Printf("p2 count: %d\n", count)
}

func main() {

	antennas = make(map[string][]Pos)

	fmt.Printf(col.BgBlue + "Day 8" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines = strings.Split(string(file_data), "\n")

	rows = len(lines)
	cols = len(lines[0])

	fmt.Printf("width %d height %d\n", cols, rows)

	for r := range lines {
		for c, char := range lines[r] {
			if lines[r][c] != '.' {
				name := string(char)
				antennas[name] = append(antennas[name], Pos{r, c})
			}
		}
	}

	fmt.Println(antennas)

	p2()

	for r := range lines {
		for c, _ := range lines[r] {
			_, f := found_spots[Pos{r, c}]
			if f {
				fmt.Printf("#")

			} else {
				fmt.Printf("%s", string(lines[r][c]))
			}
		}
		fmt.Printf("\n")
	}

}

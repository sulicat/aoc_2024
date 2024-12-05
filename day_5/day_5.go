package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	utils "github.com/sulicat/goboi/utils"
)

type Rules map[int][]int // page depends on a list of pages

func care_about_rules(rules []int, care_about map[int]int) []int {
	out := make([]int, 0)
	for _, r := range rules {
		_, do_i_care := care_about[r]
		if do_i_care {
			out = append(out, r)
		}
	}
	return out
}

func insert(slice []int, index, value int) []int {
	// Create a new slice with one extra element
	newSlice := make([]int, len(slice)+1)

	// Copy elements before the index
	copy(newSlice[:index], slice[:index])

	// Insert the new value
	newSlice[index] = value

	// Copy the remaining elements
	copy(newSlice[index+1:], slice[index:])

	return newSlice
}

func RemoveIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		// Handle invalid index
		fmt.Println("Index out of range")
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func p1(rules Rules, updates [][]int) {
	total := 0

	fmt.Printf("Rules: %d\n", rules)

	for _, update := range updates {
		fmt.Printf("running: %d\n", update)

		// build a set of the updates to know which rules we can ignore
		seen := make(map[int]int)
		care_about := make(map[int]int)
		for _, u := range update {
			care_about[u] = 1
		}

		good := true
		for i := range update {
			u := update[i]
			seen[u] = 1 // mark that we have seen the number

			dependencies := care_about_rules(rules[u], care_about)
			//fmt.Printf("%d\n", u)
			//fmt.Printf("depen: %d\n", dependencies)
			//fmt.Printf("seen: %d\n", seen)

			for _, d := range dependencies {
				_, ok := seen[d]
				if !ok {
					//fmt.Printf("notseen\n")
					good = false
				}
			}
		}

		if good {
			mid := len(update) / 2
			total += update[mid]

		}
	}

	fmt.Printf(col.Green+"TOTAL p1: %d\n"+col.Reset, total)

}

func is_good(rules Rules, update []int) bool {
	seen := make(map[int]int)
	care_about := make(map[int]int)
	for _, u := range update {
		care_about[u] = 1
	}

	good := true
	for i := range update {
		u := update[i]
		seen[u] = 1 // mark that we have seen the number

		dependencies := care_about_rules(rules[u], care_about)
		//fmt.Printf("%d\n", u)
		//fmt.Printf("depen: %d\n", dependencies)
		//fmt.Printf("seen: %d\n", seen)

		for _, d := range dependencies {
			_, ok := seen[d]
			if !ok {
				//fmt.Printf("notseen\n")
				good = false
			}
		}
	}

	return good
}

func p2(rules Rules, updates [][]int) {
	total := 0

	fmt.Printf("Rules: %d\n", rules)

	for _, update := range updates {
		//fmt.Printf("running: %d\n", update)

		care_about := make(map[int]int)
		for _, u := range update {
			care_about[u] = 1
		}

		good := is_good(rules, update)

		if !good {

			for {
				good = is_good(rules, update)
				if good {
					break
				}

				reset := false
				for pos := 0; pos < len(update); pos += 1 {
					v := update[pos]
					dependencies := care_about_rules(rules[v], care_about)

					for _, d := range dependencies {
						depen_pos := slices.Index(update, d)
						if depen_pos > pos {
							temp := update[depen_pos]
							update[depen_pos] = update[pos]
							update[pos] = temp
							reset = true
							break
						}
					}

					if reset {
						break
					}

				}

			}

			fmt.Printf("%d\n", update)
			mid := len(update) / 2
			total += update[mid]

		}

		//fmt.Printf("\n")
	}

	fmt.Printf(col.Green+"TOTAL p1: %d\n"+col.Reset, total)

}

func main() {
	fmt.Printf("Day 5\n")
	var input_file []byte
	input_file, _ = os.ReadFile("./input.txt")
	//fmt.Printf("%s\n\n", input_file)
	lines := strings.Split(string(input_file), "\n")

	break_i := 0

	for i := range lines {
		if lines[i] == "" {
			break_i = i
		}
	}

	str_rules := lines[:break_i]
	str_updates := lines[break_i+1:]
	updates := utils.MapArray(str_updates,
		func(in string) []int {
			out := utils.MapArray(strings.Split(in, ","),
				func(num string) int {
					return utils.First(strconv.Atoi(num))
				})

			return out
		})

	rules := make(Rules)

	for _, r := range str_rules {
		s := strings.Split(r, "|")
		one := utils.First(strconv.Atoi(s[0]))
		two := utils.First(strconv.Atoi(s[1]))

		_, e := rules[two]

		if !e {
			rules[two] = make([]int, 0)
		}

		rules[two] = append(rules[two], one)
	}

	//p1(rules, updates)
	p2(rules, updates)
}

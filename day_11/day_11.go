package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

var data []int

// 1 -> 0
// 0 -> 1
// 1 -> 2048
// 2048 -> 20 48
// 20 48 -> 2 0 4 8
// 2 0 4 8 ->

type DD struct {
	Val    int
	Blinks int
}

var cache map[DD]int

func blink(in []int) []int {
	out := make([]int, 0)
	for _, v := range in {

		nv := int(0)

		if v == 0 {
			nv = 1
			out = append(out, nv)
		} else {

			digits := strconv.Itoa(int(v))
			num_digits := len(digits)

			if num_digits%2 == 0 {
				a1 := utils.First(strconv.Atoi(digits[:num_digits/2]))
				a2 := utils.First(strconv.Atoi(digits[num_digits/2:]))
				out = append(out, int(a1))
				out = append(out, int(a2))

			} else {
				out = append(out, v*2024)
			}
		}

	}
	return out
}

func count(v int, blinks_left int) int {

	cached, in_cache := cache[DD{v, blinks_left}]
	if in_cache {
		//fmt.Printf("chahe use: %d->%d\n", v, cached)
		return cached
	}

	if blinks_left == 0 {
		return 1
	}

	if v == 0 {
		retval := count(1, blinks_left-1)
		cache[DD{v, blinks_left}] = retval
		return retval

	} else {
		digits := strconv.Itoa(int(v))
		num_digits := len(digits)

		if num_digits%2 == 0 {
			a1 := utils.First(strconv.Atoi(digits[:num_digits/2]))
			a2 := utils.First(strconv.Atoi(digits[num_digits/2:]))

			retval := count(a1, blinks_left-1) + count(a2, blinks_left-1)
			cache[DD{v, blinks_left}] = retval
			return retval

		} else {

			retval := count(v*2024, blinks_left-1)
			cache[DD{v, blinks_left}] = retval
			return retval
		}
	}
}

func p1() {
	num_blinks := 0

	for {
		fmt.Printf("blink: %d ---> %v\n", num_blinks, data)
		if num_blinks >= 10 {
			fmt.Printf("Num stones: %d\n", len(data))
			break
		}

		new_data := blink(data)
		data = new_data
		num_blinks += 1

	}
}

func p2() {
	num_blinks := 75

	total := 0

	for _, d := range data {
		total += count(int(d), num_blinks)
	}

	fmt.Printf("total p2: %d\n", total)
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 10" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), " ")

	data = make([]int, 0)

	cache = make(map[DD]int)

	fmt.Printf("cache: %v\n", cache)

	for _, l := range lines {
		data = append(data, int(utils.First(strconv.Atoi(l))))
	}

	fmt.Printf("%v\n", data)

	//p1()
	//fmt.Printf("------\n")
	p2()
}

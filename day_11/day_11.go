package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

var data []int64

func blink(in []int64) []int64 {
	out := make([]int64, 0)
	for _, v := range in {

		nv := int64(0)

		if v == 0 {
			nv = 1
			out = append(out, nv)
		} else {

			digits := strconv.Itoa(int(v))
			num_digits := len(digits)

			if num_digits%2 == 0 {
				a1 := utils.First(strconv.Atoi(digits[:num_digits/2]))
				a2 := utils.First(strconv.Atoi(digits[num_digits/2:]))
				out = append(out, int64(a1))
				out = append(out, int64(a2))

			} else {
				out = append(out, v*2024)
			}
		}

	}
	return out
}

func p1() {
	num_blinks := 0

	for {
		//fmt.Printf("blink: %d ---> %v\n", num_blinks, data)
		if num_blinks >= 75 {
			fmt.Printf("Num stones: %d\n", len(data))
			break
		}

		fmt.Printf("%d %d\n", num_blinks, len(data))

		new_data := blink(data)
		data = new_data
		num_blinks += 1

	}
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 10" + col.Reset + "\n")

	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), " ")

	data = make([]int64, 0)

	for _, l := range lines {
		data = append(data, int64(utils.First(strconv.Atoi(l))))
	}

	fmt.Printf("%v\n", data)

	p1()

}

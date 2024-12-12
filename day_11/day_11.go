package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

var data []int64

// 1 -> 0
// 0 -> 1
// 1 -> 2048
// 2048 -> 20 48
// 20 48 -> 2 0 4 8
// 2 0 4 8 ->

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

func count(v int, blinks_left int) int {

	if blinks_left == 0 {
		return 1
	}

	if v == 0 {
		return count(1, blinks_left-1)
	} else {
		digits := strconv.Itoa(int(v))
		num_digits := len(digits)

		if num_digits%2 == 0 {
			a1 := utils.First(strconv.Atoi(digits[:num_digits/2]))
			a2 := utils.First(strconv.Atoi(digits[num_digits/2:]))
			return count(a1, blinks_left-1) + count(a2, blinks_left-1)
		} else {
			return count(v*2024, blinks_left-1)
		}
	}
}

func p1() {
	num_blinks := 0

	for {
		//fmt.Printf("blink: %d ---> %v\n", num_blinks, data)
		if num_blinks >= 10 {
			fmt.Printf("Num stones: %d\n", len(data))
			break
		}

		fmt.Printf("%d %d -> %v\n", num_blinks, len(data), data)

		new_data := blink(data)
		data = new_data
		num_blinks += 1

	}
}

func p2() {
	num_blinks := 75

	c := make(chan int)
	var wg sync.WaitGroup

	for _, d := range data {

		fmt.Printf("d: %v\n", d)
		wg.Add(1)

		go func(din int, sum chan int, wgin *sync.WaitGroup) {
			defer wgin.Done()
			fmt.Printf("go routine: %d\n", d)
			c <- count(din, num_blinks)
		}(int(d), c, &wg)

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	total := 0
	for t := range c {
		total += t
	}

	fmt.Printf("total p2: %d\n", total)
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

	//p1()
	//fmt.Printf("------\n")
	p2()
}

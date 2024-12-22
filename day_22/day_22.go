package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

var nums []int64
var sequences [][]int

func mix(secret int64, val int64) int64 {
	return secret ^ val
}

func prune(secret int64) int64 {
	return secret % 16777216
}

func p1() {

	sum := int64(0)

	for _, n := range nums {
		secret := n

		fmt.Printf("%v: ", secret)

		for i := 0; i < 2000; i += 1 {
			secret = prune(mix(secret, secret*64))
			secret = prune(mix(secret, int64(secret/32)))
			secret = prune(mix(secret, secret*2048))
		}

		fmt.Printf(" %v\n", secret)
		sum += secret
	}

	fmt.Printf("SUM %v\n", sum)
}

func last_digit(in int64) int {
	str := strconv.Itoa(int(in))
	out, _ := strconv.Atoi(string(str[len(str)-1]))
	return out
}

func p2() {
	for _, n := range nums {
		secret := n
		new_seq := []int{}
		new_seq = append(new_seq, last_digit(secret))

		for i := 0; i < 2000; i += 1 {
			secret = prune(mix(secret, secret*64))
			secret = prune(mix(secret, int64(secret/32)))
			secret = prune(mix(secret, secret*2048))
			new_seq = append(new_seq, last_digit(secret))
		}

		sequences = append(sequences, new_seq)
	}

	fmt.Printf("%v %v\n", len(sequences), len(sequences[0]))

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 22" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	nums = []int64{}
	sequences = [][]int{}

	for _, l := range lines {
		number, _ := strconv.Atoi(l)
		nums = append(nums, int64(number))
	}

	p2()
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

var nums []int64
var sequences [][]int
var diff_sequences [][]int

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

func match_seq(seq []int, pattern []int) int {
	indx := -1
	return indx
}

func p2() {
	for _, n := range nums {
		secret := n

		new_seq := []int{}
		new_seq = append(new_seq, last_digit(secret))

		new_diff_seq := []int{}
		new_diff_seq = append(new_diff_seq, 0)

		for i := 0; i < 2000; i += 1 {
			secret = prune(mix(secret, secret*64))
			secret = prune(mix(secret, int64(secret/32)))
			secret = prune(mix(secret, secret*2048))

			digit := last_digit(secret)
			new_diff_seq = append(new_diff_seq, digit-new_seq[len(new_seq)-1])
			new_seq = append(new_seq, digit)
		}

		sequences = append(sequences, new_seq)
		diff_sequences = append(diff_sequences, new_diff_seq)
	}

	max_score := 0
	for i, test_seq := range sequences {

		// find max index
		// check all other sequences and sum the total
		max_index := utils.MaxIndex(test_seq[4:]) + 5
		score := 0
		score += test_seq[max_index]

		diff_seq := diff_sequences[i]

		pattern := []int{
			diff_seq[max_index-3],
			diff_seq[max_index-2],
			diff_seq[max_index-1],
			diff_seq[max_index],
		}

		fmt.Printf("Pattern: %d - > %v\n", max_index, pattern)

		for other_i, other_seq := range sequences {
			if i == other_i {
				continue
			}

			matched_seq := match_seq(diff_sequences[i], pattern)
			if matched_seq >= 0 {
				score += other_seq[matched_seq]
			}

		}

		if score > max_score {
			max_score = score
		}

	}
	fmt.Printf("MAX: %v \n", max_score)

}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 22" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	nums = []int64{}
	sequences = [][]int{}
	diff_sequences = [][]int{}

	for _, l := range lines {
		number, _ := strconv.Atoi(l)
		nums = append(nums, int64(number))
	}

	p2()
}

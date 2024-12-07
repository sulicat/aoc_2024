package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

var answers []int64
var nums [][]int64

//nums := [][]int

func check(base int64, nums []int64) (bool, []string) {
	// exit condition
	if len(nums) == 1 {
		return nums[0] == int64(base), []string{"D"}
	}

	if base < 0 {
		return false, []string{"D"}
	}

	l := len(nums)
	last_item := nums[l-1]

	good_div, path_div := check(int64(base)/int64(last_item), nums[:l-1])
	good_sub, path_sub := check(int64(base)-int64(last_item), nums[:l-1])

	retpath := path_div
	if good_sub {
		retpath = path_sub
		retpath = append(retpath, "+")
	} else {
		retpath = append(retpath, "*")
	}

	return (good_div || good_sub), retpath
}

func p1() {
	var count uint64
	count = uint64(0)

	for i := range nums {

		fmt.Printf("CHECKING: %d\n", answers[i])
		good, path := check(int64(answers[i]), nums[i])
		fmt.Printf("%v %d %v\n %v \n", good, answers[i], nums[i], path)

		if good {
			count += uint64(answers[i])
		}
	}

	fmt.Printf("COUNT: %d\n", count)
}

func main() {

	fmt.Printf(col.BgBlue + "DAY 7" + col.Reset + "\n")

	file_data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file_data), "\n")

	for i, l := range lines {

		f := strings.Split(l, ": ")
		res, _ := strconv.ParseInt(f[0], 10, 64)
		answers = append(answers, res)

		n := strings.Split(f[1], " ")
		nums = append(nums, []int64{})

		for _, num := range n {
			res, _ := strconv.ParseInt(num, 10, 64)
			nums[i] = append(nums[i], res)
		}

		fmt.Printf("%d %d\n", res, nums[i])
	}
	fmt.Printf("\n")

	p1()
}

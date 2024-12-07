package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
)

var answers []int64
var nums [][]int64

//nums := [][]int

func check(target int64, nums []int64, index int, val int64) bool {

	if val > target {
		return false
	}

	if index >= len(nums) {
		return val == target
	}

	check_mult := false

	if index == 0 {
		check_mult = check(target, nums, index+1, nums[index])
	} else {
		check_mult = check(target, nums, index+1, val*nums[index])

	}
	check_add := check(target, nums, index+1, val+nums[index])

	return check_mult || check_add
}

func concat(v1, v2 int64) int64 {
	str := strconv.FormatInt(v1, 10) + strconv.FormatInt(v2, 10)
	return utils.First(strconv.ParseInt(str, 10, 64))
}

func check2(target int64, nums []int64, index int, val int64) bool {

	if val > target {
		return false
	}

	if index >= len(nums) {
		return val == target
	}

	check_mult := false
	if index == 0 {
		check_mult = check2(target, nums, index+1, nums[index])
	} else {
		check_mult = check2(target, nums, index+1, val*nums[index])
	}

	check_add := check2(target, nums, index+1, val+nums[index])

	check_concat := check2(target, nums, index+1, concat(val, nums[index]))

	return check_mult || check_add || check_concat
}

func p1() {
	var count int64
	count = 0

	for i := range nums {

		fmt.Printf("CHECKING: %d\n", answers[i])
		good := check(answers[i], nums[i], 0, 0)
		fmt.Printf("%v %d %v\n \n", good, answers[i], nums[i])

		if good {
			count += answers[i]
		}
	}

	fmt.Printf("COUNT: %d\n", count)
}

func p2() {
	var count int64
	count = 0

	for i := range nums {

		fmt.Printf("CHECKING: %d\n", answers[i])
		good := check2(answers[i], nums[i], 0, 0)
		fmt.Printf("%v %d %v\n \n", good, answers[i], nums[i])

		if good {
			count += answers[i]
		}
	}

	fmt.Printf("COUNT2: %d\n", count)
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

	a := 1234
	b := 5678
	fmt.Printf("aaa %d\n", concat(int64(a), int64(b)))

	//p1()
	p2()
}

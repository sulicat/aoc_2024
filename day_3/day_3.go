package main

import (
	"fmt"
	"strconv"

	col "github.com/sulicat/goboi/colors"
	utils "github.com/sulicat/goboi/utils"
)

func main() {
	fmt.Printf("%s\n", col.BrightBlue+"Day 3"+col.Reset)

	test_arr := make([]string, 0)
	test_arr = append(test_arr, "1", "2", "3", "4")

	int_arr := utils.MapArray(test_arr, func(a string) int {
		return utils.First(strconv.Atoi(a))
	})

	utils.ForEach(test_arr, func(a string) {
		fmt.Printf("asdasd %s\n", a)
	})

	for _, val := range int_arr {
		fmt.Printf("%d\n", val)
	}

}

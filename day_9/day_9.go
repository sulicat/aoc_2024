package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/sulicat/goboi/utils"
)

var data_string string
var data []int64

func p1() {
	//fmt.Printf("data: %v\n", data)

	p1 := 0
	p2 := len(data) - 1

	for {

		// find the empty space
		if data[p1] == -1 {

			// find the first non empty space
			for {
				if data[p2] != -1 {
					break
				} else {
					p2 -= 1
				}
			}

			if p1 > p2 {
				break
			}

			//temp := data[p1]
			data[p1] = data[p2]
			data[p2] = -1

		} else {
			p1 += 1
		}

		if p1 > len(data)-1 {
			break
		}

	}

	sum := int64(0)
	for i, d := range data {
		if d == -1 {
			break
		}

		sum += int64(i) * d
	}

	//fmt.Printf("data: %v\n", data)
	//fmt.Printf("SUM: %d\n", sum)
}

type Segment struct {
	Index int
	Size  int
}

func p2() {
	//fmt.Printf("data: %v\n", data)

	current := data[:]

	p1 := 0

	// make an map of segments
	empty_segments := make([]Segment, 0)

	for {
		// find the first empty segment
		first := slices.Index(current, -1)
		if first == -1 {
			break
		}

		segment_size := slices.IndexFunc(current[first:],
			func(a int64) bool {
				return a != -1
			})

		if segment_size == -1 {
			break
		}

		p1 += first + segment_size + 1
		current = data[p1:]
		segment_index := p1 - segment_size - 1
		empty_segments = append(empty_segments, Segment{Index: segment_index, Size: segment_size})

	}

	// fmt.Printf("data start %v\n", data)
	// fmt.Printf("empty segments %v\n", empty_segments)

	p2 := len(data) - 1
	end := p2
	move_size := 9999

	for {
		if end <= 0 {
			break
		}

		// find first non -1
		for {
			if data[end] != -1 {
				break
			} else {
				end -= 1

				if end < 1 {
					break
				}
			}
		}

		// find the start
		start := end - 1
		for {
			if start < 0 {
				break
			}

			if data[start] != data[end] {
				break
			} else {
				start -= 1
				if end < 0 {
					break
				}
			}
		}

		move_size = end - start

		// fmt.Printf("start %d end %d data %d len: %d\n", start, end, data[end], move_size)

		if end < 0 || start < 0 || end < start {
			break
		}

		moveme := data[start+1 : start+move_size+1]

		// look for a segment to fill
		for si, s := range empty_segments {

			if s.Size >= move_size && start > s.Index {

				// fmt.Printf("moveme: %v\n", moveme)
				copy(data[s.Index:], moveme)
				for x := range moveme {
					moveme[x] = -1
				}

				empty_segments[si].Size -= move_size
				empty_segments[si].Index += move_size
				// fmt.Printf("moved: %d %d N:%d   -> %d %v\n", start, end, data[end], si, s)
				// fmt.Println(data)
				break
			}

		}

		end = start

		// fmt.Printf("\n")
	}

	//fmt.Printf("data: %v\n", data)

	sum := int64(0)
	for i, d := range data {
		if d == -1 {
			continue
		}

		sum += int64(i) * d
	}
	fmt.Printf("SUM2: %d\n", sum)

	//fmt.Printf("SUM: %d\n", sum)
}

func main() {
	file_data, _ := os.ReadFile("input.txt")
	//fmt.Printf("%s\n", file_data)
	//current_char := ""

	file_id := 0

	for i, c := range file_data {

		if i%2 == 0 {
			//current_char = string(c)
			num_data, _ := strconv.Atoi(string(c))
			for range num_data {
				//data_string += string(current_char)
				data_string += utils.First(strconv.Itoa(file_id))
				data = append(data, int64(file_id))
			}

			file_id += 1

		} else {
			num_empty, _ := strconv.Atoi(string(c))
			for range num_empty {
				data_string += "."
				data = append(data, -1)
			}
		}
	}

	p2()

	//p1()

	//fmt.Printf("%s\n", data_string)
}

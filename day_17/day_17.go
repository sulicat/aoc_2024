package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

var reg_a int
var reg_b int
var reg_c int

type Instruction struct {
	instr int
}

var program []Instruction
var instruction_ptr int

func combo(oper int) int {
	if oper >= 0 && oper <= 3 {
		return oper
	}

	if oper == 4 {
		return reg_a
	}

	if oper == 5 {
		return reg_b
	}

	if oper == 6 {
		return reg_c
	}

	return 0

}

func process(opcode int, operand int) {
	switch opcode {

	case 0: //adv
		num := float64(reg_a)
		den := math.Pow(2.0, float64(combo(operand)))
		reg_a = int(num / den)

	case 1: //bxl
		reg_b = reg_b ^ operand

	case 2: // bst
		reg_b = combo(operand) % 8

	case 3: // jnz
		if reg_a != 0 {
			instruction_ptr = (operand) - 2
		}

	case 4: // bxc
		reg_b = reg_b ^ reg_c

	case 5: //out
		fmt.Printf("%d,", combo(operand)%8)

	case 6: //bdv
		num := float64(reg_a)
		den := math.Pow(2.0, float64(combo(operand)))
		reg_b = int(num / den)

	case 7: //cdv
		num := float64(reg_a)
		den := math.Pow(2.0, float64(combo(operand)))
		reg_c = int(num / den)
	}
}

func p1() {
	instruction_ptr = 0

	for {
		if int(instruction_ptr) > len(program)-1 {
			fmt.Printf("\nDONE\n")
			break
		}

		process(
			program[instruction_ptr].instr,
			program[instruction_ptr+1].instr,
		)

		instruction_ptr += 2
	}
}

func main() {
	fmt.Printf(col.BgBrightCyan + "Day 16" + col.Reset + "\n")
	file_data, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file_data), "\n")

	reg_a_str := strings.Split(lines[0], "Register A: ")[1]
	reg_a_int, _ := strconv.Atoi(reg_a_str)
	reg_a = int(reg_a_int)

	reg_b_str := strings.Split(lines[1], "Register B: ")[1]
	reg_b_int, _ := strconv.Atoi(reg_b_str)
	reg_b = int(reg_b_int)

	reg_c_str := strings.Split(lines[2], "Register C: ")[1]
	reg_c_int, _ := strconv.Atoi(reg_c_str)
	reg_c = int(reg_c_int)

	program_str := strings.Split(lines[4], "Program: ")[1]
	program_codes := strings.Split(program_str, ",")
	program = []Instruction{}

	for _, pc := range program_codes {
		code, _ := strconv.Atoi(pc)
		program = append(program, Instruction{instr: int(code)})
	}

	fmt.Printf("a:%v b:%v c:%v \n", reg_a, reg_b, reg_c)

	fmt.Printf("\nprogram:\n  %v\n\n\n", program)

	p1()
}

package main

type Stack []Pos

func (s *Stack) Push(x Pos) {
	*s = append(*s, x)
}

func (s *Stack) Empty() bool {
	return len(*s) <= 0
}

func (s *Stack) Pop() Pos {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

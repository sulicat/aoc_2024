package main

type PosHeap []Pos

func (h PosHeap) Len() int {
	return len(h)
}

func (h PosHeap) Less(i int, j int) bool {
	return h[i].cost < h[j].cost
}

func (h PosHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h PosHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h = append(h, x.(Pos))
}

func (h PosHeap) Pop() interface{} {
	old := h
	n := len(old)
	x := old[n-1]
	h = old[0 : n-1]
	return x
}

func (h PosHeap) Peek() interface{} {
	old := h
	n := len(old)
	x := old[n-1]
	return x
}

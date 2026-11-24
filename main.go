package main

import (
	"container/heap"
	"fmt"
)

type cellItem struct{ t, r, c int }

type cellHeap []cellItem

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x any)        { *h = append(*h, x.(cellItem)) }
func (h *cellHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// SwimInWater returns the earliest time the bottom-right is reachable from
// the top-left when water at time t covers all cells of elevation ≤ t —
// i.e. the minimal possible "highest cell on the path". Dijkstra with the
// path cost being a MAX of elevations instead of a sum.
func SwimInWater(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &cellHeap{{grid[0][0], 0, 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cellItem)
		if visited[cur.r][cur.c] {
			continue
		}
		visited[cur.r][cur.c] = true
		if cur.r == n-1 && cur.c == n-1 {
			return cur.t
		}
		for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			r, c := cur.r+d[0], cur.c+d[1]
			if r < 0 || r >= n || c < 0 || c >= n || visited[r][c] {
				continue
			}
			t := cur.t
			if grid[r][c] > t {
				t = grid[r][c] // the path's cost is its highest cell
			}
			heap.Push(h, cellItem{t, r, c})
		}
	}
	return -1 // unreachable only for malformed input
}

func main() {
	fmt.Println(SwimInWater([][]int{{0, 2}, {1, 3}}))
	fmt.Println(SwimInWater([][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	}))
}

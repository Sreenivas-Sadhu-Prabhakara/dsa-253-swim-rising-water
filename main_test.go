package main

import "testing"

func TestSwimInWater(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{
			{0, 1, 2, 3, 4},
			{24, 23, 22, 21, 5},
			{12, 13, 14, 15, 16},
			{11, 17, 18, 19, 20},
			{10, 9, 8, 7, 6},
		}, 16},
		{[][]int{{0}}, 0},
		{[][]int{{3, 2}, {0, 1}}, 3}, // start cell's own height counts
	}
	for _, c := range cases {
		if got := SwimInWater(c.grid); got != c.want {
			t.Errorf("SwimInWater = %d, want %d", got, c.want)
		}
	}
}

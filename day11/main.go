package main

import (
	"fmt"
	"math"
)

const Serial = 9445
const Size = 301

func makeGrid() [Size][Size]int {
	var grid [Size][Size]int
	for y := 1; y < Size; y++ {
		for x := 1; x < Size; x++ {
			id := x + 10
			power := id*y + Serial
			power = (power*id)/100%10 - 5
			grid[y][x] = power
		}
	}
	return grid
}

func makeSumTable(grid [Size][Size]int) [Size][Size]int {
	var table [Size][Size]int
	for y := 1; y < Size; y++ {
		for x := 1; x < Size; x++ {
			table[y][x] = grid[y][x] + table[y-1][x] + table[y][x-1] - table[y-1][x-1]
		}
	}
	return table
}

// dimension describes the box width and height. So 1 = 1x1, 4 = 4x4
func best(dimension int, table [Size][Size]int) (int, int, int) {
	bx, by := 0, 0
	max := math.MinInt64
	for y := dimension; y < Size; y++ {
		for x := dimension; x < Size; x++ {
			yt, xt := y-dimension, x-dimension
			t := table[y][x] - table[yt][x] - table[y][xt] + table[yt][xt]
			if t > max {
				bx, by, max = x-dimension+1, y-dimension+1, t
			}
		}
	}
	return bx, by, max
}

func main() {
	grid := makeGrid()
	table := makeSumTable(grid)
	bx, by, max := best(3, table)
	fmt.Println(fmt.Sprintf("Start: %v,%v Max: %v", bx, by, max))
	max = math.MinInt64
	bw := 0
	for i := 1; i < Size; i++ {
		x, y, m := best(i, table)
		if m > max {
			max = m
			bx, by = x, y
			bw = i
		}
	}
	fmt.Println(fmt.Sprintf("Start: %v,%v,%v Max: %v", bx, by, bw, max))
}

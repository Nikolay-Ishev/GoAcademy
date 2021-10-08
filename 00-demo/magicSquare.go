package main

import (
	"fmt"
	"math"
)


func formingMagicSquare(s [][]int32) int32 {
	pre := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
    //    (00, 01, 02) (10 11 12) (20 21 22)
	var diffSl [8]float64
	for i:= 0; i<8; i++ {
		diff:=0.
		for row:= 0; row<3; row++ {
			for col := 0; col <3; col++{
				diff += math.Abs(float64(pre[i][row][col] - s[row][col]))
			}

		}
		diffSl[i] =  diff
	}
	min := diffSl[0]
	for _, v := range diffSl {
		if (v < min) {
			min = v
		}
	}
	return int32(min)
}

func main() {
	var a, b, c, d, e, f, g, h, i   int32
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i)
	s :=[][] int32 {{a, b, c}, {d, e, f}, {g, h, i}}
	fmt.Println(formingMagicSquare(s))
}





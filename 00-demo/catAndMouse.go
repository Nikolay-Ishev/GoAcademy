package main

import (
	"fmt"
	"math"
)

func catAndMouse(x float64, y float64, z float64) string {

	catA := int(math.Abs(float64(z - x)))
	catB := int(math.Abs(float64(z - y)))

	switch {
	case catA < catB:
		return "Cat A"
	case catB < catA:
		return "Cat B"
	default:
		return "Mouse C"
	}
}
func main() {
	var q int
	fmt.Scan(&q)
	for ; q > 0; q-- {
		var catA, catB, mouse float64
		_, err := fmt.Scan(&catA, &catB, &mouse)
		if err != nil {
			return
		}
		fmt.Println(catAndMouse(catA, catB, mouse))
	}
}


package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var asci_map = make(map[int]int32)
	for i:=97; i<123; i++ {
		asci_map[i] = h[i-97]
	}
	heights := make([]int32, len(word))
	for i:=0; i<len(word); i++{
		ascii := int(word[i])
		heights = append(heights, asci_map[ascii])
	}
	max := heights[0]
	for _, value := range heights {
		if max < value {
			max = value
		}
	}
	return max * int32(len(word))
}

func main() {
	b := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	fmt.Println(designerPdfViewer(b,"abc"))
}
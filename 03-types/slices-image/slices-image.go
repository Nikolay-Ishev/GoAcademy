package main

import (
	"fmt"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/mypic"
	"log"
	"os"
	"path"
)

const baseDir = "C:/Users/Nikolay/Desktop/Softuni - Python/Go/Git/coursegowork/03-types/slices-image"

// Pic returns a grayscale pic of size dy * dx
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range(result) {
		result[i] = make([]uint8, dx)
		for ci := range(result[i]) {
			result[i][ci] = uint8((i+ci)/2)
		}
	}
	return result
}

func findUniqueValue(arr []int32) map[int32]int32{
	//Create a   dictionary of values for each element
	dict:= make(map[int32]int32)
	for _ , num :=  range arr {
		dict[num] = dict[num]+1
	}
	return dict
}

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// programPath := "os.Args[0]"
	// fmt.Println(programPath)
	// dir := path.Dir(programPath)
	// fmt.Println(dir)

	imageFile := path.Join(baseDir, "image.png")
	fmt.Println(imageFile)
	file, err := os.Create(imageFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	mypic.Encode(Pic, file)
}

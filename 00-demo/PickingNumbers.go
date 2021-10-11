package main

import "fmt"

func findUniqueValue(arr []int32) map[int32]int32{
	//Create a   dictionary of values for each element
	dict:= make(map[int32]int32)
	for _ , num :=  range arr {
		dict[num] = dict[num]+1
	}
	return dict
}

func pickingNumbers(a []int32) int32 {
	dictNumbers := findUniqueValue(a)
	var longestValues int32 = 0
	for i:=0;i<len(a);i++{
		currentNum := a[i]
		uniqueValues := dictNumbers[currentNum] + dictNumbers[currentNum - 1]
		if uniqueValues > longestValues {
			longestValues = uniqueValues
		}
	}
	return longestValues
}

func main() {
	var customArr = []int32 {1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	fmt.Println(pickingNumbers(customArr))
}
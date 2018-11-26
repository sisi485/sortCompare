package main

import (
	"fmt"
	"math/rand"
	"sortCompare/Sort"
)

func PrintArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]:%d\n", i, arr[i])
	}
}


func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func CreateRandomArray(n int) []int {
	b := make([]int, n)

	for i := 0; i < n; i++  {
		b[i] = 	randomInt(0, SIZE)
		//fmt.Printf("arr[%d]:%d\n", i, b[i])
	}
	return b
}


//compares, if the array checkMe is a sorted array from compare
func CheckArray(checkMe, compare []int) bool {

	compareSorted := make([]int, len(compare))
	copy(compareSorted, compare)
	Sort.Quicksort(compareSorted)

	for i := 0; i < len(compareSorted); i++ {
		if compareSorted[i] != checkMe[i] {
			return false
		}
	}
	return true
}
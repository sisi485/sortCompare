package main

import (
	"fmt"
	"math/rand"
	"sortCompare/Sort"
	"time"
)

const (
	Selection = 0
	Insertion = 1
	Merge     = 2
	Quick     = 3
	Heap      = 4
)

const (
	SIZE       = 15
	REPETITION = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("start")
	run()
	fmt.Println("ende")
}

func run() {

	data := CreateRandomArray(SIZE)

	fmt.Println("Running selection sort:\n Run\t\t Time")
	initSort(data, Selection)

	fmt.Println("Running insertion sort:\n Run\t\t Time")
	initSort(data, Insertion)

	fmt.Println("Running merge sort:\n Run\t\t Time")
	initSort(data, Merge)

	fmt.Println("Running quick sort:\n Run\t\t Time")
	initSort(data, Quick)

	fmt.Println("Running heap sort:\n Run\t\t Time")
	initSort(data, Heap)

}

func initSort(data []int, sort int) {

	for i := 0; i < REPETITION; i++ {

		//create copy
		sortMe := make([]int, SIZE)
		copy(sortMe, data)
		timeStart := time.Now()

		switch sort {

		case Selection:
			Sort.SelectionSort(sortMe)
			break

		case Insertion:
			Sort.Insertionsort(sortMe)
			break

		case Merge:
			Sort.MergeSort(sortMe, SIZE)
			break

		case Quick:
			Sort.Quicksort(sortMe)
			break

		case Heap:
			Sort.HeapSort(sortMe)
			break

		default:
			fmt.Println("could not get sorting algorithm")
		}

		elapsed := time.Since(timeStart)
		fmt.Printf(" %d\t\t\t %s\n", i+1, elapsed)
		fmt.Println()

		if CheckArray(sortMe, data) {
			fmt.Println("is sorted")
		} else {
			fmt.Println("is not sorted\n Source:")
			PrintArr(data)
			fmt.Println("\nDestination:")
			PrintArr(sortMe)
		}

	}
	fmt.Println()
}

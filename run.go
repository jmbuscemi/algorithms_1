package main

import (
	"coursera_algorithms_1/week_1"
	"fmt"
)

//return if array has 0 or 1 elements
//if array has odd lengths, then break it into two pieces (n-1)/2 and (n+1)/2

func main() {
	slice := []int{4, 75, 3, 765, 9, 3, 2, 2, 3, 23, 987, 74, 100, 8747, 5, 6, 2, 0}

	fmt.Printf("%v\n%v\n", slice, mergesort.MergeSort(slice))
}

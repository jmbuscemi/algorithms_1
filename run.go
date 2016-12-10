package main

import (
	"coursera_algorithms_1/week_1"
	"fmt"
	"math/rand"
	"time"
)

//return if array has 0 or 1 elements
//if array has odd lengths, then break it into two pieces (n-1)/2 and (n+1)/2

func main() {
	slice := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 21; i++ {
		slice = append(slice, rand.Int())
	}

	fmt.Printf("%v\n%v\n", slice, mergesort.MergeSort(slice))
}

package main

import "fmt"

func main() {
	// var arr1 [5]int
	arr1 := [5]int{}
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

}

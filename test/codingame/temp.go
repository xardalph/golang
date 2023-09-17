package main

import "fmt"

func main() {
	stack := make([][]rune, 1)

	stack = append(stack, make([]rune, 99))

	// declaring a slice of slices of
	// type integer with a length of 3
	slice_of_slices := make([][]int, 3)

	for i := 0; i < 3; i++ {

		new_length := i*2 + 1
		// looping through the slice to declare
		// slice of slice of a variable length
		slice_of_slices[i] = make([]int, new_length)

		// assigning values to each
		// slice of a slice
		for j := 0; j < new_length; j++ {
			slice_of_slices[i][j] = i*j + 1
		}
	}

	// printing the slice of slices matrix
	for i := 0; i < 3; i++ {
		fmt.Println(slice_of_slices[i])
	}
}

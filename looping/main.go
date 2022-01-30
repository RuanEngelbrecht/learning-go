package main

import "fmt"

func main() {
	// Only loop in Go is the for loop.

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// Do while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	// Continue statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// Nesting loops using a label to break out of the outer loop
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// You can also use a for loop with a range statement.
	// The range statement is used to iterate over a slice, map, or string.
	// The range statement returns a value of type `range` that contains
	// two values, the key and the value of the element.

	s := []int{1, 2, 3, 4, 5}

	for k, v := range s {
		fmt.Println(k, v)
	}

	for _, v := range "Hello" {
		fmt.Println(v, string(v))
	}

}

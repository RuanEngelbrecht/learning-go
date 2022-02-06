package main

import "fmt"

func main() {
	// 1. Arrays
	// Arrays have a fixed size that has to be known at compile time.
	// Best use case is to use an array to back a slice.
	// Syntax: [size]type{values}

	// Fixed size array.
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)

	// Instead of double declaration (size and values), you can use the array literal syntax.
	grades := [...]int{97, 85, 92}
	fmt.Printf("Grades: %v\n", grades)

	// Uninitialized array.
	var students [3]string
	students[0] = "John"
	students[1] = "Paul"
	students[2] = "George"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a // arrays are always copied in Go. Use pointers (prefix & -> address) if you want to share the same array.
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// 2. Slices
	// Slice is a projection of the underlying array.
	// Slices are reference types.
	// Syntax: []type{values}

	c := []int{1, 2, 3} // slice literal
	fmt.Println(c)
	fmt.Println("Length:", len(c))
	fmt.Println("Capacity:", cap(c)) // capacity is the number of elements that can be stored in the underlying array.

	d := c // d and c are pointing to the same underlying array!
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	e := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array literal that we want to slice.

	// e[0:5] -> [1, 2, 3, 4, 5]
	// first number (0) is inclusive, second (5) is exclusive.

	f := e[:] // slice of all elements.
	fmt.Println(f)
	g := e[3:] // slice of all elements after index 3.
	fmt.Println(g)
	h := e[:6] // slice of all elements before index 6.
	fmt.Println(h)
	i := e[3:6] // slice of elements from index 3 to index 6.
	fmt.Println(i)

	// Creating slices via make function.
	// make([]type, size, capacity?)

	j := make([]int, 3) // initialized with zero values.
	fmt.Println(j)
	fmt.Printf("Length: %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	// Slice append function.

	k := []int{}
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	k = append(k, 1, 2, 3) //variatric function i.e. you can add multiple values of the same type at once.
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))

	// Append a slice to another slice.

	l := []int{1, 2, 3}
	m := []int{4, 5, 6}
	n := append(l, m...) // ... -> spread operator.
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	// Stack operations on slices.

	o := []int{1, 2, 3}
	p := o[:len(o)-1] // pop the last element.
	fmt.Println(p)

	// "js slice" a slice. Beware of this behavior because the slice is not copied.
	q := []int{1, 2, 3, 4, 5}
	r := append(q[:2], q[3:]...)
	fmt.Println(r)
	fmt.Println(q) // q has been modified.
}

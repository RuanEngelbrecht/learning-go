package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a // b is a pointer to an int value (*int), that points to the memory address of a.
	fmt.Println(&a, b)
	fmt.Println(a, *b) // deferencing operator (*) -> read the value at the memory address of b.
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	// Pointer arithmetic -> intentionally left out of Go, but you can use the "unsafe" package if you need to.
	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1] // &c[1] - 8 will throw an error.
	fmt.Printf("%v %p %p\n", c, d, e)

	// Struct initialization syntax wih pointers.
	var ms *myStruct = &myStruct{foo: 42}
	fmt.Println(ms)

	// Using the new keyword (new allocates memory, but doesn't initialize it.)
	var ms2 *myStruct = new(myStruct)
	ms2.foo = 42 // compiler will automatically dereference the pointer.
	fmt.Println(ms2.foo)

	// Important to check for nil pointers
	var ms3 *myStruct // uninitialized pointer
	if ms3 == nil {
		fmt.Println("ms3 is nil")
	}

	// Slices and Maps contain internal pointers to the underlying data.
	// So copies point to the same underlying data.
}

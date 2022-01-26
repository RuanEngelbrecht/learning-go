package main

import (
	"fmt"
	"strconv"
)

// Package level scope

// Variable declaration -> no shorthand declarations here in package level scope.
// Variables with inner-most scope take precedence i.e. declaring var i in main().
var i int = 42

// Declare multiple variables at once in a var block
var (
	actorName     string = "Elisabeth Sladen"
	companionName string = "Sarah Jane Smith"
	doctorNumber  int    = 3
	season        int    = 11
)

// variable declaration with lowercase -> scoped to the package.
var inside string = "Hello from the inside"

// variable declaration with uppercase -> export from package.
var Outside string = "Hello from the outside"

func main() {
	// Package level variable i will be shadowed by the local variable i, i.e. 'hiding' it after the declaration.
	// Can't redeclare variables, but can shadow them. See below.
	fmt.Printf("%v %T\n", i, i)

	// Simple variable declaration.
	var i int = 43

	// Shorthand -> Go infers the type of the variable.
	j := 99.3

	// Variable declaration with type -> more control over variable types.
	var k float32 = 21.3

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)

	// Remember to use all the variables you declare.
	fmt.Println(actorName, companionName, doctorNumber, season)

	fmt.Printf("%s\n%s\n", inside, Outside)

	// Convert variable types.
	var a int = 42
	fmt.Printf("%v %T\n", a, a)
	var b float32 = float32(a)
	fmt.Printf("%v %T\n", b, b)

	// Careful with the other way around as you can lose data.
	var c float32 = 42.3
	fmt.Printf("%v %T\n", c, c)
	var d int = int(c)
	fmt.Printf("%v %T\n", d, d)

	// Converting int to string requires the use of strconv package.
	var e int = 42
	fmt.Printf("%v %T\n", e, e)
	var f string = string(e) // converts int bytes into "*". Use strconv.Itoa(e)
	fmt.Printf("%v %T\n", f, f)
	var g string = strconv.Itoa(e) // converts int to string. Itoa -> Int to ASCII.
	fmt.Printf("%v %T\n", g, g)
}

/*
	Naming conventions:
		- Pascal or camelCase
			- Capitalize acronyms i.e. HTTP, URL
		- As short as reasonable
			- Longer names for longer lives
*/

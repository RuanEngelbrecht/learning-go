package main

import "fmt"

// Enumerated constants

const (
	c = iota // iota is scoped to const block.
	d        // compiler will infer the type i.e. no need to do d = iota.
	e
)

const (
	// common to set first value to an error value.
	// underscore is Go's only write-only variable i.e. we don't need the zero value.
	// using _ won't assign it in memory.

	_ = iota //offset starting value by iota + 3.
	cat
	dog
	snake
)

// Bit shifting

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Using bit shifting to set flags in a boolean byte
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// Has to be assignable at compile time.
	// e.g. const a float64 = math.Sin(1.57) // compile error.
	// Immutable but can be shadowed.

	const a = 42

	fmt.Printf("%v %T\n", a, a)

	var b int16 = 27

	fmt.Printf("%v %T\n", a+b, a+b)

	// Enumerated constants

	fmt.Printf("%v %T\n", c, c) //0
	fmt.Printf("%v %T\n", d, d) //1
	fmt.Printf("%v %T\n", e, e) //2

	var specialist int = cat
	fmt.Printf("%v\n", specialist == cat)

	var otherSpecialist int // initialized with a zero value. Hence the need for the error value.
	fmt.Printf("%v\n", otherSpecialist == cat)

	filesize := 4000000000.
	fmt.Printf("%.2fGB\n", filesize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", roles&isAdmin == isAdmin) // bit mask
	fmt.Printf("Is HQ? %v\n", roles&isHeadquarters == isHeadquarters)
}

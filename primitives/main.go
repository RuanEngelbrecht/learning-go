package main

import (
	"fmt"
)

func main() {
	// Boolean values
	// Not an alias for other types (int)
	var n bool = false
	fmt.Printf("%v %T\n", n, n)

	// Boolean shorthand
	m := true
	fmt.Printf("%v %T\n", m, m)

	// Initialize boolean == 0 == false
	var p bool
	fmt.Printf("%v %T\n", p, p)

	// Unsigned integers
	var q uint16 = 42
	fmt.Printf("%v %T\n", q, q)

	// Byte = unsigned 8-bit integer
	var r byte = 'a'
	fmt.Printf("%v %T\n", r, r)

	//	Can't mix types of the same family. Convert.
	var a int = 1
	var b int8 = 2
	fmt.Println(a + int(b))

	// Bit operators
	var s uint8 = 10    //1010
	var t uint8 = 3     //0011
	fmt.Println(s & t)  //0010. AND
	fmt.Println(s | t)  //1011. OR
	fmt.Println(s ^ t)  //1001. XOR
	fmt.Println(s &^ t) //0100. AND NOT

	// Bit shifting
	u := 8              //2^3
	fmt.Println(u << 3) //Shift left by 3 bits. 2^3 * 2^3 = 64 == 2^6
	fmt.Println(u >> 3) //Shift right by 3 bits. 2^3 / 2^3 = 1 == 2^0

	// Floating point numbers
	var v float64 = 3.14
	v = 13.7e72
	v = 2e10
	fmt.Printf("%v %T\n", v, v)

	//Complex numbers
	var w complex64 = 1 + 2i
	fmt.Printf("%v %T\n", w, w)

	var x complex128 = complex(5, 12)
	fmt.Printf("%v %T\n", x, x)

	// Strings
	// Immutable
	c := "this is a string"           // UTF-8
	fmt.Printf("%v %T\n", c[2], c[2]) //105 uint8. Strings in Go are aliases for bytes.
	fmt.Printf("%v %T\n", string(c[2]), string(c[2]))

	// Convert string to slice of bytes
	d := []byte(c)
	fmt.Printf("%v %T\n", d, d) //ASCII characters. UTF-8.

	// Runes. Represents a UTF-32 character.
	// Type alias for int32.
	var e rune = 'a'
	fmt.Printf("%v %T\n", e, e)
	// Special methods for runes:
	// string.Reader.ReadRune()
}

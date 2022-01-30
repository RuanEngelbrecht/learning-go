package main

import "fmt"

// Methods
type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Ghost"
}

type counter int

func (c counter) increment() counter {
	return c + 1
}

func main() {
	// Basic functions
	sayMessage("Hello Go!")

	greeting := "Hello"
	name := "Stacy"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	total := sum(1, 2, 3, 4, 5)
	fmt.Println(*total)

	total2 := sum2(1, 2, 3, 4, 5)
	fmt.Println(total2)

	d, err := divide(10.0, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	// Anonymous functions -> IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("Anonymous function")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// Functions as variables -> func() is a type
	f := func(i int) {
		fmt.Println(i)
	}
	f(5)

	g := greeter{
		greeting: "Hello",
		name:     "Stacy",
	}
	g.greet()
	fmt.Println(g.name)

	var count counter = 1
	fmt.Println(count.increment())
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// Use pointers to avoid copying the value
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// Variatric function -> wraps values in a slice
// Can only have one variatic parameter
// Variatric parameters need to be the last parameter
func sum(values ...int) *int {
	total := 0
	for _, value := range values {
		total += value
	}
	return &total
}

// Named return values
func sum2(values ...int) (result int) {
	for _, value := range values {
		result += value
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

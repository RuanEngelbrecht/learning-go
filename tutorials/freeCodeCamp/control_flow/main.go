package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12801989,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	if pop, ok := statePopulations["Georgia"]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println("No data for Georgia")
	}

	// Switch statements.

	switch 1 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("another number")
	}

	// Switch statements with an initializer statement.

	switch i := 2 + 3; i {
	case 1, 5, 10: // Can test for multiple cases -> has to be unique.
		fmt.Println("one, five, or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("another number")
	}

	// Tagless switch statement.

	// Break statements are implied but you can still explicitly break out of a switch.
	// Can fall through to next case if needed. Use fallthrough keyword.

	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough // fallthrough is logicless! So it will always execute the next case.
	case j <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than 20")
	}

	// Type switch

	var k interface{} = [3]int{} // interface{} is a type that can hold any type.

	switch k.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("[3]int")
	default:
		fmt.Println("another type")
	}

	// Defer.

	// Defer is used to ensure that a function call is executed after all other
	// functions have been executed in the current function call stack.

	// Defer is useful for closing resources, or for waiting for a group of goroutines
	// to finish before continuing.
	// LIFO order.

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// Common pattern: group the opening and closing of a resource.
	// Easy to forget to close a resource.
	// Use defer to ensure that the resource is closed at the end of the function call.

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// When working with resources in a loop, you might not want to defer the closing of the resource.
	// You can delegate the closing of a resource to another function.

	// defer takes the value at the time of defer (eager), not the value at the time of the defer call.
	a := "a = start"
	defer fmt.Println(a)
	a = "a = end"

	// Panic.

	// In Go we don't have exceptions, we use error handling.
	// Panic is a built-in function that stops the execution of the current function and
	// returns control to the caller.

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// serverErr := http.ListenAndServe(":8080", nil)
	// if serverErr != nil {
	// 	panic(serverErr.Error())
	// }

	// Recover.

	// Recover is a built-in function that is used to handle panics.
	// Only useful in deferred functions.
	// Panics happen after defer statements are executed.
	// Current function will not attempt to continue, but higher functions in the call stack will.

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Recovered in f", err)
	// 		panic(err) // If you can't recover, you can rethrow the error.
	// 	}
	// }()
	// panic("something went wrong")
	// fmt.Println("end")
}

package main

import "fmt"

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
}

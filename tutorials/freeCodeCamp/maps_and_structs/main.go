package main

import (
	"fmt"
	"reflect"
)

// 2. Structs
// Declaration: struct struct-name{key: value, key: value}
// Collections of disparate data types that describe a single concept.
// When exporting a struct (uppercase), all fields must be exported as well.

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

// Embedded structs.
// Structs can be embedded in other structs.
// When to use: authoring a library / web framework / heavy base controller.

// Go does not support traditional Object Oriented Principles such as inheritance.
// Composition via embedding is the recommended approach.
// Use interfaces when you want to describe common behavior.

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // Embedded struct. Has-a relationship / characteristics.
	SpeedKPH float32
	CanFly   bool
}

// Tags
// Structs can be tagged with metadata.
// Validation / Encoding / Decoding / ORM / Defaults / Config / Field mapping / etc.

type Person struct {
	Name string `required:"true" max:"100"`
}

func main() {
	// 1. Maps
	// Key-value pairs.
	// Syntax: map[key-type]value-type{key: value}
	// Multiple assignments refer to the same underlying data.

	// Can use make function to create a map.
	a := make(map[string]int)
	a["x"] = 1
	a["y"] = 2
	a["z"] = 3
	fmt.Println(a)

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12801989,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))

	// Get value from map.
	fmt.Println(statePopulations["Ohio"])

	// Add a new key-value pair.
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations) // Return order of a map cannot guaranteed.

	// Delete a key-value pair.
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// Comma ok syntax to check if a key exists in a map.
	population, ok := statePopulations["Georgia"] // use _ to ignore the value.
	fmt.Println(population, ok)

	// 2. Structs
	// Structs are value types.

	aDoctor := Doctor{
		Number:     3,
		ActorName:  "Jon Pertwee",
		Companions: []string{"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.Companions[1])

	// Anonymous structs.
	// Can't use the struct anywhere else.
	// Very short-lived.
	bDoctor := struct{ name string }{name: "Jon Pertwee"}
	fmt.Println(bDoctor)

	anotherDoctor := bDoctor // Copy the struct i.e. independent dataset. Use pointer (&) to use same underlying data.
	anotherDoctor.name = "Tom Baker"
	fmt.Println(bDoctor)
	fmt.Println(anotherDoctor)

	// Embedded structs. See package level scope.
	aBird := Bird{}
	aBird.Name = "Emu"
	aBird.Origin = "Australia"
	aBird.SpeedKPH = 48
	aBird.CanFly = false
	fmt.Println(aBird)
	fmt.Println(aBird.Name)

	// Literal syntax for embedded struct.
	cBird := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(cBird)

	// Tags.
	// Using reflection to get tags and pass to some validation framework etc.
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

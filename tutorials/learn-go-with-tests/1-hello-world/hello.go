package main

import "fmt"

/*
	It is good to separate your "domain" code from the outside world (side-effects).
	The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.

	So let's separate these concerns so it's easier to test
*/

const spanish = "Spanish"
const french = "French"
const afrikaans = "Afrikaans"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const afrikaansHelloPrefix = "Awe, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case afrikaans:
		prefix = afrikaansHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

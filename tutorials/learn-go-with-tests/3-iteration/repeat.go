package iteration

import "strings"

func Repeat(char string, repetitions int) string {
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += char
	}
	return repeated
}

func RepeatUsingStdLib(char string, repetitions int) string {
	return strings.Repeat(char, repetitions)
}

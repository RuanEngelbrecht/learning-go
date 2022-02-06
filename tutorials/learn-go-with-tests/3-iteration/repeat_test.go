package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

// Example
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// Benchmarking
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeatUsingStdLib(t *testing.T) {
	repeated := RepeatUsingStdLib("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeatUsingStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatUsingStdLib("a", 5)
	}
}

/*
	Run "go test -bench=." to run all benchmarks.
	Function using the "strings" standard library is way faster than the one not using it and allocates less memory.
*/

package main

import "testing"

/*
	Writing a test is just like writing a function, with a few rules:
	- It needs to be in a file with a name like xxx_test.go
	- The test function must start with the word Test
	- The test function takes one argument only t *testing.T
	- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

	TDD Discipline:
	- Write a test
	- Make the compiler pass
	- Run the test, see that it fails and check the error message is meaningful
	- Write enough code to make the test pass
	- Refactor

	Why the TDD steps are important:
	- Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
	- Writing the smallest amount of code to make it pass so we know we have working software
	- Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
*/

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		//t.Helper() needed to tell the test suite that this method is a helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ruan", "")
		want := "Hello, Ruan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ruan", "Spanish")
		want := "Hola, Ruan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Ruan", "French")
		want := "Bonjour, Ruan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Afrikaans", func(t *testing.T) {
		got := Hello("Ruan", "Afrikaans")
		want := "Awe, Ruan"
		assertCorrectMessage(t, got, want)
	})
}

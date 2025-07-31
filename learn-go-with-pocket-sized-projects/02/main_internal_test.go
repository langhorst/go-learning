package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	want := "Hello world"

	got := greet()

	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}

}

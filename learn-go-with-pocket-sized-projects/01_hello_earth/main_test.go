package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello, world!
}

func TestGreeting(t *testing.T) {
	want := "Hello, world!"
	got := greet()

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

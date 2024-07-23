package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello, world!
}

func TestGreet_EN(t *testing.T) {
	lang := language("EN")
	want := "Hello, world!"
	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_FR(t *testing.T) {
	lang := language("FR")
	want := "Bonjour, le monde!"
	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_unsupported(t *testing.T) {
	lang := language("FOO")
	want := "unsupported language: FOO"
	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

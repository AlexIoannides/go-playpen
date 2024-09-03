package main

import "testing"

func TestSet_Contains(t *testing.T) {
	s := set{handmaidsTale: {}, oryxAndCrake: {}, theBellJar: {}}
	if !s.Contains(theBellJar) {
		t.Fatalf("expected set to contain book: %v", theBellJar)
	}
	if s.Contains(janeEyre) {
		t.Fatalf("expected set not to contain book: %v", janeEyre)
	}
}

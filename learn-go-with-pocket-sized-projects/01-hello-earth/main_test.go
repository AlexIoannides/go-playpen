package main

import (
	"os/exec"
	"testing"
)

func Example() {
	main()
	// Output:
	// Hello, world!
}

func Test(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--lang=FR")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("failed to run command: %v", err)
	}

	want := "Bonjour, le monde!\n"
	got := string(output)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var scenarios = map[string]testCase{
		"English":     {lang: "EN", want: "Hello, world!"},
		"French":      {lang: "FR", want: "Bonjour, le monde!"},
		"Unsupported": {lang: "FOO", want: "unsupported language: FOO"},
	}

	for name, scenario := range scenarios {
		t.Run(name, func(t *testing.T) {
			want := scenario.want
			got := greet(scenario.lang)

			if want != got {
				t.Errorf("expected: %q, got: %q", want, got)
			}
		})
	}
}

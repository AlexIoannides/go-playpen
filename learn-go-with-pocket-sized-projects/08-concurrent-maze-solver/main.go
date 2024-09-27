package main

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/08-concurrent-maze-solver/internal/solver"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)

	s, err := solver.New(inputFile)
	if err != nil {
		exit(err)
	}

	err = s.Solve()
	if err != nil {
		exit(err)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		exit(err)
	}
}

// usage displays the usage of the binary and exits the program.
func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}

// exit prints the error and exits the program.
func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}

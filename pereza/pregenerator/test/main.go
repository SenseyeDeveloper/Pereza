package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	intFixtures(pwd + "/fixtures/pregen/")
	benchmarks(pwd + "/benchmarks/pregen/")
}

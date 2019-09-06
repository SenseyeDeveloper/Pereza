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

	fixtures(pwd + "/fixtures/json/pregen/")
	benchmarks(pwd + "/benchmarks/json/pregen/")
}

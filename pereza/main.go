package main

import (
	"fmt"
	"github.com/senseyedeveloper/pereza/bootstrap"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "files required")

		os.Exit(1)
	}

	generator := bootstrap.NewGenerator(bootstrap.Settings{
		Filenames: files,
	})

	err := generator.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	// verbose
	fmt.Println("success generated")
}

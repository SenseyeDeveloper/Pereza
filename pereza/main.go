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

	generator.Run()
}

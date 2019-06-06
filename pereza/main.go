package main

import (
	"fmt"
	"github.com/senseyedeveloper/pereza/bootstrap"
	"github.com/senseyedeveloper/pereza/parser"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "files required")

		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Fprintln(os.Stderr, "GOPATH required")

		os.Exit(1)
	}

	generator := bootstrap.NewGenerator(
		parser.NewParser(gopath, pwd),
		bootstrap.Settings{
			Filenames: files,
		},
	)

	err = generator.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	// verbose
	fmt.Println("success generated")
}

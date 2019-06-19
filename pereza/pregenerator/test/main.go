package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	packageHeader = "package pregen\n\n"

	outputFilePattern = "%s_state.go"

	intStateStructPattern = "%s\ntype %sState struct {\nState %s `json:\"state\"`\n}\n"

	easy   = "Easy"
	pereza = "Pereza"
)

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	var (
		fixtures   = pwd + "/fixtures/pregen/"
		benchmarks = pwd + "/fixtures/benchmarks/"
	)

	_ = fixtures
	_ = benchmarks

	ints := []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
	}

	uints := []string{
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
	}

	types := make([]string, 0, len(ints)+len(uints))
	types = append(types, ints...)
	types = append(types, uints...)

	for _, t := range types {
		title := strings.Title(t)

		output := fixtures + fmt.Sprintf(outputFilePattern, t)

		content := []byte(packageHeader +
			fmt.Sprintf(intStateStructPattern, "", title, t) +
			fmt.Sprintf(intStateStructPattern, "// easyjson:json", easy+title, t) +
			fmt.Sprintf(intStateStructPattern, "// pereza:json", pereza+title, t),
		)

		ioutil.WriteFile(output, content, 0666)
	}
}

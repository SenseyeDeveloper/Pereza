package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	packageHeader = "package pregen\n\n"

	fixtureOutputFilePattern = "%s_state.go"

	intStatePattern = "%s\ntype %sState struct {\nState %s `json:\"state\"`\n}\n"

	easy   = "Easy"
	pereza = "Pereza"
)

func fixtures(path string) {
	types := make([]string, 0, len(ints)+len(uints))
	types = append(types, ints...)
	types = append(types, uints...)

	for _, t := range types {
		title := strings.Title(t)

		output := path + fmt.Sprintf(fixtureOutputFilePattern, t)

		content := []byte(packageHeader +
			fmt.Sprintf(intStatePattern, "", title, t) +
			fmt.Sprintf(intStatePattern, "// easyjson:json", easy+title, t) +
			fmt.Sprintf(intStatePattern, "// pereza:json", pereza+title, t),
		)

		err := ioutil.WriteFile(output, content, 0666)

		if err != nil {
			log.Fatalf("store %s fixture with %+v", t, err)
		}
	}
}

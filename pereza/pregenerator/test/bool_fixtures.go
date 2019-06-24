package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	boolStateStructStartPattern = "%s\ntype %sBoolState struct {\n"
	boolStateStructFieldPattern = "%c bool `json:\"%c\"`\n"
	boolStateStructEnd          = "}\n"
)

func boolFixtures(path string) {
	hex := boolFields(16)
	octo := hex[:8]

	typeFields := map[string][]string{
		"octo": octo,
		"hex":  hex,
	}

	for _, t := range bools {
		title := strings.Title(t)

		output := path + fmt.Sprintf(fixtureOutputFilePattern, t+"_bool")

		fields := strings.Join(typeFields[t], "")

		content := []byte(packageHeader +
			boolFixture("", title, fields) +
			boolFixture("// easyjson:json", easy+title, fields) +
			boolFixture("// pereza:json", pereza+title, fields),
		)

		err := ioutil.WriteFile(output, content, 0666)

		if err != nil {
			log.Fatalf("store %s fixture with %+v", t, err)
		}
	}
}

func boolFields(size int) []string {
	result := make([]string, size)

	for i := 0; i < size; i++ {
		result[i] = fmt.Sprintf(boolStateStructFieldPattern, byte('A'+i), byte('a'+i))
	}

	return result
}

func boolFixture(comment, title string, fields string) string {
	return fmt.Sprintf(boolStateStructStartPattern, comment, title) +
		fields +
		boolStateStructEnd
}

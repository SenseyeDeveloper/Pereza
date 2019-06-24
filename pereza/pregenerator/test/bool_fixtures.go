package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	boolStateStructStartPattern       = "%s\ntype %sBoolState struct {\n"
	boolStateStructFieldPattern       = "%c bool `json:\"%c\"`\n"
	boolStateStructDoubleFieldPattern = "%c%c bool `json:\"%c%c\"`\n"
	boolStateStructEnd                = "}\n"
)

func boolFixtures(path string) {
	hex := boolFields(16)
	octo := hex[:8]

	typeFields := map[string][]string{
		"octo":     octo,
		"hex":      hex,
		"alphabet": squareBoolFields(26),
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

func squareBoolFields(size int) []string {
	result := make([]string, 0, size*size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result = append(
				result,
				fmt.Sprintf(boolStateStructDoubleFieldPattern, byte('A'+i), byte('A'+j), byte('a'+i), byte('a'+j)),
			)
		}
	}

	return result
}

func boolFixture(comment, title string, fields string) string {
	return fmt.Sprintf(boolStateStructStartPattern, comment, title) +
		fields +
		boolStateStructEnd
}

package main

const (
	packageHeader = "package pregen\n\n"

	fixtureOutputFilePattern = "%s_state.go"

	easy   = "Easy"
	pereza = "Pereza"
)

func fixtures(path string) {
	intFixtures(path)
	boolFixtures(path)
}

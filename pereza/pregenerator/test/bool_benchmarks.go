package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	benchmarkBoolPattern = `package pregen

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"testing"
)

func Benchmark{Size}BoolStateEncodingJSON(b *testing.B) {
	source := fixtures.{Size}BoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func Benchmark{Size}BoolStateEasyJSON(b *testing.B) {
	source := fixtures.Easy{Size}BoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func Benchmark{Size}BoolStatePerezaJSON(b *testing.B) {
	source := fixtures.Pereza{Size}BoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
`
)

func boolBenchmarks(path string) {
	names := []string{
		"octo",
		"hexa",
	}

	for _, name := range names {
		output := path + fmt.Sprintf(benchmarkOutputFilePattern, name)

		replacer := strings.NewReplacer(
			"{Size}", strings.Title(name),
		)

		err := ioutil.WriteFile(output, []byte(replacer.Replace(benchmarkBoolPattern)), 0666)

		if err != nil {
			log.Fatalf("store %s benchmark with %+v", name, err)
		}
	}
}

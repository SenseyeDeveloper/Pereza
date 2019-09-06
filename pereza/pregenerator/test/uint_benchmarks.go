package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const (
	benchmarkUintPattern = `package pregen

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/json/pregen"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var (
	{uint}StateDataProvider = map[{uint}]string{
		0: "{\"state\":0}",
		1: "{\"state\":1}",
		{mathMax}: "{\"state\":{max}}",
	}
)

func Test{Uint}StateEncodingJSON(t *testing.T) {
	for state, expect := range {uint}StateDataProvider {
		source := pregen.{Uint}State{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Test{Uint}StateEasyJSON(t *testing.T) {
	for state, expect := range {uint}StateDataProvider {
		source := pregen.Easy{Uint}State{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Test{Uint}StatePereza(t *testing.T) {
	for state, expect := range {uint}StateDataProvider {
		source := pregen.Pereza{Uint}State{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Benchmark{Uint}StateEncodingJSON(b *testing.B) {
	source := pregen.{Uint}State{
		State: {mathMax},
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func Benchmark{Uint}StateEasyJSON(b *testing.B) {
	source := pregen.Easy{Uint}State{
		State: {mathMax},
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func Benchmark{Uint}StatePerezaJSON(b *testing.B) {
	source := pregen.Pereza{Uint}State{
		State: {mathMax},
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
`
)

func uintBenchmarks(path string) {
	minMaxMap := map[string]minmax{
		"uint":   {"", fmt.Sprintf("%d", uint(math.MaxUint64)), "", "math.MaxUint64"},
		"uint8":  {"", fmt.Sprintf("%d", uint(math.MaxUint8)), "", "math.MaxUint8"},
		"uint16": {"", fmt.Sprintf("%d", uint(math.MaxUint16)), "", "math.MaxUint16"},
		"uint32": {"", fmt.Sprintf("%d", uint(math.MaxUint32)), "", "math.MaxUint32"},
		"uint64": {"", fmt.Sprintf("%d", uint(math.MaxUint64)), "", "math.MaxUint64"},
	}

	for _, t := range uints {
		title := strings.Title(t)

		output := path + fmt.Sprintf(benchmarkOutputFilePattern, t)

		minmax := minMaxMap[t]

		replacer := strings.NewReplacer(
			"{uint}", t,
			"{Uint}", title,
			"{max}", minmax.max,
			"{mathMax}", minmax.mathMax,
		)

		err := ioutil.WriteFile(output, []byte(replacer.Replace(benchmarkUintPattern)), 0666)

		if err != nil {
			log.Fatalf("store %s benchmark with %+v", t, err)
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const (
	benchmarkIntPattern = `package pregen

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/pregen"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var (
	{int}StateDataProvider = map[{int}]string{
		{mathMin}: "{\"state\":{min}}",
		0:  "{\"state\":0}",
		{mathMax}:  "{\"state\":{max}}",
	}
)

func Test{Int}StateEncodingJSON(t *testing.T) {
	for state, expect := range {int}StateDataProvider {
		source := pregen.{Int}State{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Test{Int}StateEasyJSON(t *testing.T) {
	for state, expect := range {int}StateDataProvider {
		source := pregen.Easy{Int}State{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Test{Int}StatePereza(t *testing.T) {
	for state, expect := range {int}StateDataProvider {
		source := pregen.Pereza{Int}State{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func Benchmark{Int}StateEncodingJSON(b *testing.B) {
	source := pregen.{Int}State{
		State: {mathMin},
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func Benchmark{Int}StateEasyJSON(b *testing.B) {
	source := pregen.Easy{Int}State{
		State: {mathMin},
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func Benchmark{Int}StatePerezaJSON(b *testing.B) {
	source := pregen.Pereza{Int}State{
		State: {mathMin},
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
`
)

func intBenchmarks(path string) {
	minMaxMap := map[string]minmax{
		"int":   {fmt.Sprintf("%d", math.MinInt64), fmt.Sprintf("%d", math.MaxInt64), "math.MinInt64", "math.MaxInt64"},
		"int8":  {fmt.Sprintf("%d", math.MinInt8), fmt.Sprintf("%d", math.MaxInt8), "math.MinInt8", "math.MaxInt8"},
		"int16": {fmt.Sprintf("%d", math.MinInt16), fmt.Sprintf("%d", math.MaxInt16), "math.MinInt16", "math.MaxInt16"},
		"int32": {fmt.Sprintf("%d", math.MinInt32), fmt.Sprintf("%d", math.MaxInt32), "math.MinInt32", "math.MaxInt32"},
		"int64": {fmt.Sprintf("%d", math.MinInt64), fmt.Sprintf("%d", math.MaxInt64), "math.MinInt64", "math.MaxInt64"},
	}

	for _, t := range ints {
		title := strings.Title(t)

		output := path + fmt.Sprintf(benchmarkOutputFilePattern, t)

		minmax := minMaxMap[t]

		replacer := strings.NewReplacer(
			"{int}", t,
			"{Int}", title,
			"{min}", minmax.min,
			"{max}", minmax.max,
			"{mathMin}", minmax.mathMin,
			"{mathMax}", minmax.mathMax,
		)

		err := ioutil.WriteFile(output, []byte(replacer.Replace(benchmarkIntPattern)), 0666)

		if err != nil {
			log.Fatalf("store %s benchmark with %+v", t, err)
		}
	}
}

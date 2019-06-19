package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	const (
		output = `package core

import "reflect"

type IntSizeComment struct {
	Size    int
	Comment string
}

var reflectIntSize = map[reflect.Kind]IntSizeComment{
%s
}

func IntToStringMaxSize(t reflect.Kind) IntSizeComment {
	return reflectIntSize[t]
}
`
		pattern = "\t%s:{\n\tSize: %d,\n\tComment: \"len(`%s`)\",\n\t},"
	)

	type reflectSize struct {
		reflectNameAsKey string
		intAsString      string
	}

	// https://stackoverflow.com/questions/16474594/how-can-i-print-out-an-constant-uint64-in-go-using-fmt
	ints := []reflectSize{
		{"reflect.Int", fmt.Sprintf("%d", math.MinInt64)},
		{"reflect.Int8", fmt.Sprintf("%d", math.MaxInt8)},
		{"reflect.Int16", fmt.Sprintf("%d", math.MaxInt16)},
		{"reflect.Int32", fmt.Sprintf("%d", math.MaxInt32)},
		{"reflect.Int64", fmt.Sprintf("%d", math.MinInt64)},
		{"reflect.Uint", fmt.Sprintf("%d", uint64(math.MaxUint64))},
		{"reflect.Uint8", fmt.Sprintf("%d", math.MaxUint8)},
		{"reflect.Uint16", fmt.Sprintf("%d", math.MaxUint16)},
		{"reflect.Uint32", fmt.Sprintf("%d", math.MaxUint32)},
		{"reflect.Uint64", fmt.Sprintf("%d", uint64(math.MaxUint64))},
	}

	result := make([]string, len(ints))

	for i, value := range ints {
		result[i] = fmt.Sprintf(pattern, value.reflectNameAsKey, len(value.intAsString), value.intAsString)
	}

	fmt.Printf(output, strings.Join(result, "\n"))
}

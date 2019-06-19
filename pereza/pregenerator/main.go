package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	const (
		output = `package pregen

import "reflect"

type IntSizeComment struct {
	Size         int
	SizeAsString string
	Comment      string
	TypeCast     bool
	Signed       bool
}

var reflectIntSize = map[reflect.Kind]IntSizeComment{
%s
}

func IntToStringMaxSize(t reflect.Kind) IntSizeComment {
	return reflectIntSize[t]
}
`
		pattern = "%s:{\nSize: %d,\nSizeAsString: %q,\nComment: %q,\nTypeCast: %t,\nSigned: %t,\n},"
	)

	type reflectSize struct {
		reflectNameAsKey string
		intAsString      string
		typeCast         bool
		signed           bool
	}

	// https://stackoverflow.com/questions/16474594/how-can-i-print-out-an-constant-uint64-in-go-using-fmt
	ints := []reflectSize{
		{"reflect.Int", fmt.Sprintf("%d", math.MinInt64), true, true},
		{"reflect.Int8", fmt.Sprintf("%d", math.MaxInt8), true, true},
		{"reflect.Int16", fmt.Sprintf("%d", math.MaxInt16), true, true},
		{"reflect.Int32", fmt.Sprintf("%d", math.MaxInt32), true, true},
		{"reflect.Int64", fmt.Sprintf("%d", math.MinInt64), false, true},
		{"reflect.Uint", fmt.Sprintf("%d", uint64(math.MaxUint64)), true, false},
		{"reflect.Uint8", fmt.Sprintf("%d", math.MaxUint8), true, false},
		{"reflect.Uint16", fmt.Sprintf("%d", math.MaxUint16), true, false},
		{"reflect.Uint32", fmt.Sprintf("%d", math.MaxUint32), true, false},
		{"reflect.Uint64", fmt.Sprintf("%d", uint64(math.MaxUint64)), false, false},
	}

	result := make([]string, len(ints))

	for i, value := range ints {
		size := len(value.intAsString)
		sizeAsString := strconv.Itoa(size)
		result[i] = fmt.Sprintf(pattern,
			value.reflectNameAsKey,
			size,
			sizeAsString,
			value.intAsString,
			value.typeCast,
			value.signed,
		)
	}

	fmt.Printf(output, strings.Join(result, "\n"))
}

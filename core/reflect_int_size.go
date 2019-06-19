package core

import "reflect"

type IntSizeComment struct {
	Size         int
	SizeAsString string
	Comment      string
	TypeCast     bool
	Signed       bool
}

var reflectIntSize = map[reflect.Kind]IntSizeComment{
	reflect.Int: {
		Size:         20,
		SizeAsString: "20",
		Comment:      "-9223372036854775808",
		TypeCast:     true,
		Signed:       true,
	},
	reflect.Int8: {
		Size:         3,
		SizeAsString: "3",
		Comment:      "127",
		TypeCast:     true,
		Signed:       true,
	},
	reflect.Int16: {
		Size:         5,
		SizeAsString: "5",
		Comment:      "32767",
		TypeCast:     true,
		Signed:       true,
	},
	reflect.Int32: {
		Size:         10,
		SizeAsString: "10",
		Comment:      "2147483647",
		TypeCast:     true,
		Signed:       true,
	},
	reflect.Int64: {
		Size:         20,
		SizeAsString: "20",
		Comment:      "-9223372036854775808",
		TypeCast:     false,
		Signed:       true,
	},
	reflect.Uint: {
		Size:         20,
		SizeAsString: "20",
		Comment:      "18446744073709551615",
		TypeCast:     true,
		Signed:       false,
	},
	reflect.Uint8: {
		Size:         3,
		SizeAsString: "3",
		Comment:      "255",
		TypeCast:     true,
		Signed:       false,
	},
	reflect.Uint16: {
		Size:         5,
		SizeAsString: "5",
		Comment:      "65535",
		TypeCast:     true,
		Signed:       false,
	},
	reflect.Uint32: {
		Size:         10,
		SizeAsString: "10",
		Comment:      "4294967295",
		TypeCast:     true,
		Signed:       false,
	},
	reflect.Uint64: {
		Size:         20,
		SizeAsString: "20",
		Comment:      "18446744073709551615",
		TypeCast:     false,
		Signed:       false,
	},
}

func IntToStringMaxSize(t reflect.Kind) IntSizeComment {
	return reflectIntSize[t]
}

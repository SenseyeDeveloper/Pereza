package core

import "reflect"

type IntSizeComment struct {
	Size    int
	Comment string
}

var reflectIntSize = map[reflect.Kind]IntSizeComment{
	reflect.Int: {
		Size:    20,
		Comment: "len(`-9223372036854775808`)",
	},
	reflect.Int8: {
		Size:    3,
		Comment: "len(`127`)",
	},
	reflect.Int16: {
		Size:    5,
		Comment: "len(`32767`)",
	},
	reflect.Int32: {
		Size:    10,
		Comment: "len(`2147483647`)",
	},
	reflect.Int64: {
		Size:    20,
		Comment: "len(`-9223372036854775808`)",
	},
	reflect.Uint: {
		Size:    20,
		Comment: "len(`18446744073709551615`)",
	},
	reflect.Uint8: {
		Size:    3,
		Comment: "len(`255`)",
	},
	reflect.Uint16: {
		Size:    5,
		Comment: "len(`65535`)",
	},
	reflect.Uint32: {
		Size:    10,
		Comment: "len(`4294967295`)",
	},
	reflect.Uint64: {
		Size:    20,
		Comment: "len(`18446744073709551615`)",
	},
}

func IntToStringMaxSize(t reflect.Kind) IntSizeComment {
	return reflectIntSize[t]
}

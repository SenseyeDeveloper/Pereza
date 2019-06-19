package core

import "reflect"

var reflectIntSize = map[reflect.Kind]int{
	reflect.Int:    20, // len(`-9223372036854775808`)
	reflect.Int8:   3,  // len(`127`)
	reflect.Int16:  5,  // len(`32767`)
	reflect.Int32:  10, // len(`2147483647`)
	reflect.Int64:  20, // len(`-9223372036854775808`)
	reflect.Uint:   20, // len(`18446744073709551615`)
	reflect.Uint8:  3,  // len(`255`)
	reflect.Uint16: 5,  // len(`65535`)
	reflect.Uint32: 10, // len(`4294967295`)
	reflect.Uint64: 20, // len(`18446744073709551615`)
}

func IntToStringMaxSize(t reflect.Kind) int {
	return reflectIntSize[t]
}

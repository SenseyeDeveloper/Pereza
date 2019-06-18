package gen

import (
	"reflect"
	"strings"
)

func getTagName(f reflect.StructField) string {
	json := f.Tag.Get("json")

	index := strings.IndexByte(json, ',')

	if index == -1 {
		return json
	}

	return json[:index]
}

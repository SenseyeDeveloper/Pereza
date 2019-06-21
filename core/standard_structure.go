package core

import "reflect"

func StandardStructureField(field reflect.StructField) (string, bool) {
	tags := ParseFieldTags(field.Tag.Get("json"))

	if tags.OnlyName() {
		return tags.Name, true
	}

	return "", false
}

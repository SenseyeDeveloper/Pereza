package core

import "reflect"

func StandardStructureField(field reflect.StructField) (string, bool) {
	tags := ParseFieldTags(field.Tag.Get("json"))

	if tags.OnlyName() {
		return tags.Name, true
	}

	return "", false
}

func MultiBoolStandardStructure(t reflect.Type) ([]string, []string, bool) {
	length := t.NumField()

	fieldsNames := make([]string, length)
	jsonNames := make([]string, length)

	for i := 0; i < length; i++ {
		field := t.Field(i)

		jsonName, standard := StandardStructureField(field)

		if !standard {
			return nil, nil, false
		}

		fieldsNames[i] = field.Name
		jsonNames[i] = jsonName
	}

	return fieldsNames, jsonNames, true
}

func MatchAllBooleanFields(t reflect.Type) bool {
	length := t.NumField()

	for i := 0; i < length; i++ {
		field := t.Field(i)

		if field.Type.Kind() != reflect.Bool {
			return false
		}
	}

	return true
}

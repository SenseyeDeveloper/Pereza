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

		if field.Type.Kind() != reflect.Bool {
			return nil, nil, false
		}

		jsonName, standard := StandardStructureField(field)

		if standard {
			fieldsNames[i] = field.Name
			jsonNames[i] = jsonName

			continue
		}

		return nil, nil, false
	}

	return fieldsNames, jsonNames, true
}

package complexstub

import (
	"github.com/gopereza/pereza/core/boolstub"
	"reflect"
)

func StandardStub(t reflect.Type, fieldNames, jsonNames []string) []byte {
	wrapSize := WrapMultiSize(jsonNames)

	totalSize := wrapSize

	stub := make([]byte, 0)

	{
		var fieldSize int
		stub, fieldSize = AppendFirstField(stub, t.Field(0).Type.Kind(), fieldNames[0], jsonNames[0])
		totalSize += fieldSize
	}

}

func AppendFirstField(dst []byte, kind reflect.Kind, fieldName, jsonName string) ([]byte, int) {
	const firstFieldWrapSize = 4 // len(`{"":`)
	firstFieldSize := firstFieldWrapSize + len(jsonName)

	switch kind {
	case reflect.Bool:
		return boolstub.AppendFirstField(dst, boolstub.Condition(fieldName), jsonName), firstFieldSize + boolstub.MaxBoolSize
	}

	return nil, 0
}

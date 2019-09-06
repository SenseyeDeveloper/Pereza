package complexstub

import (
	"github.com/gopereza/pereza/core/boolstub"
	"github.com/gopereza/pereza/core/common"
	"github.com/gopereza/pereza/core/intstub"
	"github.com/gopereza/pereza/core/stringstub"
	"github.com/gopereza/pereza/pregen"
	"reflect"
	"strconv"
	"strings"
)

func StandardStub(t reflect.Type, fieldNames, jsonNames []string) []byte {
	wrapSize := WrapMultiSize(jsonNames)

	totalSize := make([]string, 0, 4)
	totalSize = append(totalSize, strconv.FormatUint(uint64(wrapSize), 10))

	var fieldSize string
	var fieldImports []string

	body := make([]byte, 0)
	imports := make([]string, 0)

	{
		body, fieldSize, fieldImports = AppendFirstField(body, t.Field(0).Type.Kind(), fieldNames[0], jsonNames[0])
		totalSize = append(totalSize, fieldSize)
		imports = append(imports, fieldImports...)
	}

	length := len(fieldNames)

	for i := 1; i < length; i++ {
		body, fieldSize, fieldImports = AppendField(body, t.Field(i).Type.Kind(), fieldNames[i], jsonNames[i])
		totalSize = append(totalSize, fieldSize)
		imports = append(imports, fieldImports...)
	}

	body = append(body, '\n')
	body = append(body, "\tresult = append(result, '}')"...)
	body = append(body, '\n')

	result := make([]byte, 0)

	result = common.AppendImports(result, imports)

	result = common.AppendHeader(result, t.Name())
	result = append(result, '\t')
	result = append(result, "result := make([]byte, 0, "+strings.Join(totalSize, " + ")+")"...)
	result = append(result, '\n')
	result = append(result, body...)
	result = append(result, '\n')
	result = append(result, '\t')
	result = append(result, "return result, nil"...)

	result = common.AppendFooter(result)

	return result
}

func AppendFirstField(dst []byte, kind reflect.Kind, fieldName, jsonName string) ([]byte, string, []string) {
	switch kind {
	case reflect.Bool:
		return boolstub.AppendFirstField(dst, boolstub.Condition(fieldName), jsonName), boolstub.MaxBoolSizeAsString, nil
	case reflect.String:
		return stringstub.AppendFirstField(dst, fieldName, jsonName), "len(v." + fieldName + ") + 2", nil
	}

	if intSizeComment, ok := pregen.IntToStringMaxSizeOk(kind); ok {
		return intstub.AppendFirstField(dst, fieldName, jsonName, intSizeComment), intSizeComment.SizeAsString, []string{intstub.IntImport}
	}

	return dst, "0", nil
}

func AppendField(dst []byte, kind reflect.Kind, fieldName, jsonName string) ([]byte, string, []string) {
	switch kind {
	case reflect.Bool:
		return boolstub.AppendField(dst, boolstub.Condition(fieldName), jsonName), boolstub.MaxBoolSizeAsString, []string{intstub.IntImport}
	case reflect.String:
		return stringstub.AppendField(dst, fieldName, jsonName), "len(v." + fieldName + ") + 2", nil
	}

	if intSizeComment, ok := pregen.IntToStringMaxSizeOk(kind); ok {
		return intstub.AppendField(dst, fieldName, jsonName, intSizeComment), intSizeComment.SizeAsString, nil
	}

	return dst, "0", nil
}

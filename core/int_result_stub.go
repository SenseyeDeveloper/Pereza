package core

import (
	"github.com/gopereza/pereza/pregen"
	"reflect"
	"strconv"
)

const (
	intImport = `import "strconv"

`
)

var signedTypeCastSize = map[bool]map[bool]int{
	true: {
		true:  7,
		false: 0,
	},
	false: {
		true:  9,
		false: 1,
	},
}

/**
// MarshalJSON supports json.Marshaler interface
func (v *PerezaIntState) PerezaMarshalJSON() []byte {
	const start = 9  // len([]byte(`{"state":`))
	const value = 11 // len([]byte(`-2147483648`))
	const end = 1    // len([]byte(`}`))

	result := make([]byte, 0, start+value+end)

	result = append(result, '{', '"', 's', 't', 'a', 't', 'e', '"', ':')
	result = strconv.AppendInt(result, int64(v.State), 10)
	result = append(result, '}')

	return result
}
*/

func IntResultStubByType(typeName, fieldName, jsonName string, t reflect.Kind) []byte {
	return IntResultStubBySettings(typeName, fieldName, jsonName, pregen.IntToStringMaxSize(t))
}

func IntResultStubBySettings(typeName, fieldName, jsonName string, comment pregen.IntSizeComment) []byte {
	result := make([]byte, 0, getIntResultStubSizeBySettings(typeName, fieldName, jsonName, comment))

	result = append(result, intImport...)
	result = append(result, resultStubHeader...)
	result = append(result, resultStubFuncSignatureStart...)
	result = append(result, typeName...)
	result = append(result, resultStubFuncSignatureEnd...)

	result = append(result, "	const start = "...)
	result = append(result, strconv.Itoa(getStringStartConst(jsonName))...)
	result = append(result, " // len([]byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":`))\n"...)
	result = append(result, "	const value = "...)
	result = append(result, comment.SizeAsString...)
	result = append(result, " // len(`"...)
	result = append(result, comment.Comment...)
	result = append(result, "`)\n"...)
	result = append(result, "	const end = 1    // len([]byte{'}'})\n\n"...)

	result = append(result, "	result := make([]byte, 0, start+value+end)\n"...)

	result = append(result, `	result = append(result, '{', '"'`...)
	result = appendJSONFieldNameAsBytes(result, jsonName)
	result = append(result, `, '"', ':')`...)
	result = append(result, n)

	if comment.Signed {
		if comment.TypeCast {
			result = append(result, `	result = strconv.AppendInt(result, int64(v.`...)
			result = append(result, fieldName...)
			result = append(result, `), 10)`...)
		} else {
			result = append(result, `	result = strconv.AppendInt(result, v.`...)
			result = append(result, fieldName...)
			result = append(result, `, 10)`...)
		}
	} else {
		if comment.TypeCast {
			result = append(result, `	result = strconv.AppendUint(result, uint64(v.`...)
			result = append(result, fieldName...)
			result = append(result, `), 10)`...)
		} else {
			result = append(result, `	result = strconv.AppendUint(result, v.`...)
			result = append(result, fieldName...)
			result = append(result, `, 10)`...)
		}
	}

	result = append(result, n)
	result = append(result, `	result = append(result, '}')`...)
	result = append(result, n, n)

	result = append(result, `	return result, nil`...)
	result = append(result, n, '}', n)

	return result
}

func getIntResultStubSizeByType(typeName, fieldName, jsonName string, t reflect.Kind) int {
	return getIntResultStubSizeBySettings(typeName, fieldName, jsonName, pregen.IntToStringMaxSize(t))
}

func getIntResultStubSizeBySettings(typeName, fieldName, jsonName string, comment pregen.IntSizeComment) int {
	const (
		fixedSize = len(intImport) +
			len(resultStubHeader) +
			len(resultStubFuncSignatureStart) +
			len(resultStubFuncSignatureEnd) +
			291 // func other
	)

	return fixedSize +
		intSize(getStringStartConst(jsonName)) +
		len(typeName) +
		len(fieldName) +
		6*len(jsonName) +
		len(comment.Comment) +
		len(comment.SizeAsString) +
		signedTypeCastSize[comment.Signed][comment.TypeCast]
}

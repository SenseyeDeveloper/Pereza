package intstub

import (
	"github.com/gopereza/pereza/core/common"
	"github.com/gopereza/pereza/pregen"
	"reflect"
	"strconv"
)

const (
	IntImport = `import "strconv"`
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

// Dynamic allocate
func IntResultStubByType(typeName, fieldName, jsonName string, t reflect.Kind) []byte {
	return IntResultStubBySettings(typeName, fieldName, jsonName, pregen.IntToStringMaxSize(t))
}

func IntResultStubBySettings(typeName, fieldName, jsonName string, comment pregen.IntSizeComment) []byte {
	result := make([]byte, 0)

	result = append(result, IntImport...)
	result = append(result, '\n', '\n')

	result = common.AppendHeader(result, typeName)

	result = append(result, "	const start = "...)
	result = strconv.AppendUint(result, uint64(common.StringStartConst(jsonName)), 10)
	result = append(result, " // len([]byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":`))\n"...)
	result = append(result, "	const value = "...)
	result = append(result, comment.SizeAsString...)
	result = append(result, " // len(`"...)
	result = append(result, comment.Comment...)
	result = append(result, "`)\n"...)
	result = append(result, "	const end = 1    // len([]byte{'}'})\n\n"...)

	result = append(result, "	result := make([]byte, 0, start+value+end)"...)

	result = AppendFirstField(result, fieldName, jsonName, comment)

	result = append(result, '\n')
	result = append(result, `	result = append(result, '}')`...)
	result = append(result, '\n', '\n')

	result = append(result, `	return result, nil`...)
	result = append(result, '\n', '}', '\n')

	return result
}

func AppendFirstField(result []byte, fieldName, jsonName string, comment pregen.IntSizeComment) []byte {
	result = append(result, '\n')
	result = append(result, "\tresult = append(result, `{"...)
	result = append(result, `"`...)
	result = append(result, jsonName...)
	result = append(result, `":`...)
	result = append(result, "`...)"...)

	result = append(result, '\n')
	result = appendInt(result, fieldName, comment)
	result = append(result, '\n')

	return result
}

func AppendField(result []byte, fieldName, jsonName string, comment pregen.IntSizeComment) []byte {
	result = append(result, '\n')
	result = append(result, "\tresult = append(result, `,"...)
	result = append(result, `"`...)
	result = append(result, jsonName...)
	result = append(result, `":`...)
	result = append(result, "`...)"...)

	result = append(result, '\n')
	result = appendInt(result, fieldName, comment)
	result = append(result, '\n')

	return result
}

func appendInt(result []byte, fieldName string, comment pregen.IntSizeComment) []byte {
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

	return result
}

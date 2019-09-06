package stringstub

import (
	"github.com/gopereza/pereza/core/common"
	"strconv"
)

/**
// MarshalJSON supports json.Marshaler interface
func (v *PerezaStringState) PerezaMarshalJSON() []byte {
	const start = 10 // len([]byte(`{"state":"`))
	const end = 2    // len([]byte{'"', '}'})

	result := make([]byte, 0, start+len(v.State)+end)

	result = append(result, '{', '"', 's', 't', 'a', 't', 'e', '"', ':', '"')
	result = append(result, v.State...)
	result = append(result, '"', '}')

	return result
}
*/

// Dynamic allocate
func StringResultStub(typeName, fieldName, jsonName string) []byte {
	result := make([]byte, 0, getStringResultStubSize(typeName, fieldName, jsonName))

	result = common.AppendHeader(result, typeName)

	result = append(result, "	const start = "...)
	result = strconv.AppendUint(result, uint64(common.StringStartConst(jsonName)), 10)
	result = append(result, " // len([]byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":\"`))\n"...)
	result = append(result, `	const end = 2    // len([]byte{'"', '}'})`...)
	result = append(result, '\n', '\n')

	result = append(result, `	result := make([]byte, 0, start+len(v.`...)
	result = append(result, fieldName...)
	result = append(result, `)+end)`...)
	result = append(result, '\n')

	result = AppendFirstField(result, fieldName, jsonName)
	result = append(result, `	result = append(result, '}')`...)
	result = append(result, '\n', '\n')

	result = append(result, `	return result, nil`...)
	result = append(result, '\n', '}', '\n')

	return result
}

func getStringResultStubSize(typeName, fieldName, jsonName string) int {
	const (
		fixedSize = common.WrapSignatureSize +
			268 // func other
	)

	return fixedSize +
		common.DigitsSize(common.StringStartConst(jsonName)) +
		len(typeName) +
		2*len(fieldName) +
		6*len(jsonName)
}

func AppendFirstField(result []byte, fieldName, jsonName string) []byte {
	result = append(result, "\tresult = append(result, `{"...)
	result = append(result, `"`...)
	result = append(result, jsonName...)
	result = append(result, `":"`...)
	result = append(result, "`...)"...)
	result = append(result, '\n')
	result = append(result, `	result = append(result, v.`...)
	result = append(result, fieldName...)
	result = append(result, `...)`...)
	result = append(result, '\n')
	result = append(result, `	result = append(result, '"')`...)
	result = append(result, '\n')

	return result
}

func AppendField(result []byte, fieldName, jsonName string) []byte {
	result = append(result, "\tresult = append(result, `,"...)
	result = append(result, `"`...)
	result = append(result, jsonName...)
	result = append(result, `":"`...)
	result = append(result, "`...)"...)
	result = append(result, '\n')
	result = append(result, `	result = append(result, v.`...)
	result = append(result, fieldName...)
	result = append(result, `...)`...)
	result = append(result, '\n')
	result = append(result, `	result = append(result, '"')`...)
	result = append(result, '\n')

	return result
}

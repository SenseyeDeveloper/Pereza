package core

import "strconv"

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

func StringResultStub(typeName, fieldName, jsonName string) []byte {
	result := make([]byte, 0, getStringResultStubSize(typeName, fieldName, jsonName))

	result = append(result, resultStubHeader...)
	result = append(result, resultStubFuncSignatureStart...)
	result = append(result, typeName...)
	result = append(result, resultStubFuncSignatureEnd...)

	result = append(result, "	const start = "...)
	result = append(result, strconv.Itoa(getStringStartConst(jsonName))...)
	result = append(result, " // len([]byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":\"`))\n"...)
	result = append(result, `	const end = 2    // len([]byte{'"', '}'})`...)
	result = append(result, n, n)

	result = append(result, `	result := make([]byte, 0, start+len(v.`...)
	result = append(result, fieldName...)
	result = append(result, `)+end)`...)
	result = append(result, n)

	result = append(result, `	result = append(result, '{', '"'`...)
	result = appendJSONFieldNameAsBytes(result, jsonName)
	result = append(result, `, '"', ':', '"')`...)
	result = append(result, n)
	result = append(result, `	result = append(result, v.`...)
	result = append(result, fieldName...)
	result = append(result, `...)`...)
	result = append(result, n)
	result = append(result, `	result = append(result, '"', '}')`...)
	result = append(result, n, n)

	result = append(result, `	return result, nil`...)
	result = append(result, n, '}', n)

	return result
}

func getStringResultStubSize(typeName, fieldName, jsonName string) int {
	const (
		fixedSize = len(resultStubHeader) +
			len(resultStubFuncSignatureStart) +
			len(resultStubFuncSignatureEnd) +
			270 // func other
	)

	return fixedSize +
		intSize(getStringStartConst(jsonName)) +
		len(typeName) +
		2*len(fieldName) +
		6*len(jsonName)
}

func appendJSONFieldNameAsBytes(bytes []byte, jsonName string) []byte {
	for _, l := range jsonName {
		bytes = append(bytes, ',', ' ', '\'', byte(l), '\'')
	}

	return bytes
}

func getStringStartConst(jsonName string) int {
	// len([]byte(`{"json_name":"`))

	const wrapperSize = 5

	return wrapperSize + len(jsonName)
}

func intSize(i int) int {
	result := 0

	for i > 0 {
		i /= 10

		result += 1
	}

	return result
}

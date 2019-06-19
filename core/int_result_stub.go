package core

import "strconv"

const (
	intImport = `import "strconv"

`
)

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

func IntResultStub(typeName, fieldName, jsonName string) []byte {
	result := make([]byte, 0, getIntResultStubSize(typeName, fieldName, jsonName))

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
	result = append(result, "	const value = 20 // len([]byte(`-9223372036854775808`))\n"...)
	result = append(result, `	const end = 1    // len([]byte{'}'})`...)
	result = append(result, n, n)

	result = append(result, "	result := make([]byte, 0, start+value+end)\n"...)

	result = append(result, `	result = append(result, '{', '"'`...)
	result = appendJSONFieldNameAsBytes(result, jsonName)
	result = append(result, `, '"', ':')`...)
	result = append(result, n)

	result = append(result, `	result = strconv.AppendInt(result, int64(v.`...)
	result = append(result, fieldName...)
	result = append(result, `), 10)`...)

	result = append(result, n)
	result = append(result, `	result = append(result, '}')`...)
	result = append(result, n, n)

	result = append(result, `	return result, nil`...)
	result = append(result, n, '}', n)

	return result
}

func getIntResultStubSize(typeName, fieldName, jsonName string) int {
	const (
		fixedSize = len(intImport) +
			len(resultStubHeader) +
			len(resultStubFuncSignatureStart) +
			len(resultStubFuncSignatureEnd) +
			328 // func other
	)

	return fixedSize +
		intSize(getStringStartConst(jsonName)) +
		len(typeName) +
		len(fieldName) +
		6*len(jsonName)
}

package core

/**
// MarshalJSON supports json.Marshaler interface
func (v PerezaBoolState) MarshalJSON() []byte {
	if v.State {
		return []byte(`{"state":true}`), nil
	}

	return []byte(`{"state":false}`), nil
}
*/

func BoolResultStub(typeName, fieldName, jsonName string) []byte {
	result := make([]byte, 0, getBoolResultStubSize(typeName, fieldName, jsonName))

	result = append(result, resultStubHeader...)
	result = append(result, resultStubFuncSignatureStart...)
	result = append(result, typeName...)
	result = append(result, resultStubFuncSignatureEnd...)

	result = append(result, "	if v."...)
	result = append(result, fieldName...)
	result = append(result, ' ', '{', n)
	result = append(result, "		return []byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":true}`), nil\n"...)
	result = append(result, "	}"...)
	result = append(result, n, n)

	result = append(result, "	return []byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":false}`), nil"...)
	result = append(result, n, '}', n)

	return result
}

func getBoolResultStubSize(typeName, fieldName, jsonName string) int {
	const (
		fixedSize = len(resultStubHeader) +
			len(resultStubFuncSignatureStart) +
			len(resultStubFuncSignatureEnd) +
			83 // func other
	)

	return fixedSize + len(typeName) + len(fieldName) + 2*len(jsonName)
}

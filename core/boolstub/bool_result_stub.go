package boolstub

import "github.com/gopereza/pereza/core/common"

/**
// MarshalJSON supports json.Marshaler interface
func (v PerezaBoolState) MarshalJSON() []byte {
	if v.State {
		return []byte(`{"state":true}`), nil
	}

	return []byte(`{"state":false}`), nil
}

// MarshalJSON supports json.Marshaler interface
func (v PerezaBoolState) MarshalJSON() ([]byte, error) {
	if v.State {
		return []byte{'{', '"', 's', 't', 'a', 't', 'e', '"', ':', 't', 'r', 'u', 'e', '}'}, nil
	}

	return []byte{'{', '"', 's', 't', 'a', 't', 'e', '"', ':', 'f', 'a', 'l', 's', 'e', '}'}, nil
}
*/

// Static allocate
func BoolResultStub(typeName, fieldName, jsonName string) []byte {
	result := make([]byte, 0, getBoolResultStubSize(typeName, fieldName, jsonName))

	result = common.AppendHeader(result, typeName)

	result = append(result, "	if v."...)
	result = append(result, fieldName...)
	result = append(result, ' ', '{', '\n')
	result = append(result, "		return []byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":true}`), nil\n"...)
	result = append(result, "	}"...)
	result = append(result, '\n', '\n')

	result = append(result, "	return []byte(`{\""...)
	result = append(result, jsonName...)
	result = append(result, "\":false}`), nil"...)
	result = append(result, '\n', '}', '\n')

	return result
}

// Dynamic allocate
func CombinatorBoolResultStub(typeName string, fieldNames, jsonNames []string) []byte {
	generator := NewCombinatorGenerator(fieldNames, jsonNames)

	body := generator.Generate()

	result := make([]byte, 0, common.WrapSignatureSize+len(body))

	result = common.AppendHeader(result, typeName)
	result = append(result, body...)
	result = common.AppendFooter(result)

	return result
}

func getBoolResultStubSize(typeName, fieldName, jsonName string) int {
	const (
		fixedSize = common.WrapSignatureSize +
			81 // func other
	)

	return fixedSize + len(typeName) + len(fieldName) + 2*len(jsonName)
}

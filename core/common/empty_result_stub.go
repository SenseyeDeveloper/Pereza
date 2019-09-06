package common

/**
// MarshalJSON supports json.Marshaler interface
func (PerezaEmptyState) MarshalJSON() ([]byte, error) {
	return []byte{'{', '}'}, nil
}
*/

const (
	emptyResultStubFuncStart = "func ("
	emptyResultStubFuncBody  = "	return []byte{'{', '}'}, nil\n"
)

// Static allocate
func EmptyResultStub(name string) []byte {
	result := make([]byte, 0, getEmptyResultStubSize(name))

	result = append(result, resultStubHeader...)
	result = append(result, emptyResultStubFuncStart...)
	result = append(result, name...)
	result = append(result, resultStubFuncSignatureEnd...)
	result = append(result, emptyResultStubFuncBody...)
	result = append(result, '}', '\n')

	return result
}

func getEmptyResultStubSize(name string) int {
	const (
		fixedSize = len(resultStubHeader) +
			len(emptyResultStubFuncStart) +
			len(resultStubFuncSignatureEnd) +
			len(emptyResultStubFuncBody) +
			2 // func end
	)

	return fixedSize + len(name)
}

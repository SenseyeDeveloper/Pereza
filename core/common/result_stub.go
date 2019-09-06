package common

const (
	resultStubHeader             = "// MarshalJSON supports json.Marshaler interface\n"
	resultStubFuncSignatureStart = "func (v "
	resultStubFuncSignatureEnd   = ") MarshalJSON() ([]byte, error) {\n"
)

const (
	WrapSignatureSize = len(resultStubHeader) +
		len(resultStubFuncSignatureStart) +
		len(resultStubFuncSignatureEnd) +
		2 // '\n', '}'
)

func AppendImports(result []byte, imports []string) []byte {
	unique := make(map[string]bool)

	for _, importString := range imports {
		if unique[importString] {
			continue
		}

		unique[importString] = true

		result = append(result, importString...)
		result = append(result, '\n')
	}

	return result
}

func AppendHeader(result []byte, typeName string) []byte {
	result = append(result, resultStubHeader...)
	result = append(result, resultStubFuncSignatureStart...)
	result = append(result, typeName...)
	result = append(result, resultStubFuncSignatureEnd...)

	return result
}

func AppendFooter(source []byte) []byte {
	return append(source, '\n', '}')
}

package core

const (
	resultStubHeader             = "// MarshalJSON supports json.Marshaler interface\n"
	resultStubFuncSignatureStart = "func (v "
	resultStubFuncSignatureEnd   = ") MarshalJSON() ([]byte, error) {\n"
)

const (
	wrapSignatureSize = len(resultStubHeader) +
		len(resultStubFuncSignatureStart) +
		len(resultStubFuncSignatureEnd) +
		2 // '\n', '}'
)

func appendHeader(source []byte, typeName string) []byte {
	result := append(source, resultStubHeader...)
	result = append(result, resultStubFuncSignatureStart...)
	result = append(result, typeName...)
	result = append(result, resultStubFuncSignatureEnd...)

	return result
}

func appendFooter(source []byte) []byte {
	return append(source, n, '}')
}

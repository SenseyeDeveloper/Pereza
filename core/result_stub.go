package core

const (
	resultStubHeader             = "// MarshalJSON supports json.Marshaler interface\n"
	resultStubFuncSignatureStart = "func (v "
	resultStubFuncSignatureEnd   = ") MarshalJSON() ([]byte, error) {\n"
)

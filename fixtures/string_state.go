package fixtures

const (
	ExpectStringStateEmpty = `{"state":""}`
	ExpectStringStateSmall = `{"state":"abc"}`
)

type StringState struct {
	State string `json:"state"`
}

// easyjson:json
type EasyStringState struct {
	State string `json:"state"`
}

type PerezaStringState struct {
	State string `json:"state"`
}

func (v *PerezaStringState) PerezaMarshalJSON() []byte {
	const start = 10 // len([]byte(`{"state":"`))
	const end = 2    // len([]byte{'"', "}"})

	result := make([]byte, 0, start+len(v.State)+end)

	result = append(result, '{', '"', 's', 't', 'a', 't', 'e', '"', ':', '"')
	result = append(result, v.State...)
	result = append(result, '"', '}')

	return result
}

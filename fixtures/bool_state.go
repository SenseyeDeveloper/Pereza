package fixtures

const (
	ExpectBoolStateTrue  = `{"state":true}`
	ExpectBoolStateFalse = `{"state":false}`
)

type BoolState struct {
	State bool `json:"state"`
}

// easyjson:json
type EasyBoolState struct {
	State bool `json:"state"`
}

type PerezaBoolState struct {
	State bool `json:"state"`
}

func (v *PerezaBoolState) PerezaMarshalJSON() []byte {
	if v.State {
		return []byte(`{"state":true}`)
	}

	return []byte(`{"state":false}`)
}

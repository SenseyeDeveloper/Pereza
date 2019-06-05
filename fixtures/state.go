package fixtures

const (
	ExpectBoolStateTrue  = `{"state":true}`
	ExpectBoolStateFalse = `{"state":false}`
)

type BoolState struct {
	State bool `json:"state"`
}

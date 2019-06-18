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

// pereza:json
type PerezaStringState struct {
	State string `json:"state"`
}

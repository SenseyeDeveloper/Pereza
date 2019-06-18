package fixtures

const (
	ExpectIntStateN1 = `{"state":-1}`
	ExpectIntState0  = `{"state":0}`
	ExpectIntState1  = `{"state":1}`
)

type IntState struct {
	State int `json:"state"`
}

// easyjson:json
type EasyIntState struct {
	State int `json:"state"`
}

// pereza:json
type PerezaIntState struct {
	State int `json:"state"`
}

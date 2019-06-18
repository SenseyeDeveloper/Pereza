package fixtures

const (
	ExpectUintState0 = `{"state":0}`
	ExpectUintState1 = `{"state":1}`
)

type UintState struct {
	State uint `json:"state"`
}

// easyjson:json
type EasyUintState struct {
	State uint `json:"state"`
}

// pereza:json
type PerezaUintState struct {
	State uint `json:"state"`
}

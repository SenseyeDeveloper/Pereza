package boolean

type DoubleBoolState struct {
	Active  bool `json:"active"`
	Approve bool `json:"approve"`
}

// easyjson:json
type EasyDoubleBoolState struct {
	Active  bool `json:"active"`
	Approve bool `json:"approve"`
}

// pereza:json
type PerezaDoubleBoolState struct {
	Active  bool `json:"active"`
	Approve bool `json:"approve"`
}

package fixtures

import "strconv"

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

type PerezaIntState struct {
	State int `json:"state"`
}

func (v *PerezaIntState) PerezaMarshalJSON() []byte {
	const start = 9  // len([]byte(`{"state":`))
	const end = 1    // len([]byte(`}`))
	const value = 11 // len([]byte(`-2147483648`))

	result := make([]byte, 0, start+value+end)

	result = append(result, []byte(`{"state":`)...)
	result = strconv.AppendInt(result, int64(v.State), 10)
	result = append(result, byte('}'))

	return result
}

package complex

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/complex"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	expect = `{"id":1,"first_name":"Senseye","last_name":"Developer","country":"UA","created_at":1567695264,"updated_at":1567695265,"enabled":true}`
)

var (
	shortUser = complex.ShortUser{
		ID:        1,
		FirstName: "Senseye",
		LastName:  "Developer",
		Country:   "UA",
		CreatedAt: 1567695264,
		UpdatedAt: 1567695265,
		Enabled:   true,
	}

	easyShortUser = complex.EasyShortUser{
		ID:        1,
		FirstName: "Senseye",
		LastName:  "Developer",
		Country:   "UA",
		CreatedAt: 1567695264,
		UpdatedAt: 1567695265,
		Enabled:   true,
	}

	perezaShortUser = complex.PerezaShortUser{
		ID:        1,
		FirstName: "Senseye",
		LastName:  "Developer",
		Country:   "UA",
		CreatedAt: 1567695264,
		UpdatedAt: 1567695265,
		Enabled:   true,
	}
)

func TestShortUserEncodingJSON(t *testing.T) {
	actual, err := json.Marshal(shortUser)
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func TestShortUserEasyJSON(t *testing.T) {
	actual, err := easyShortUser.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func TestShortUserPereza(t *testing.T) {
	actual, err := perezaShortUser.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func BenchmarkShortUserEncodingJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(shortUser)
	}
}

func BenchmarkShortUserEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = easyShortUser.MarshalJSON()
	}
}

func BenchmarkShortUserPerezaJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = perezaShortUser.MarshalJSON()
	}
}

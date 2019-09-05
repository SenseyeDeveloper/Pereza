package complex

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/complex"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	expect = `{"id":1,"first_name":"Senseye","last_name":"Developer","country":"UA","created_at":1567695264,"updated_at":1567695265}`
)

func TestShortUserEncodingJSON(t *testing.T) {
	source := complex.ShortUser{
		ID:        1,
		FirstName: "Senseye",
		LastName:  "Developer",
		Country:   "UA",
		CreatedAt: 1567695264,
		UpdatedAt: 1567695265,
	}

	actual, err := json.Marshal(source)
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func TestShortUserEasyJSON(t *testing.T) {
	source := complex.EasyShortUser{
		ID:        1,
		FirstName: "Senseye",
		LastName:  "Developer",
		Country:   "UA",
		CreatedAt: 1567695264,
		UpdatedAt: 1567695265,
	}

	actual, err := source.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func TestShortUserPereza(t *testing.T) {
	source := complex.PerezaShortUser{}

	actual, err := source.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expect), actual)
}

func BenchmarkShortUserEncodingJSON(b *testing.B) {
	source := complex.EasyShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkShortUserEasyJSON(b *testing.B) {
	source := complex.EasyShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkShortUserPerezaJSON(b *testing.B) {
	source := complex.PerezaShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

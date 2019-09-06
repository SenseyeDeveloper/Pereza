package boolean

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/boolean"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	boolStateDataProvider = map[bool]string{
		false: `{"state":false}`,
		true:  `{"state":true}`,
	}
)

func TestBoolStateEncodingJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := boolean.BoolState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestBoolStateEasyJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := boolean.EasyBoolState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestBoolStatePereza(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := boolean.PerezaBoolState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkBoolStateEncodingJSON(b *testing.B) {
	source := boolean.BoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkBoolStateEasyJSON(b *testing.B) {
	source := boolean.EasyBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkBoolStatePerezaJSON(b *testing.B) {
	source := boolean.PerezaBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

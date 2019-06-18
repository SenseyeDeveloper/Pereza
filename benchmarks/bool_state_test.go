package benchmarks

import (
	"encoding/json"
	"github.com/senseyedeveloper/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	boolStateDataProvider = map[bool]string{
		false: fixtures.ExpectBoolStateFalse,
		true:  fixtures.ExpectBoolStateTrue,
	}
)

func TestBoolStateEncodingJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.BoolState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestBoolStateEasyJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.EasyBoolState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestBoolStatePereza(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.PerezaBoolState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkBoolStateEncodingJSON(b *testing.B) {
	source := fixtures.BoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkBoolStateEasyJSON(b *testing.B) {
	source := fixtures.EasyBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkBoolStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

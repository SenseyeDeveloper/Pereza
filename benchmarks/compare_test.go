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

func TestEncodingJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.BoolState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestEasyJSON(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.EasyBoolState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestPereza(t *testing.T) {
	for state, expect := range boolStateDataProvider {
		source := fixtures.PerezaBoolState{
			State: state,
		}

		actual := source.PerezaMarshalJSON()
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkEncodingJSON(b *testing.B) {
	source := fixtures.BoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	source := fixtures.EasyBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkPerezaJSON(b *testing.B) {
	source := fixtures.PerezaBoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_ = source.PerezaMarshalJSON()
	}
}

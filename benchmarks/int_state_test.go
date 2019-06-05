package benchmarks

import (
	"encoding/json"
	"github.com/senseyedeveloper/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var (
	intStateDataProvider = map[int]string{
		-1: fixtures.ExpectIntStateN1,
		0:  fixtures.ExpectIntState0,
		1:  fixtures.ExpectIntState1,
	}
)

func TestIntStateEncodingJSON(t *testing.T) {
	for state, expect := range intStateDataProvider {
		source := fixtures.IntState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestIntStateEasyJSON(t *testing.T) {
	for state, expect := range intStateDataProvider {
		source := fixtures.EasyIntState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestIntStatePereza(t *testing.T) {
	for state, expect := range intStateDataProvider {
		source := fixtures.PerezaIntState{
			State: state,
		}

		actual := source.PerezaMarshalJSON()
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkIntStateEncodingJSON(b *testing.B) {
	source := fixtures.IntState{
		State: math.MinInt32,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkIntStateEasyJSON(b *testing.B) {
	source := fixtures.EasyIntState{
		State: math.MinInt32,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkIntStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaIntState{
		State: math.MinInt32,
	}

	for i := 0; i < b.N; i++ {
		_ = source.PerezaMarshalJSON()
	}
}

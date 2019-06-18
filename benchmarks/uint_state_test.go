package benchmarks

import (
	"encoding/json"
	"github.com/senseyedeveloper/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var (
	uintStateDataProvider = map[uint]string{
		0: fixtures.ExpectUintState0,
		1: fixtures.ExpectUintState1,
	}
)

func TestUintStateEncodingJSON(t *testing.T) {
	for state, expect := range uintStateDataProvider {
		source := fixtures.UintState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestUintStateEasyJSON(t *testing.T) {
	for state, expect := range uintStateDataProvider {
		source := fixtures.EasyUintState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestUintStatePereza(t *testing.T) {
	for state, expect := range uintStateDataProvider {
		source := fixtures.PerezaUintState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkUintStateEncodingJSON(b *testing.B) {
	source := fixtures.UintState{
		State: math.MaxUint32,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkUintStateEasyJSON(b *testing.B) {
	source := fixtures.EasyUintState{
		State: math.MaxUint32,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkUintStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaUintState{
		State: math.MaxUint32,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStub(t *testing.T) {
	expectSize := getStubSize("fixtures", stubTypes)

	assert.Equal(t, expectSize, len(Stub("fixtures", stubTypes)))
}

func BenchmarkStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Stub("fixtures", stubTypes)
	}
}

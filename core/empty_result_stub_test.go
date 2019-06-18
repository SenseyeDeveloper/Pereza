package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyResultStub(t *testing.T) {
	assertEmptyResultStubOneAllocation(t, "EmptyState")
	assertEmptyResultStubOneAllocation(t, "PerezaEmptyState")
}

func BenchmarkEmptyResultStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = EmptyResultStub("PerezaEmptyState")
	}
}

func assertEmptyResultStubOneAllocation(t *testing.T, name string) {
	t.Helper()

	expectSize := getEmptyResultStubSize(name)
	assert.Equal(t, expectSize, len(EmptyResultStub(name)))
}

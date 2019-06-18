package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntResultStub(t *testing.T) {
	assertIntResultStubOneAllocation(t, "BoolState", "SomeState", "some_state")
	assertIntResultStubOneAllocation(t, "PerezaBoolState", "State", "state")
}

func BenchmarkIntResultStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IntResultStub("PerezaBoolState", "State", "state")
	}
}

func assertIntResultStubOneAllocation(t *testing.T, typeName, fieldName, jsonName string) {
	t.Helper()

	expectSize := getIntResultStubSize(typeName, fieldName, jsonName)
	assert.Equal(t, expectSize, len(IntResultStub(typeName, fieldName, jsonName)))
}

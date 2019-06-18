package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringResultStub(t *testing.T) {
	assertStringResultStubOneAllocation(t, "BoolState", "SomeState", "some_state")
	assertStringResultStubOneAllocation(t, "PerezaBoolState", "State", "state")
}

func BenchmarkStringResultStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringResultStub("PerezaBoolState", "State", "state")
	}
}

func assertStringResultStubOneAllocation(t *testing.T, typeName, fieldName, jsonName string) {
	t.Helper()

	expectSize := getStringResultStubSize(typeName, fieldName, jsonName)
	assert.Equal(t, expectSize, len(StringResultStub(typeName, fieldName, jsonName)))
}

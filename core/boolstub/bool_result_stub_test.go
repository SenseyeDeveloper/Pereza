package boolstub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolResultStub(t *testing.T) {
	assertBoolResultStubOneAllocation(t, "BoolState", "SomeState", "some_state")
	assertBoolResultStubOneAllocation(t, "PerezaBoolState", "State", "state")
}

func BenchmarkBoolResultStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BoolResultStub("PerezaBoolState", "State", "state")
	}
}

func assertBoolResultStubOneAllocation(t *testing.T, typeName, fieldName, jsonName string) {
	t.Helper()

	expectSize := getBoolResultStubSize(typeName, fieldName, jsonName)
	assert.Equal(t, expectSize, len(BoolResultStub(typeName, fieldName, jsonName)))
}

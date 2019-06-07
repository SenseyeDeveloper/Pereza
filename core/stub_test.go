package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStub(t *testing.T) {
	assertStubOneAllocation(t, "fixtures", stubTypes)
	assertStubOneAllocation(t, "fixtures", stubLongTypes)
}

func BenchmarkStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Stub("fixtures", stubTypes)
	}
}

func assertStubOneAllocation(t *testing.T, packageName string, types []string) {
	t.Helper()

	expectSize := getStubSize(packageName, types)
	assert.Equal(t, expectSize, len(Stub(packageName, types)))
}

package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunnerStub(t *testing.T) {
	assertRunnerStubOneAllocation(t, "fixtures", "github.com/gopereza/pereza/fixtures", "fixtures", stubTypes)
	assertRunnerStubOneAllocation(t, "fixtures", "github.com/gopereza/pereza/fixtures", "fixtures", stubLongTypes)
}

func BenchmarkRunnerStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RunnerStub("fixtures", "github.com/gopereza/pereza/fixtures", "fixtures", stubTypes)
	}
}

func assertRunnerStubOneAllocation(t *testing.T, name, packagePath, packageName string, types []string) {
	t.Helper()

	expectSize := getRunnerStubSize(name, packagePath, packageName, types)
	assert.Equal(t, expectSize, len(RunnerStub(name, packagePath, packageName, types)))
}

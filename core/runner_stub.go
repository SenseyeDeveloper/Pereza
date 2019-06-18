package core

const (
	runnerHeader = `// +build ignore

// TEMPORARY AUTOGENERATED FILE: pereza bootstapping code to launch
// the actual generator.

package main

import (
	"fmt"
	"github.com/senseyedeveloper/pereza/gen"
	"os"

`

	runnerFooter = `
	if err := g.Run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`
)

func RunnerStub(name, packagePath, packageName string, types []string) []byte {
	content := make([]byte, 0, getRunnerStubSize(name, packagePath, packageName, types))

	content = append(content, runnerHeader...)
	content = append(content, `	pkg "`...)
	content = append(content, packagePath...)
	content = append(content, '"', n, ')', n, n)

	content = append(content, "func main() {\n"...)
	content = append(content, `	g := gen.NewGenerator("`...)
	content = append(content, packagePath...)
	content = append(content, '"', ',', ' ', '"')
	content = append(content, packageName...)
	content = append(content, '"', ',', ' ', '"')
	content = append(content, name...)
	content = append(content, '"', ')')

	for _, t := range types {
		content = append(content, "\n	g.Add(pkg.PerezaJSON_exporter_"...)
		content = append(content, t...)
		content = append(content, "(nil))"...)
	}

	content = append(content, runnerFooter...)

	return content
}

func getRunnerStubSize(name, packagePath, packageName string, types []string) int {
	const (
		headerSize      = len(runnerHeader)
		bodySize        = 59
		typeWrapperSize = 38
		footerSize      = len(runnerFooter)
	)

	return headerSize +
		bodySize +
		len(packagePath)*2 +
		len(name) +
		len(packageName) +
		typeWrapperSize*len(types) +
		stringSliceSize(types) +
		footerSize
}

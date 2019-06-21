package core

import "fmt"

type MultiBoolStubGenerator struct {
	fieldNames  []string
	pattern     string
	buffer      []byte
	returnDepth int
}

func NewMultiBoolStubGenerator(fieldNames, jsonNames []string) *MultiBoolStubGenerator {
	return &MultiBoolStubGenerator{
		fieldNames:  fieldNames,
		pattern:     WrapAsResult(MultiBoolJSONPattern(jsonNames)),
		buffer:      nil, // dynamic allocate
		returnDepth: len(fieldNames) - 1,
	}
}

func (g *MultiBoolStubGenerator) Generate() []byte {
	g.generate(0, FillBooleans(len(g.fieldNames), true))

	return g.buffer
}

func (g *MultiBoolStubGenerator) generate(depth int, states []bool) {
	fieldName := g.fieldNames[depth]

	if depth == g.returnDepth {
		g.append("if v." + fieldName + " {\n")
		g.append(g.f(ReplaceBool(states, depth, true)))
		g.append("}\n")

		g.append(g.f(ReplaceBool(states, depth, false)))

		return
	}

	g.append("if v." + fieldName + " {\n")
	g.generate(depth+1, ReplaceBool(states, depth, true))
	g.append("}\n")

	g.generate(depth+1, ReplaceBool(states, depth, false))
}

func (g *MultiBoolStubGenerator) append(code string) {
	g.buffer = append(g.buffer, code...)
}

func (g *MultiBoolStubGenerator) f(states []bool) string {
	args := make([]interface{}, len(states))

	for i, state := range states {
		args[i] = state
	}

	return fmt.Sprintf(g.pattern, args...)
}

// helper
func MultiBoolJSONPattern(jsonNames []string) string {
	length := len(jsonNames)

	fieldLength := stringSliceSize(jsonNames)

	patternLength := fieldLength + 6*length + 1

	// 1 alloc
	result := make([]byte, 0, patternLength)

	result = append(result, '{')

	skipLastCommaSize := length - 1
	for i := 0; i < skipLastCommaSize; i++ {
		result = append(result, '"')
		result = append(result, jsonNames[i]...)
		result = append(result, '"', ':', '%', 't', ',')
	}

	result = append(result, '"')
	result = append(result, jsonNames[skipLastCommaSize]...)
	result = append(result, '"', ':', '%', 't') // skip comma

	result = append(result, '}')

	// 2 alloc
	return string(result)
}

package core

import "fmt"

type MultiBoolStubGenerator struct {
	fieldNames  []string
	pattern     string
	buffer      []byte
	returnDepth int
}

func NewMultiBoolStubGenerator(fieldNames, jsonNames []string) *MultiBoolStubGenerator {
	pattern := WrapAsResult(MultiBoolJSONPattern(jsonNames))
	length := len(fieldNames)

	return &MultiBoolStubGenerator{
		fieldNames:  fieldNames,
		pattern:     pattern,
		buffer:      make([]byte, 0, len(pattern)*length*length*4), // dynamic allocate
		returnDepth: length - 1,
	}
}

func (g *MultiBoolStubGenerator) Generate() []byte {
	g.generate(0, FillBooleans(len(g.fieldNames), true))

	return g.buffer
}

func (g *MultiBoolStubGenerator) generate(depth int, states []interface{}) {
	fieldName := g.fieldNames[depth]

	if depth == g.returnDepth {
		g.append("if v." + fieldName + " {\n")
		g.append(g.f(states))
		g.append("}\n")

		g.append(g.f(ReplaceBool(states, depth, false)))

		return
	}

	g.append("if v." + fieldName + " {\n")
	g.generate(depth+1, states)
	g.append("}\n")

	g.generate(depth+1, ReplaceBool(states, depth, false))
}

func (g *MultiBoolStubGenerator) append(code string) {
	g.buffer = append(g.buffer, code...)
}

func (g *MultiBoolStubGenerator) f(states []interface{}) string {
	return fmt.Sprintf(g.pattern, states)
}

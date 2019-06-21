package core

type MultiBoolStubGenerator struct {
	fieldNames       []string
	fastConditionMap map[string][]byte
	pattern          *MultiBoolJSONResultGenerator
	buffer           []byte
	returnDepth      int
}

func NewMultiBoolStubGenerator(fieldNames, jsonNames []string) *MultiBoolStubGenerator {
	pattern := NewMultiBoolJSONResultGenerator(jsonNames)

	length := len(fieldNames)

	fastConditionMap := make(map[string][]byte, length)
	for _, fieldName := range fieldNames {
		fastConditionMap[fieldName] = []byte("if v." + fieldName + " {\n")
	}

	return &MultiBoolStubGenerator{
		fieldNames:       fieldNames,
		fastConditionMap: fastConditionMap,
		pattern:          pattern,
		buffer:           make([]byte, 0, pattern.Capacity()*length*length*6), // dynamic allocate
		returnDepth:      length - 1,
	}
}

func (g *MultiBoolStubGenerator) Generate() []byte {
	g.generate(0, FillBooleans(len(g.fieldNames), true))

	return g.buffer
}

func (g *MultiBoolStubGenerator) generate(depth int, states []bool) {
	fieldName := g.fieldNames[depth]

	if depth == g.returnDepth {
		g.append(g.fastConditionMap[fieldName])
		g.append(g.pattern.Generate(ReplaceBool(states, depth, true)))
		g.conditionClose()

		g.append(g.pattern.Generate(ReplaceBool(states, depth, false)))

		return
	}

	g.append(g.fastConditionMap[fieldName])
	g.generate(depth+1, ReplaceBool(states, depth, true))
	g.conditionClose()

	g.generate(depth+1, ReplaceBool(states, depth, false))
}

func (g *MultiBoolStubGenerator) appendString(code string) {
	g.buffer = append(g.buffer, code...)
}

func (g *MultiBoolStubGenerator) append(data []byte) {
	g.buffer = append(g.buffer, data...)
}

func (g *MultiBoolStubGenerator) conditionClose() {
	g.buffer = append(g.buffer, '}', '\n')
}

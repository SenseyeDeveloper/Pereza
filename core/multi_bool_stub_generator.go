package core

type MultiBoolStubGenerator struct {
	fieldNames       []string
	fastConditionMap map[string][]byte
	pattern          *MultiBoolJSONResultGenerator
	buffer           []byte
	returnDepth      int
	capacity         int
}

func NewMultiBoolStubGenerator(fieldNames, jsonNames []string) *MultiBoolStubGenerator {
	pattern := NewMultiBoolJSONResultGenerator(jsonNames)

	length := len(fieldNames)

	capacity := pattern.Capacity() * length * length * length * length * 2

	return &MultiBoolStubGenerator{
		fieldNames:       fieldNames,
		fastConditionMap: FastConditionMap(fieldNames),
		pattern:          pattern,
		buffer:           make([]byte, 0), // dynamic allocate
		returnDepth:      length - 1,
		capacity:         capacity,
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

func FastConditionMap(fieldNames []string) map[string][]byte {
	length := len(fieldNames)

	fastConditionMap := make(map[string][]byte, length)

	const (
		start     = "if v."
		end       = " {\n"
		fixedSize = len(start) + len(end)
	)

	capacity := stringSliceSize(fieldNames) + length*fixedSize
	once := make([]byte, 0, capacity)

	for _, fieldName := range fieldNames {
		current := fixedSize + len(fieldName)

		once = append(once, start...)
		once = append(once, fieldName...)
		once = append(once, end...)

		fastConditionMap[fieldName] = once[:current]

		once = once[current:]
	}

	return fastConditionMap
}

package core

const (
	multiBoolJSONResultHeader    = "return []byte(`{"
	multiBoolJSONResultFooter    = "}`), nil"
	multiBoolJSONResultFixedSize = len(multiBoolJSONResultHeader) + len(multiBoolJSONResultFooter)
)

type MultiBoolJSONResultGenerator struct {
	jsonNames []string
	buffer    []byte
	last      int
}

func NewMultiBoolJSONResultGenerator(jsonNames []string) *MultiBoolJSONResultGenerator {
	length := len(jsonNames)

	const wrap = 8 // len(`"":false`)

	commaCount := length - 1

	return &MultiBoolJSONResultGenerator{
		jsonNames: jsonNames,
		buffer:    make([]byte, 0, multiBoolJSONResultFixedSize+length*wrap+commaCount),
		last:      commaCount,
	}
}

func (g *MultiBoolJSONResultGenerator) Generate(values []bool) []byte {
	result := g.buffer[:0]

	result = append(result, multiBoolJSONResultHeader...)

	for i := 0; i < g.last; i++ {
		result = AppendBool(result, g.jsonNames[i], values[i])

		result = append(result, ',')
	}

	result = AppendBool(result, g.jsonNames[g.last], values[g.last])

	result = append(result, multiBoolJSONResultFooter...)

	return result
}

// helper
func GenerateMultiBoolJSONPattern(jsonNames []string) string {
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

func AppendBool(source []byte, jsonName string, value bool) []byte {
	result := append(source, '"')
	result = append(result, jsonName...)
	result = append(result, '"', ':')

	if value {
		result = append(result, "true"...)
	} else {
		result = append(result, "false"...)
	}

	return result
}

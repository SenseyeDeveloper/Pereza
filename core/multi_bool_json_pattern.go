package core

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

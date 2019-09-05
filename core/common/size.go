package common

func StringSliceSize(slice []string) int {
	result := 0

	for _, s := range slice {
		result += len(s)
	}

	return result
}

func StringStartConst(jsonName string) int {
	// len([]byte(`{"json_name":"`))

	const wrapperSize = 5

	return wrapperSize + len(jsonName)
}

func DigitsSize(i int) int {
	result := 0

	for i > 0 {
		i /= 10

		result += 1
	}

	return result
}

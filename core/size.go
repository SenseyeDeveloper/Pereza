package core

func stringSliceSize(slice []string) int {
	result := 0

	for _, s := range slice {
		result += len(s)
	}

	return result
}

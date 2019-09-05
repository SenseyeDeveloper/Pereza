package common

func AppendJSONFieldNameAsBytes(bytes []byte, jsonName string) []byte {
	for _, l := range jsonName {
		bytes = append(bytes, ',', ' ', '\'', byte(l), '\'')
	}

	return bytes
}

package fixtures

const (
	ExpectEmptyState = `{}`
)

type EmptyState struct {
}

// easyjson:json
type EasyEmptyState struct {
}

type PerezaEmptyState struct {
}

func (v *PerezaEmptyState) PerezaMarshalJSON() []byte {
	return []byte(`{}`)
}

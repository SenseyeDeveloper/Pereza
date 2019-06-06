package parser

type Result struct {
	Name        string
	PackagePath string
	PackageName string
	StructNames []string
	Explicit    bool
}

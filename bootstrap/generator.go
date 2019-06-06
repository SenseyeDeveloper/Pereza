package bootstrap

import "fmt"

type Generator struct {
	settings Settings
}

func NewGenerator(settings Settings) *Generator {
	return &Generator{settings: settings}
}

func (g *Generator) Run() {
	fmt.Printf("Generator::Run with %+v\n", g.settings)
}

package bootstrap

import (
	"github.com/senseyedeveloper/pereza/parser"
	"os"
)

type Generator struct {
	parser   *parser.Parser
	settings Settings
}

func NewGenerator(parser *parser.Parser, settings Settings) *Generator {
	return &Generator{parser: parser, settings: settings}
}

func (g *Generator) Run() error {
	for _, filename := range g.settings.Filenames {
		err := g.generate(filename)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generate(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	err = g.parser.Parse(filename, fileInfo.IsDir())
	if err != nil {
		return err
	}

	return nil
}

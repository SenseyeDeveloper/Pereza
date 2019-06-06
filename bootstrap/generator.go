package bootstrap

import (
	"errors"
	"fmt"
	"github.com/senseyedeveloper/pereza/parser"
	"os"
	"path/filepath"
	"strings"
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

	result, err := g.parser.Parse(filename, fileInfo.IsDir())
	if err != nil {
		return err
	}

	const suffix = "_pereza.go"

	var outName string
	if fileInfo.IsDir() {
		outName = filepath.Join(filename, result.PackageName+suffix)
	} else {
		s := strings.TrimSuffix(filename, ".go")

		if s == filename {
			return errors.New("filename must end in '.go'")
		}

		outName = s + suffix
	}

	fmt.Printf("filename %s parser result %+v to %s\n", filename, result, outName)

	return nil
}

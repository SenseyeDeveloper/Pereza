package bootstrap

import (
	"errors"
	"github.com/gopereza/pereza/core/runnerstub"
	"github.com/gopereza/pereza/parser"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const suffix = "_perezajson.go"

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

	var outName string

	if fileInfo.IsDir() {
		outName = filepath.Join(filename, result.PackageName+suffix)
	} else {
		s := strings.TrimSuffix(filename, ".go")

		// FIXME: up before parse
		if s == filename {
			return errors.New("filename must end in '.go'")
		}

		outName = s + suffix
	}

	//fmt.Printf("source filename %s\n"+
	//	"parser result %+v\n"+
	//	"output filename %s\n", filename, result, outName)

	err = ioutil.WriteFile(outName, runnerstub.Stub(result.PackageName, result.StructNames), 0666)
	if err != nil {
		return err
	}

	bootstrapPath, err := g.writeMain(outName, result.PackagePath, result.PackageName, result.StructNames)
	if err != nil {
		return err
	}
	defer os.Remove(bootstrapPath)

	tempOutput, err := g.runBootstrap(bootstrapPath, outName)
	if err != nil {
		return err
	}

	return os.Rename(tempOutput, outName)
}

func (g *Generator) writeMain(outName, packagePath, packageName string, types []string) (path string, err error) {
	base := filepath.Base(outName)

	bootstrapFilename := filepath.Join(
		filepath.Dir(outName),
		strings.TrimSuffix(base, suffix)+"-easyjson-bootstrap.go",
	)

	content := runnerstub.RunnerStub(base, packagePath, packageName, types)

	return bootstrapFilename, ioutil.WriteFile(bootstrapFilename, content, 0644)
}

func (g *Generator) runBootstrap(bootstrapPath, outName string) (string, error) {
	temp, err := os.Create(outName + ".tmp")
	if err != nil {
		return "", err
	}
	defer temp.Close()

	cmd := exec.Command("go", "run", filepath.Base(bootstrapPath))
	cmd.Stdout = temp
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Dir(bootstrapPath)

	return temp.Name(), cmd.Run()
}

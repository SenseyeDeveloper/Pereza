// +build ignore

// TEMPORARY AUTOGENERATED FILE: pereza bootstapping code to launch
// the actual generator.

package main

import (
	"fmt"
	"github.com/gopereza/pereza/gen"
	"os"

	pkg "github.com/gopereza/pereza/fixtures"
)

func main() {
	g := gen.NewGenerator("github.com/gopereza/pereza/fixtures", "fixtures", "double_bool_state_perezajson.go")
	g.Add(pkg.PerezaJSON_exporter_PerezaDoubleBoolState(nil))
	if err := g.Run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
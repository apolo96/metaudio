package main

import (
	"os"

	"github.com/apolo96/metaudio/cmd/cli/helpers"
)

func main() {
	cmds := bundle()
	parser := helpers.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

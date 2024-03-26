package main

import (
	"net/http"
	"os"

	"github.com/apolo96/metaudio/cmd/cli/commands"
	"github.com/apolo96/metaudio/cmd/cli/helpers"
	"github.com/apolo96/metaudio/internal/interfaces"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		commands.NewGetCommand(client),
		commands.NewListCommand(client),
		commands.NewUploadCommand(client),
	}
	parser := helpers.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

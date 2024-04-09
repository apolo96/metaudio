package main

import (
	"net/http"

	"github.com/apolo96/metaudio/cmd/cli/commands"
	"github.com/apolo96/metaudio/internal/interfaces"
)

var isProPlan bool

func bundle() []interfaces.Command {
	client := &http.Client{}
	cmds := []interfaces.Command{
		commands.NewGetCommand(client),
		commands.NewListCommand(client),
		commands.NewUploadCommand(client),
		commands.NewDeleteCommand(client),
	}
	if isProPlan {
		cmds = append(
			cmds,
			commands.NewSearchCommand(client),
			commands.NewPlayCommand(client),
		)
	}
	return cmds
}

package commands

import (
	"flag"
	"fmt"

	"github.com/apolo96/metaudio/internal/interfaces"
)

type UpdateCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	filename    string
}

func NewUpdateCommand(client interfaces.Client) *UpdateCommand {
	gc := &UpdateCommand{
		flag:        flag.NewFlagSet("list", flag.ContinueOnError),
		description: "Upload audio-file",
		client:      client,
	}
	gc.flag.StringVar(&gc.filename, "filename", "", "full path of filename to be uploaded")
	return gc
}

func (cmd *UpdateCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: metaudio upload -filename <filename>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *UpdateCommand) Run() error {
	fmt.Println("Upload audio file")
	return nil
}

func (cmd *UpdateCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *UpdateCommand) Description() string {
	return cmd.description
}

func (cmd *UpdateCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

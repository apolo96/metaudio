package commands

import (
	"flag"
	"fmt"

	"github.com/apolo96/metaudio/cmd/cli/client"
	"github.com/apolo96/metaudio/internal/interfaces"
)

type GetCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	id          string
}

func NewGetCommand(client interfaces.Client) *GetCommand {
	gc := &GetCommand{
		flag:        flag.NewFlagSet("get", flag.ContinueOnError),
		description: "Get metadata for a particular audio file by id",
		client:      client,
	}
	gc.flag.StringVar(&gc.id, "id", "", "ID of audio-file requested")
	return gc
}

func (cmd *GetCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: metaudio get -id <id>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *GetCommand) Run() error {
	audio, err := client.GetByID(cmd.id, cmd.client)
	if err != nil {
		return err
	}
	if audio == "" {
		fmt.Println("Resource not found, please verify the ID field")
	}
	fmt.Print(audio)
	return nil
}

func (cmd *GetCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *GetCommand) Description() string {
	return cmd.description
}

func (cmd *GetCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

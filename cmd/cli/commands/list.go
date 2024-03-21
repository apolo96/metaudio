package commands

import (
	"flag"
	"fmt"

	"github.com/apolo96/metaudio/internal/interfaces"
)

type ListCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
}

func NewListCommand(client interfaces.Client) *ListCommand {
	return &ListCommand{
		flag:        flag.NewFlagSet("list", flag.ContinueOnError),
		description: "List all metadata",
		client:      client,
	}
}

func (cmd *ListCommand) ParseFlags(flags []string) error {
	return cmd.flag.Parse(flags)
}

func (cmd *ListCommand) Run() error {
	fmt.Println("List all metadata")
	return nil
}

func (cmd *ListCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *ListCommand) Description() string {
	return cmd.description
}

func (cmd *ListCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

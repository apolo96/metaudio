package commands

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/apolo96/metaudio/cmd/cli/client"
	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/apolo96/metaudio/models"
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
		var id string
		prompt := &survey.Input{
			Message: "What is the id of the audio-file?",
		}
		survey.AskOne(prompt, &id)
		if id == "" {
			fmt.Println("usage: metaudio get -id <id>")
			return fmt.Errorf("missing flags")
		}
		flags = append(flags, "-id", id)
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
	model := models.Audio{}
	if err := json.Unmarshal([]byte(audio), &model); err != nil {
		return err
	}
	fmt.Print(model.Table())
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

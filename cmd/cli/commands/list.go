package commands

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/apolo96/metaudio/cmd/cli/config"
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
	req, err := http.NewRequest(http.MethodGet, config.API_LIST_URL, &bytes.Buffer{})
	if err != nil {
		return err
	}
	res, err := cmd.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	/* bs := string(b)
	if bs == ""{
		fmt.Println("Resources not found, please upload an audio file")
	} */
	fmt.Print(string(b))
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

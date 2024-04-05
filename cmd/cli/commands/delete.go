package commands

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
)

type DeleteCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	id          string
}

func NewDeleteCommand(client interfaces.Client) *DeleteCommand {
	dc := &DeleteCommand{
		flag:        flag.NewFlagSet("delete", flag.ContinueOnError),
		description: "Delete audio by ID",
		client:      client,
	}
	dc.flag.StringVar(&dc.id, "id", "", "ID of audio-file")
	return dc
}

func (cmd *DeleteCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: metaudio get -id <id>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *DeleteCommand) Run() error {
	url := strings.Replace(config.API_DELETE_URL, "{id}", cmd.id, 1)
	req, err := http.NewRequest(http.MethodDelete, url, &bytes.Buffer{})
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
	bs := string(b)
	if bs == "" {
		fmt.Println("Resource not found, please verify the ID field")
	}
	fmt.Print(bs)
	return nil
}

func (cmd *DeleteCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *DeleteCommand) Description() string {
	return cmd.description
}

func (cmd *DeleteCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

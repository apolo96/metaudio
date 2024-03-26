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
	url := strings.Replace(config.API_GET_URL, "{id}", cmd.id, 1)
	req, err := http.NewRequest(http.MethodGet, url, &bytes.Buffer{})
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
	if bs == ""{
		fmt.Println("Resource not found, please verify the ID field")
	}
	fmt.Print(bs)
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

package commands

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
)

type SearchCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	value       string
}

func NewSearchCommand(client interfaces.Client) *SearchCommand {
	sc := &SearchCommand{
		flag:        flag.NewFlagSet("search", flag.ContinueOnError),
		description: "List all metadata",
		client:      client,
	}
	sc.flag.StringVar(&sc.value, "value", "", "text value for search")
	return sc
}

func (cmd *SearchCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: metaudio search -value <text>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *SearchCommand) Run() error {
	url := strings.Replace(config.API_SEARCH_URL, "{text}", url.QueryEscape(cmd.value), 1)
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
	fmt.Print(string(b))
	return nil
}

func (cmd *SearchCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *SearchCommand) Description() string {
	return cmd.description
}

func (cmd *SearchCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

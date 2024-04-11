package commands

import (
	"bytes"
	"flag"
	"fmt"
	"net/url"

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
)

type BugCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
}

func NewBugCommand(client interfaces.Client) *BugCommand {
	return &BugCommand{
		flag:        flag.NewFlagSet("bug", flag.ContinueOnError),
		description: "opens the default browser to start a bug report",
		client:      client,
	}
}

func (cmd *BugCommand) ParseFlags(flags []string) error {
	return cmd.flag.Parse(flags)
}

func (cmd *BugCommand) Run() error {
	var buffer bytes.Buffer
	buffer.WriteString("**Audiofile version**\n" + config.CLI_VERSION + "\n\n")
	buffer.WriteString(`**Description**
	A clear description of the bug encountered.
	
	`)
	buffer.WriteString(`**To reproduce**
	Steps to reproduce the bug.
	
	`)
	buffer.WriteString(`**Expected behavior**
	Expected behavior.
	
	`)
	buffer.WriteString(`**Additional details**
	Any other useful data to share.
	
	`)
	issue := buffer.String()
	url := "https://github.com/apolo96/metaudio/issues/new?title=Bug Report&body=" + url.QueryEscape(issue)
	if err := openBrowser(url); err != nil {
		return err
	}
	fmt.Print("Please file a new issue at https://github.com/apolo96/metaudio/issues/new using this template:\n\n")
	fmt.Print(issue)
	return nil
}

func (cmd *BugCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *BugCommand) Description() string {
	return cmd.description
}

func (cmd *BugCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

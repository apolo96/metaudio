package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/apolo96/metaudio/cmd/cli/client"
	"github.com/apolo96/metaudio/models"
	"github.com/pterm/pterm"

	"github.com/apolo96/metaudio/internal/interfaces"
)

type PlayCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	id          string
}

func NewPlayCommand(client interfaces.Client) *PlayCommand {
	gc := &PlayCommand{
		flag:        flag.NewFlagSet("play", flag.ContinueOnError),
		description: "Get metadata for a particular audio file by id",
		client:      client,
	}
	gc.flag.StringVar(&gc.id, "id", "", "ID of audio-file requested")
	return gc
}

func (cmd *PlayCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: metaudio get -id <id>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *PlayCommand) Run() error {
	result, err := client.GetByID(cmd.id, cmd.client)
	if err != nil {
		return err
	}
	if result == "" {
		return fmt.Errorf("resource not found, please verify the ID field")
	}
	audio := &models.Audio{}
	if err := json.Unmarshal([]byte(result), &audio); err != nil {
		return err
	}

	if err := cmd.play(runtime.GOOS, audio.Path); err != nil {
		return err
	}
	return nil
}

func (cmd *PlayCommand) play(os string, audioPath string) error {
	var program string
	switch os {
	case "darwin":
		program = "afplay"
	case "windows":
		program = "cmd /C start"
	case "linux":
		program = "aplay"
	}
	if program == "" {
		return fmt.Errorf("operating system is not support for playing music")
	}
	play := exec.Command(program, audioPath)
	if err := play.Start(); err != nil {
		fmt.Println("ok")
		return err
	}
	spinner := &pterm.SpinnerPrinter{}
	spinner, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	if err := play.Wait(); err != nil {
		return err
	}
	spinner.Stop()
	return nil
}

func (cmd *PlayCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *PlayCommand) Description() string {
	return cmd.description
}

func (cmd *PlayCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

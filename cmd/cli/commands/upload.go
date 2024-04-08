package commands

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/cmd/cli/format"
	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/pterm/pterm"
)

type UploadCommand struct {
	flag        *flag.FlagSet
	description string
	client      interfaces.Client
	filename    string
}

func NewUploadCommand(client interfaces.Client) *UploadCommand {
	gc := &UploadCommand{
		flag:        flag.NewFlagSet("upload", flag.ContinueOnError),
		description: "Upload audio-file",
		client:      client,
	}
	gc.flag.StringVar(&gc.filename, "filename", "", "full path of filename to be uploaded")
	return gc
}

func (cmd *UploadCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		var filename string
		prompt := &survey.Input{
			Message: "What is the filename of the audio to upload",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		}
		survey.AskOne(prompt, &filename)
		if filename == "" {
			fmt.Println("usage: metaudio upload -filename <filename>")
			return fmt.Errorf("missing flags")
		}
		flags = append(flags, "-filename", filename)

	}
	return cmd.flag.Parse(flags)
}

func (cmd *UploadCommand) Run() error {
	/* Progress Bar */
	var progress = &pterm.ProgressbarPrinter{}
	progress, _ = pterm.DefaultProgressbar.WithTotal(2).WithTitle("Initiating upload...").Start()
	time.Sleep(time.Second)
	/* End Progress Bard */

	body := &bytes.Buffer{}
	multiWriter := multipart.NewWriter(body)
	file, err := os.Open(cmd.filename)
	if err != nil {
		return err
	}
	partWrite, err := multiWriter.CreateFormFile("file", filepath.Base(cmd.filename))
	if err != nil {
		return err
	}
	_, err = io.Copy(partWrite, file)
	if err != nil {
		return err
	}
	if err := multiWriter.Close(); err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, config.API_UPDATE_URL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", multiWriter.FormDataContentType())

	/* Progress Bar */
	progress.UpdateTitle("Sending request...")
	time.Sleep(4 * time.Second)
	/* End Progress Bar */

	res, err := cmd.client.Do(req)
	if err != nil {
		return err
	}
	/* Progress Bar */
	pterm.Success.Println("Sending request...")
	progress.Increment()
	/* End Progress Bard */

	defer res.Body.Close()

	/* Progress Bar  */
	progress.UpdateTitle("Process response...")
	time.Sleep(2 * time.Second)
	/* End Progress Bard */

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	/* Progress Bar  */
	pterm.Success.Println("Process response...")
	progress.Increment()
	/* End Progress Bard */

	fmt.Println(format.EmojiCheck, " Successfully uploaded!")
	fmt.Println(format.EmojiCheck, " Audiofile ID: ", string(b))
	return nil
}

func (cmd *UploadCommand) Name() string {
	return cmd.flag.Name()
}

func (cmd *UploadCommand) Description() string {
	return cmd.description
}

func (cmd *UploadCommand) Info() string {
	return fmt.Sprintf("%s : %s", cmd.flag.Name(), cmd.description)
}

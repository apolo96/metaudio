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

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
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
		fmt.Println("usage: metaudio upload -filename <filename>")
		return fmt.Errorf("missing flags")
	}
	return cmd.flag.Parse(flags)
}

func (cmd *UploadCommand) Run() error {
	fmt.Println("Uploading audio file...")
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
	res, err := cmd.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println("Audiofile ID: ", string(b))

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

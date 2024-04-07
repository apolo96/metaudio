package client

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
)

func GetByID(id string, client interfaces.Client) (audio string, err error) {
	url := strings.Replace(config.API_GET_URL, "{id}", id, 1)
	req, err := http.NewRequest(http.MethodGet, url, &bytes.Buffer{})
	if err != nil {
		return audio, err
	}
	res, err := client.Do(req)
	if err != nil {
		return audio, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return audio, err
	}
	return string(data), nil
}

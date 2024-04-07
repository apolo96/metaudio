package client

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/apolo96/metaudio/cmd/cli/config"
	"github.com/apolo96/metaudio/internal/interfaces"
)

/* Models */
type Audio struct {
	Id       string   `json:"Id"`
	Path     string   `json:"Path"`
	Metadata Metadata `json:"Metadata"`
	Status   string   `json:"Status"`
	Error    []string `json:"Error"`
}

type Metadata struct {
	Tags       Tags   `json:"tags"`
	Transcript string `json:"transcript"`
}

type Tags struct {
	Title       string `json:"title"`
	Album       string `json:"album"`
	Artist      string `json:"artist"`
	AlbumArtist string `json:"album_artist"`
	Composer    string `json:"composer"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Lyrics      string `json:"lyrics"`
	Comment     string `json:"comment"`
}

/* API Client */

func GetByID(id string, client interfaces.Client) (s string, err error) {
	url := strings.Replace(config.API_GET_URL, "{id}", id, 1)
	req, err := http.NewRequest(http.MethodGet, url, &bytes.Buffer{})
	if err != nil {
		return s, err
	}
	res, err := client.Do(req)
	if err != nil {
		return s, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return string(b), nil
}

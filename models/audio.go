package models

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/pterm/pterm"
)

type Audio struct {
	Id       string   `json:"Id"`
	Path     string   `json:"Path"`
	Metadata Metadata `json:"Metadata"`
	Status   string   `json:"Status"`
	Error    []string `json:"Error"`
}

func (a *Audio) JSON() (string, error) {
	audioJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(audioJSON), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

/* Table */

var header = []string{
	"ID",
	"Path",
	"Status",
	"Title",
	"Album",
	"Album Artist",
	"Composer",
	"Genre",
	"Artist",
	"Lyrics",
	"Year",
	"Comment",
}

func row(audio Audio) []string {
	return []string{
		audio.Id,
		audio.Path,
		audio.Status,
		audio.Metadata.Tags.Title,
		audio.Metadata.Tags.Album,
		audio.Metadata.Tags.AlbumArtist,
		audio.Metadata.Tags.Composer,
		audio.Metadata.Tags.Genre,
		audio.Metadata.Tags.Artist,
		audio.Metadata.Tags.Lyrics,
		strconv.Itoa(audio.Metadata.Tags.Year),
		audio.Metadata.Tags.Comment,
	}
}

func (a *Audio) Table() (string, error) {
	data := pterm.TableData{header, row(*a)}
	return pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
}

type AudioList []Audio

func (al *AudioList) Table() (string, error) {
	data := pterm.TableData{header}
	for _, audio := range *al {
		data = append(data, row(audio))
	}
	return pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
}

package metadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/apolo96/metaudio/extractors/tags"
	"github.com/apolo96/metaudio/extractors/transcript"
	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/apolo96/metaudio/models"
)

type MetadataService struct {
	Storage interfaces.Storage
}

func NewMetadaService(s interfaces.Storage) *MetadataService {
	return &MetadataService{
		Storage: s,
	}
}

func (ms *MetadataService) Upload(filename string, file io.Reader) (string, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
		return "", err
	}
	defer func() {
		err = os.Remove(filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
		}
		f.Close()
	}()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("error copying file to buffer: ", err)
		return "", err
	}

	id, audioFilePath, err := ms.Storage.Upload(buf.Bytes(), filename)
	audio := &models.Audio{
		Id:   id,
		Path: audioFilePath,
	}
	err = ms.Storage.SaveMetadata(audio)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		return "", err
	}
	audio.Status = "Initiating"
	go ms.extract(audio)
	return id, nil
}

func (ms *MetadataService) extract(audio *models.Audio) {
	var errors []string
	audio.Status = "Complete"
	// tags
	err := tags.Extract(audio)
	if err != nil {
		fmt.Println("error extracting tags metadata: ", err)
		errors = append(errors, err.Error())
	}
	err = ms.Storage.SaveMetadata(audio)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		errors = append(errors, err.Error())
	}

	// transcript
	err = transcript.Extract(audio)
	if err != nil {
		fmt.Println("error extracting transcript metadata: ", err)
		errors = append(errors, err.Error())
	}

	audio.Error = errors
	audio.Status = "Complete"
	err = ms.Storage.SaveMetadata(audio)
	if err != nil {
		fmt.Println("error saving metadata: ", err)
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		fmt.Println("errors occurred extracting metadata: ")
		for i := 0; i < len(errors); i++ {
			fmt.Printf("\terror[%d]: %s\n", i, errors[i])
		}
	} else {
		fmt.Println("successfully extracted and saved audio metadata: ", audio)
	}
}

func (ms *MetadataService) List() (audios string, err error) {
	audioFiles, err := ms.Storage.List()
	if err != nil {
		return audios, err
	}
	jsonData, err := json.Marshal(audioFiles)
	if err != nil {
		return audios, err
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(jsonData), "", "  ")
	audios = prettyJSON.String()
	return audios, err
}

func (ms *MetadataService) Get(id string) (audioJson string, err error) {
	audio, err := ms.Storage.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such file or directory") {
			return audioJson, err
		}
		return audioJson, err
	}
	audioJson, err = audio.JSON()
	if err != nil {
		return audioJson, err
	}
	return audioJson, err
}

func (ms *MetadataService) Delete(id string) error {
	err := ms.Storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MetadataService) Search(text string) (audios string, err error) {
	results, err := ms.Storage.Search(text)
	if err != nil {
		return audios, err
	}
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return audios, err
	}
	audios = string(data)
	return audios, err
}

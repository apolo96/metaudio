package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/apolo96/metaudio/models"

	"github.com/google/uuid"
)

type FlatFile struct {
	Name string
}

func (f FlatFile) GetByID(id string) (*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile", id, "metadata.json")
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(metadataFilePath, os.ModePerm)
	}
	file, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, err
	}
	data := models.Audio{}
	err = json.Unmarshal([]byte(file), &data)
	return &data, err
}

func (f FlatFile) SaveMetadata(audio *models.Audio) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	audioDirPath := filepath.Join(dirname, "audiofile", audio.Id)
	metadataFilePath := filepath.Join(audioDirPath, "metadata.json")
	file, err := os.Create(metadataFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := audio.JSON()
	if err != nil {
		fmt.Println("Err: ", err)
		return err
	}
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func (f FlatFile) Upload(bytes []byte, filename string) (string, string, error) {
	// generate guid
	id := uuid.New()
	// copy file to configured storage path by tag name or id
	dirname, err := os.UserHomeDir()
	if err != nil {
		return id.String(), "", err
	}
	audioDirPath := filepath.Join(dirname, "audiofile", id.String())
	if err := os.MkdirAll(audioDirPath, os.ModePerm); err != nil {
		return id.String(), "", err
	}
	audioFilePath := filepath.Join(audioDirPath, filename)
	err = os.WriteFile(audioFilePath, bytes, 0644)
	if err != nil {
		return id.String(), "", err
	}
	return id.String(), audioFilePath, nil
}

func (f FlatFile) List() ([]*models.Audio, error) {
	dirname, err := os.UserHomeDir()
	println("Storage: listing audios in: " + dirname)
	if err != nil {
		return nil, err
	}
	metadataFilePath := filepath.Join(dirname, "audiofile")
	println("Storage: prepare dir files path: " + metadataFilePath)
	if _, err := os.Stat(metadataFilePath); errors.Is(err, os.ErrNotExist) {
		println("Storage: dir not exits, creating path: " + metadataFilePath)
		if err = os.Mkdir(metadataFilePath, os.ModePerm); err != nil {
			println("Storage: error creating path" + err.Error())
			return nil, err
		}
	}
	println("Storage: reading files")
	files, err := os.ReadDir(metadataFilePath)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	println("Storage: each files audios")
	audioFiles := []*models.Audio{}
	for _, file := range files {
		if file.IsDir() {
			println("Storage: getting file " + file.Name())
			audio, err := f.GetByID(file.Name())
			if err != nil {
				return nil, err
			}
			audioFiles = append(audioFiles, audio)
			println("Storage: prepare audio in list " + audio.Id)
		}
	}
	print("Storage audios length ", len(audioFiles))
	return audioFiles, nil
}

func (f FlatFile) Delete(id string) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(dirname, "audiofile", id)
	if err := os.RemoveAll(filePath); err != nil {
		return err
	}
	return nil
}

func (f FlatFile) Search(text string) ([]*models.Audio, error) {
	dir, err := os.UserHomeDir()
	audios := []*models.Audio{}
	if err != nil {
		return audios, err
	}
	path := filepath.Join(dir, "audiofile")
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "metadata.json" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if strings.Contains(strings.ToLower(string(content)), strings.ToLower(text)) {
				audio := models.Audio{}
				if err = json.Unmarshal(content, &audio); err != nil {
					return err
				}
				audios = append(audios, &audio)
			}
		}
		return nil
	})
	if err != nil {
		return audios, err
	}
	return audios, err
}

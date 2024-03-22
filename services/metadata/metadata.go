package metadata

import (
	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/apolo96/metaudio/storage"
)

type MetadataService struct {
	storage interfaces.Storage
}

func NewMetadaService() *MetadataService {
	return &MetadataService{
		storage: storage.FlatFile{},
	}
}

func (ms *MetadataService) Boot() {

}

func (ms *MetadataService) Upload() {

}

func (ms *MetadataService) List() {

}

func (ms *MetadataService) Get() {

}

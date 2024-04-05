package interfaces

import (
	"io"
)

type Service interface {
	Upload(filename string, file io.Reader) (string, error)
	List() (string, error)
	Get(id string) (string, error)
	Delete(id string) error
	Search(text string) (string, error)
}

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/apolo96/metaudio/services/metadata"
	"github.com/apolo96/metaudio/storage"
)

type MetadataHandler struct {
	Service interfaces.Service
}

func NewMetadataHandler(mux *http.ServeMux) {
	metadataHandler := &MetadataHandler{
		Service: metadata.NewMetadaService(storage.FlatFile{}),
	}
	mux.HandleFunc("POST /upload", metadataHandler.upload)
	mux.HandleFunc("GET /request/{id}", metadataHandler.get)
	mux.HandleFunc("GET /list", metadataHandler.list)
	mux.HandleFunc("DELETE /audio/{id}", metadataHandler.delete)
	mux.HandleFunc("GET /search", metadataHandler.search)
}

func (mh *MetadataHandler) upload(res http.ResponseWriter, req *http.Request) {
	fmt.Println("==== Reading audio file.... ")
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("error creating formfile: ", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	id, err := mh.Service.Upload(handler.Filename, file)
	if err != nil {
		fmt.Println("error upload audio: ", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(res, id)
}

func (mh *MetadataHandler) list(res http.ResponseWriter, req *http.Request) {
	audios, err := mh.Service.List()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(res, audios)
}

func (mh *MetadataHandler) get(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	audio, err := mh.Service.Get(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(res, audio)
}

func (mh *MetadataHandler) delete(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	err := mh.Service.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, fmt.Sprintf("successfully deleted audio with id: %s", id))
}

func (mh *MetadataHandler) search(res http.ResponseWriter, req *http.Request) {
	text := req.URL.Query().Get("q")
	result, err := mh.Service.Search(text)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, result)
}

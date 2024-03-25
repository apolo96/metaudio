package handlers

import (
	"fmt"
	"io"
	"net/http"

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
	if err != nil{
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

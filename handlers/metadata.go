package handlers

import (
	"fmt"
	"net/http"

	"github.com/apolo96/metaudio/internal/interfaces"
	"github.com/apolo96/metaudio/services/metadata"
)

type MetadataHandler struct {
	Service interfaces.Service
}

func NewMetadataHandler(mux *http.ServeMux){
	metadataHandler := &MetadataHandler{
		Service: metadata.NewMetadaService(),
	}
	mux.HandleFunc("/upload", metadataHandler.upload)
	mux.HandleFunc("/request", metadataHandler.get)
	mux.HandleFunc("/list", metadataHandler.list)
}

func (mh MetadataHandler) upload(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res,"upload")
}

func (mh MetadataHandler) list(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res,"list")
}

func (mh MetadataHandler) get(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res,"get")
}

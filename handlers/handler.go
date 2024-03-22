package handlers

import (
	"fmt"
	"net/http"
)

func Listen(port int) error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to MetaAudioAPI")
	})

	NewMetadataHandler(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting api: %v", err)
	}
	return nil
}

package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"log"

	"github.com/gorilla/mux"
)

func respondDial(w http.ResponseWriter, req *http.Request) {
	log.Println("[DEBUG]")
	fmt.Fprintf(w, "OK")
}

func handleSendFile(w http.ResponseWriter, req *http.Request) {
	log.Println("[DEBUG]")
	reader, err := req.MultipartReader()
	if err != nil {
		log.Println("[DEBUG] reader")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// open destination
		_, filename := filepath.Split(part.FileName())
		outfile, err := os.Create("./" + filename)
		defer outfile.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 32K buffer copy
		var written int64
		if written, err = io.Copy(outfile, part); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("File received:" + part.FileName() + "; length:" + strconv.Itoa(int(written))))
	}
}

func Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/dial", respondDial).Methods("POST")
	router.HandleFunc("/sendFile", handleSendFile).Methods("POST")

	server := &http.Server{Addr: ":3000", Handler: router}
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	return server.Serve(listener.(*net.TCPListener))
}

package server

import (
	"fmt"
	"net"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func respondDial(w http.ResponseWriter, r *http.Request) {
	log.Println("[DEBUG]")
	fmt.Fprintf(w, "OK")
}

func Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/dial", respondDial).Methods("POST")

	server := &http.Server{Addr: ":3000", Handler: router}
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	return server.Serve(listener.(*net.TCPListener))
}

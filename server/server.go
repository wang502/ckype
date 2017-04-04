package server

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"log"

	"io/ioutil"

	"bytes"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/wang502/ckype/encryption"
)

var (
	hashedDialMessage = sha256.Sum256([]byte("dial"))
	pemDir, _         = encryption.GetSnippetDir()
)

// Message represents a message sent in ckype
type Message struct {
	Content string `json:"content"`
	Time    int64  `json:"time"`
	From    string `json:"from"`
}

func (m *Message) String() string {
	i, err := strconv.ParseInt(strconv.FormatInt(m.Time, 10), 10, 64)
	if err != nil {
		return ""
	}
	time := time.Unix(i, 0)
	s := fmt.Sprintf("New Message: \nFrom: %s\nTime:%s\n%s\n", m.From, time, m.Content)
	return s
}

// ----------------------------------------------------
//
// Handlers
//
// ----------------------------------------------------

func respondDial(w http.ResponseWriter, req *http.Request) {
	log.Println("[DEBUG]")
	var b bytes.Buffer
	_, err := io.Copy(&b, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("encrypted dial message: %s", b.Bytes())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = encryption.Verify(hashedDialMessage, b.Bytes(), filepath.Join(pemDir, "public_key.pem")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verified, respond with "Verified"
	log.Printf("[Dial] Dial request is Verified")
	fmt.Fprintf(w, "Verified")
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

		// buffer copy
		var written int64
		if written, err = io.Copy(outfile, part); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("File received:" + part.FileName() + "; length:" + strconv.Itoa(int(written))))
	}
}

func handleSendMsg(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("[1]Received a message...\n")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse message
	/*
		msg := &Message{}
		if err := json.Unmarshal(data, msg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/

	fmt.Printf("[2]Decrypting the message...\n")
	//secretContent := msg.Content
	//decryptedContent, err := encryption.RsaDecrypt([]byte(secretContent), fmt.Sprintf("%s/my_private_key.pem", pemDir))

	decryptedContent, err := encryption.RsaDecrypt(data, fmt.Sprintf("%s/my_private_key.pem", pemDir))
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		msg.Content = decryptedContent
		fmt.Printf("%s", msg)
	*/
	fmt.Fprintf(color.Output, "%s\n", color.GreenString(decryptedContent))
	return
}

// Start starts the local http server to ommunicate with other ckype users
func Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/dial", respondDial).Methods("POST")
	router.HandleFunc("/sendFile", handleSendFile).Methods("POST")
	router.HandleFunc("/sendMsg", handleSendMsg).Methods("POST")

	server := &http.Server{Addr: ":3000", Handler: router}
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	return server.Serve(listener.(*net.TCPListener))
}

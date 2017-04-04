package encryption

import (
	"bytes"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func ReadPemFile(dir string) ([]byte, error) {
	pemFile, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer pemFile.Close()

	var b bytes.Buffer
	if _, err = io.Copy(&b, pemFile); err != nil {
		return nil, err
	}

	block, _ := pem.Decode(b.Bytes())
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	if block.Type != "PUBLIC KEY" && block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	return block.Bytes, nil
}

func GetSnippetDir() (dir string, err error) {
	if runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "ckype")
		}
		dir = filepath.Join(dir, "pet")
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config", "ckype")
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", fmt.Errorf("cannot create directory: %v", err)
	}
	return dir, nil
}

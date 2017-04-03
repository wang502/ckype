package encryption

import (
	"bytes"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

func readPemFile(dir string) ([]byte, error) {
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

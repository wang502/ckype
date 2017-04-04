package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getIP() (string, error) {
	bashCmd := exec.Command("curl", "ipecho.net/plain")
	var out bytes.Buffer
	bashCmd.Stdout = &out
	err := bashCmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func getSnippetDir() (dir string, err error) {
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

package cmd

import (
	"bytes"
	"os/exec"
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

package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show your public IP address",
	Long:  `Show your public IP address`,
	RunE:  ip,
}

func ip(cmd *cobra.Command, args []string) error {
	bashCmd := exec.Command("curl", "ipecho.net/plain")
	var out bytes.Buffer
	bashCmd.Stdout = &out
	err := bashCmd.Run()
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output, "%s : %s\n", color.GreenString("Your public IP address"), out.String())
	return nil
}

func init() {
	baseCmd.AddCommand(ipCmd)
}

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load and copy your friend's public key pem file into ckype folder for future use",
	Long:  `Load your friend's public key pem file into ckype folder for future use`,
	RunE:  load,
}

func load(cmd *cobra.Command, args []string) error {
	file := args[0]
	command := fmt.Sprintf("cp %s %s/public_key.pem", file, pemDir)
	_, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output, "Your friend's %s %s %s\n", color.GreenString(file), "is loaded into directory:", color.GreenString(pemDir))
	return nil
}

func init() {
	baseCmd.AddCommand(loadCmd)
}

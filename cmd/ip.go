package cmd

import (
	"fmt"

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
	ipAddr, err := getIP()
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output, "%s : %s\n", color.GreenString("Your public IP address"), ipAddr)
	return nil
}

func init() {
	baseCmd.AddCommand(ipCmd)
}

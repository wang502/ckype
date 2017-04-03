package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var genPriCmd = &cobra.Command{
	Use:   "genPrivateKey",
	Short: "Generate RSA keypair with a 2048 bit private key",
	Long:  `Generate RSA keypair with a 2048 bit private key`,
	RunE:  genPrivateKey,
}

var genPubCmd = &cobra.Command{
	Use:   "genPublicKey",
	Short: "Generate RSA public key based on your private key",
	Long:  `Generate RSA public key based on your private key`,
	RunE:  genPublicKey,
}

func genPrivateKey(cmd *cobra.Command, args []string) error {
	cmdStr := "openssl genrsa -out private_key.pem"
	_, err := exec.Command("sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output, "%s \"%s\"\n", color.GreenString("An RSA private key is generated in"), "private_key.pem")
	return nil
}

func genPublicKey(cmd *cobra.Command, args []string) error {
	cmdStr := "openssl rsa -pubout -in private_key.pem -out public_key.pem"
	_, err := exec.Command("sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output, "%s \"%s\" based on your \"private_key.pem\"\n", color.GreenString("An RSA public key is generated in"), "public_key.pem")
	return nil
}

func init() {
	baseCmd.AddCommand(genPriCmd)
	baseCmd.AddCommand(genPubCmd)
}

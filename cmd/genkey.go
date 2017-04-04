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
	fmt.Println("Generating private key....")

	destDir := fmt.Sprintf("%s/my_private_key.pem", pemDir)
	command := fmt.Sprintf("openssl genrsa -out %s 2048", destDir)
	_, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output, "%s %s\n", "An RSA private key is generated in", color.GreenString(destDir))
	fmt.Println("...")
	fmt.Println("[Done]")
	return nil
}

func genPublicKey(cmd *cobra.Command, args []string) error {
	fmt.Println("Generating public key....")

	// path of existing private ket pem file
	privateDir := fmt.Sprintf("%s/my_private_key.pem", pemDir)

	// destination directory of public key pem file
	pubDir := fmt.Sprintf("%s/my_public_key.pem", pemDir)

	command := fmt.Sprintf("openssl rsa -pubout -in %s -out %s", privateDir, pubDir)
	_, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return err
	}

	fmt.Fprintf(color.Output, "%s %s based on your private key in %s\n", "An RSA public key is generated in", color.GreenString(pubDir), color.GreenString(privateDir))
	fmt.Println("...")
	fmt.Println("[Done]")
	return nil
}

func init() {
	baseCmd.AddCommand(genPriCmd)
	baseCmd.AddCommand(genPubCmd)
}

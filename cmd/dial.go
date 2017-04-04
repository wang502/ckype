package cmd

import (
	"fmt"
	"io/ioutil"

	"bytes"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wang502/ckype/encryption"
)

var dialCmd = &cobra.Command{
	Use:   "dial",
	Short: "Check whether the other user is online",
	Long:  `Check whether the other user is online`,
	RunE:  dial,
}

func dial(cmd *cobra.Command, args []string) error {
	signature, err := encryption.Sign("dial", fmt.Sprintf("%s/%s", pemDir, "my_private_key.pem"))
	if err != nil {
		return err
	}

	var b bytes.Buffer
	//b.Write([]byte("dial"))
	b.Write(signature)
	httpResp, err := httpClient.Post(fmt.Sprintf("http://%s:3000/dial", args[0]), "ckype", &b)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	data, _ := ioutil.ReadAll(httpResp.Body)
	fmt.Fprintf(color.Output, "%s %s : %s\n", color.GreenString("Dial response from"), args[0], data)
	if string(data) != "Verified" {
		fmt.Fprintf(color.Output, "User %s rejected your dialing\n", args[0])
		return nil
	}

	fmt.Fprintf(color.Output, "%s %s\n", color.YellowString("You can now send messages and files to"), args[0])
	return nil
}

func init() {
	baseCmd.AddCommand(dialCmd)
}

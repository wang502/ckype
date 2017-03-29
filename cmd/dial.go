package cmd

import (
	"fmt"
	"io/ioutil"

	"bytes"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dialCmd = &cobra.Command{
	Use:   "dial",
	Short: "Check whether the other user is online",
	Long:  `Check whether the other user is online`,
	RunE:  dial,
}

func dial(cmd *cobra.Command, args []string) error {
	var b bytes.Buffer
	b.Write([]byte("dial"))
	httpResp, err := httpClient.Post("http://"+args[0]+"/dial", "ckype", &b)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	data, _ := ioutil.ReadAll(httpResp.Body)
	fmt.Fprintf(color.Output, "%s %s : %s\n", color.GreenString("Dial response from"), args[0], data)
	fmt.Fprintf(color.Output, "%s %s\n", color.YellowString("You can now send messages and files to"), args[0])

	return nil
}

func init() {
	baseCmd.AddCommand(dialCmd)
}

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var sendFileCmd = &cobra.Command{
	Use:   "send_file",
	Short: "Send file to your mate",
	Long:  `Send file to your mate`,
	RunE:  sendFile,
}

func sendFile(cmd *cobra.Command, args []string) error {
	ip, filePath := args[0], args[1]

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fw, err := w.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	if _, err = io.Copy(fw, f); err != nil {
		return err
	}
	w.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/sendFile", ip), &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	respByte, _ := ioutil.ReadAll(res.Body)
	log.Println(string(respByte))
	return err
}

func init() {
	baseCmd.AddCommand(sendFileCmd)
}

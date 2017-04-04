package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"io/ioutil"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wang502/ckype/encryption"
	"github.com/wang502/ckype/server"
)

// ----------------------------------------------------
//
// Send File
//
// ----------------------------------------------------

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

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:3000/sendFile", ip), &b)
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

// ----------------------------------------------------
//
// Send Message
//
// ----------------------------------------------------

var sendMsgCmd = &cobra.Command{
	Use:   "sendmsg",
	Short: "Send message to your mate",
	Long:  `Send message to your mate`,
	RunE:  sendMsg,
}

func sendMsg(cmd *cobra.Command, args []string) error {
	to, message := args[0], args[1]

	// encrypt message
	/*
		ciphertext, err := encryption.RsaEncrypt(message, fmt.Sprintf("%s/public_key.pem", pemDir))
		if err != nil {
			return err
		}
	*/

	fmt.Fprintf(color.Output, "%s: %s\n", color.GreenString("Message"), message)
	from, err := getIP()
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output, "%s: %s\n", color.GreenString("Sending message from"), from)
	fmt.Fprintf(color.Output, color.GreenString("Sending message...\n"))

	// prepare message
	msg := &server.Message{
		Content: message,
		Time:    time.Now().Unix(),
		From:    from,
	}
	/*
		data, err := json.Marshal(msg)
		if err != nil {
			return err
		}
	*/
	ciphertext, err := encryption.RsaEncrypt(msg.String(), fmt.Sprintf("%s/public_key.pem", pemDir))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	//buf.Write(data)
	buf.Write(ciphertext)
	resp, err := httpClient.Post(fmt.Sprintf("http://%s:3000/sendMsg", to), "ckype", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", resp.Status)
		return err
	}

	fmt.Fprintf(color.Output, color.GreenString("Message is received successfully!\n"))
	return nil
}

func init() {
	baseCmd.AddCommand(sendFileCmd)
	baseCmd.AddCommand(sendMsgCmd)
}

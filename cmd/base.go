package cmd

import (
	"fmt"
	"net/http"
	"time"

	"os"

	"sync"

	"github.com/spf13/cobra"
)

var baseCmd = &cobra.Command{
	Use:   "ckype",
	Short: "A command line tool for P2P chatting and file sharing",
	Long:  `ckype - A command line tool for P2P chat and file sharing`,
}

var routineGroup sync.WaitGroup
var httpClient = http.Client{Timeout: time.Second}

// Execute is wrapper of base command
func Execute() {
	if err := baseCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

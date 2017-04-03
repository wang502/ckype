package cmd

import (
	"os/exec"

	"log"

	"github.com/spf13/cobra"
	"github.com/wang502/ckype/server"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Get online",
	Long:  `Get online`,
	RunE:  start,
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start daemon server",
	Long:  `Start daemon server`,
	RunE:  daemon,
}

func start(cmd *cobra.Command, args []string) error {
	var err error

	routineGroup.Add(1)
	go func() {
		defer routineGroup.Done()
		err = server.Start()
	}()

	routineGroup.Wait()
	return err
}

func daemon(cmd *cobra.Command, args []string) error {
	strCmd := "./ckype start &"
	out, err := exec.Command("sh", "-c", strCmd).Output()
	if err != nil {
		return err
	}

	log.Println(string(out))
	return nil
}

func init() {
	baseCmd.AddCommand(startCmd)
	baseCmd.AddCommand(daemonCmd)
}

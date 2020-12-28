package cmd

import (
	"awesomeProject/DoMyDocker/container"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the env.",

	Run: func(cmd *cobra.Command, args []string) {
		//command := args[0]
		//tty :=
		//	Run(tty, command)
	},
}

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}

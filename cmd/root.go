package cmd

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "authorization-svc",
	Short: "Authorization service for check userRoles and permission",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use serve to start a server")
		fmt.Println("Use -h to see the list of command")
	},
}

//Run run command
func Run() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	serverCommand.PersistentFlags().StringVarP(&configURL, "config", "c", "config/files", "Config URL i.e. config/files")
	rootCommand.AddCommand(serverCommand)

	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

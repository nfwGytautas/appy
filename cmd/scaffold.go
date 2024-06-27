package appy_cmd

import (
	"fmt"
	"os"

	appy_pkg "github.com/nfwGytautas/appy/pkg"
	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffold the application from the appy.yaml file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := appy_pkg.Scaffold(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}

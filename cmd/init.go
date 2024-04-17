package appy_cmd

import (
	"fmt"
	"os"

	appy_pkg "github.com/nfwGytautas/appy/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init name language",
	Short: "Initialize appy project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := appy_pkg.Init(args[0], args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

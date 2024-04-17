package appy_cmd

import (
	"fmt"
	"os"

	appy_pkg "github.com/nfwGytautas/appy/pkg"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the appy configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := appy_pkg.Verify(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Configuration is valid")
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}

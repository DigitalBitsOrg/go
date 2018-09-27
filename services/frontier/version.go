package main

import (
	"fmt"

	apkg "github.com/digitalbitsorg/go/support/app"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print frontier version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(apkg.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

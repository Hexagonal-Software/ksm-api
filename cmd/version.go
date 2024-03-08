package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hexagonal.software/ksm-api/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the KSM API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("KSM API", version.Version)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

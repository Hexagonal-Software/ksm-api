package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"hexagonal.software/ksm-api/internal/app"
	"hexagonal.software/ksm-api/internal/config"
	"hexagonal.software/ksm-api/internal/logging"
	"hexagonal.software/ksm-api/internal/version"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApplication(config.Conf)
		logging.Log.Info("Version: ", version.Version)

		if err := app.Bootstrap(); err != nil {
			logging.Log.Fatal(err)
			os.Exit(1)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Println("Gracefully shutting down...")
			app.Shutdown()
		}()

		if err := app.RunServer(); err != nil {
			logging.Log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

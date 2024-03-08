package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"hexagonal.software/ksm-api/internal/app"
	"hexagonal.software/ksm-api/internal/config"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApplication(config.Conf)

		if err := app.Bootstrap(); err != nil {
			log.Fatal(err)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Println("Gracefully shutting down...")
			app.Shutdown()
		}()

		if err := app.RunServer(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

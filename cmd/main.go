package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"hexagonal.software/ksm-api/internal/config"
	"hexagonal.software/ksm-api/internal/logging"
)

var (
	cfgFile  string
	LogLevel int
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "ksm-api",
	Short: "KSM-api adds an API to the KSM app",
	Long: `KSM-api adds an API to the KSM app. It is a RESTful API that allows users to retrieve
		secrets from Keeper Vault using an REST API.
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := viper.Unmarshal(&config.Conf)
		if err != nil {
			log.Fatal("unable to decode into struct, %v", err)
		}
		logging.NewLogger(LogLevel)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetEnvPrefix("KSM_API")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__", "-", "_"))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "ksmapi.yaml", "Use specific config file")
	rootCmd.PersistentFlags().IntVar(&LogLevel, "log-level", int(log.LevelError), "Use debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("ksmapi")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.BindEnv("KV.KsmConfig", "KSM_CONFIG")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// message := fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed())
		// log.Info(message)
		return
	}
}

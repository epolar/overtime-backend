package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"orderStatistics/runtime/log"
	"os"
)

var configPath string

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config-path", "./config", "config file path")
	rootCmd.PersistentFlags().Bool("enable-db-log", false, "print db log")

	cobra.OnInitialize(initViper)

	_ = viper.BindPFlag("config-path", rootCmd.PersistentFlags().Lookup("config-path"))
	_ = viper.BindPFlag("enable-db-log", rootCmd.PersistentFlags().Lookup("enable-db-log"))

}

func initViper() {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Log.Errorf("init viper failure: %s", err)
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "morecoin-sync",
	Short: "morecoin online data monitor and async",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please select the corresponding subcommand")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

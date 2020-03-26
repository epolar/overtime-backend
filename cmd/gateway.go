package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"orderStatistics/gateway/user"
	"orderStatistics/repository"
	"os"
)

var addr string

func init() {
	gatewayCmd.PersistentFlags().StringVar(&addr, "addr", ":8001", "listen addr")

	gatewayCmd.AddCommand(userCmd)
	rootCmd.AddCommand(gatewayCmd)
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run gateway",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please specify the gateway name")
	},
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Run user gateway",
	Run: func(cmd *cobra.Command, args []string) {
		db := repository.DB() // pre init db conn
		defer func() {
			_ = db.Close()
		}()

		if err := user.Run(addr); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

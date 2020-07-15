package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "managing jqueen daemon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daemon called")
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

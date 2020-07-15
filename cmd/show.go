package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detail information of job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// showCmd.PersistentFlags().String("foo", "", "A help for foo")
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

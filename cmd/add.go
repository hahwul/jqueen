package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)


	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

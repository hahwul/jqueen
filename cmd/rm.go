package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

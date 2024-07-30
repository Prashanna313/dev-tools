package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Start debugger",
	Long:  `Start a debugger with appropriate configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting debugger...")
		// Add debugging logic here
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testType string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Long:  `Run unit, integration, and end-to-end tests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running %s tests...\n", testType)
		// Add test running logic here
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&testType, "type", "t", "unit", "Type of tests to run (unit, integration, e2e)")
}

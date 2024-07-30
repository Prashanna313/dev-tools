package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var language string
var framework string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate boilerplate code",
	Long:  `Generate boilerplate code for different languages and frameworks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating boilerplate code for %s using %s framework...\n", language, framework)
		// Add code generation logic here
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&language, "language", "l", "", "Programming language (e.g., Go, Python)")
	generateCmd.Flags().StringVarP(&framework, "framework", "f", "", "Framework to use")
}

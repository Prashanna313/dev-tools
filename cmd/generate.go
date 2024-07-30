package cmd

import (
	"fmt"

	"github.com/prashanna313/dev-tools/internal/python"
	"github.com/spf13/cobra"
)

var projectName string
var templateType string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Python boilerplate code",
	Long:  `Generate Python boilerplate code based on selected templates such as Flask, Django, or CLI applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" {
			fmt.Println("Please specify a project name using the -n or --name flag.")
			return
		}

		if !python.IsValidTemplate(templateType) {
			fmt.Printf("Unknown template type: %s. Available templates are:\n", templateType)
			for key, value := range python.AvailableTemplates() {
				fmt.Printf("- %s: %s\n", key, value)
			}
			return
		}

		err := python.GenerateProject(projectName, templateType)
		if err != nil {
			fmt.Printf("Error generating project: %v\n", err)
		} else {
			fmt.Printf("Successfully generated %s project '%s'.\n", python.AvailableTemplates()[templateType], projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
	generateCmd.Flags().StringVarP(&templateType, "template", "t", "", "Type of template to generate (flask, django, cli)")
}

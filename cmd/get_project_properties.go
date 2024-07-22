/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// getProjectPropertiesCmd represents the getProjectProperties command
var get_project_propertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "Get project properties",
	Long:  `Retrieve project's properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects.GetProjectProperties(projectId)
	},
}

func init() {
	get_projectCmd.AddCommand(get_project_propertiesCmd)

	get_project_propertiesCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "Azure DevOps project's id")
}

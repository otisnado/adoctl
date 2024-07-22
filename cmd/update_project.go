/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// updateProjectCmd represents the updateProject command
var update_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Update an Azure DevOps organization project",
	Long:  `This commandd help you to queue a project to be updated into your Azure DevOps organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects.UpdateProject(projectId, projectName, projectDescription, projectVisibility)
	},
}

func init() {
	updateCmd.AddCommand(update_projectCmd)

	update_projectCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id to queue for updating")
	update_projectCmd.Flags().StringVarP(&projectName, "new-project-name", "n", "", "new project name")
	update_projectCmd.Flags().StringVarP(&projectDescription, "new-project-description", "d", "", "new project description")
	update_projectCmd.Flags().StringVarP(&projectVisibility, "new-project-visibility", "v", "", "new project visibility")

	update_projectCmd.MarkFlagRequired("project-id")
}

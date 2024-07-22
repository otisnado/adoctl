/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// deleteProjectCmd represents the deleteProject command
var delete_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Delete an Azure DevOps organization project",
	Long:  `This command help you to queue a project to be deleted into your Azure DevOps organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects.DeleteProject(projectId)
	},
}

func init() {
	deleteCmd.AddCommand(delete_projectCmd)

	delete_projectCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id to queue for deleting")
	delete_projectCmd.MarkFlagRequired("project-id")
}

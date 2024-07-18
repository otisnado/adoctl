/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// Variables to create project
var projectName string
var projectDescription string
var projectSourceControlSystem string
var projectProcessTemplate string
var projectVisibility string

// createProjectCmd represents the createProject command
var create_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create an Azure DevOps organization project",
	Long:  `This command help you to create a project into your Azure DevOps organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := projects.CreateProject(projectName, projectDescription, projectSourceControlSystem, projectProcessTemplate, projectVisibility)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	createCmd.AddCommand(create_projectCmd)

	create_projectCmd.Flags().StringVarP(&projectName, "name", "n", "", "project's name to be created")
	create_projectCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "project's description to be created")
	create_projectCmd.Flags().StringVar(&projectSourceControlSystem, "source-control-system", "Git", "source control system to be used in project")
	create_projectCmd.Flags().StringVarP(&projectProcessTemplate, "process", "p", "SCRUM", "process template to be used in project")
	create_projectCmd.Flags().StringVarP(&projectVisibility, "project-visibility", "v", "Private", "project visibility")

	create_projectCmd.MarkFlagsRequiredTogether("name", "description", "source-control-system", "process", "project-visibility")
}

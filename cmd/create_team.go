/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/teams"
	"github.com/spf13/cobra"
)

var teamName string
var teamDescription string

// createTeamCmd represents the createTeam command
var create_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Create an Azure DevOps team",
	Long:  `This command help you to create a team into a Azure DevOps project.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := teams.CreateTeam(teamName, teamDescription, projectId)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	createCmd.AddCommand(create_teamCmd)

	create_teamCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "Project's id where team will be associated")
	create_teamCmd.Flags().StringVarP(&teamName, "name", "n", "", "Team's name")
	create_teamCmd.Flags().StringVarP(&teamDescription, "description", "d", "", "Team's description")
}

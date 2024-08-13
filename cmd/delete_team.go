/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/teams"
	"github.com/spf13/cobra"
)

// deleteTeamCmd represents the deleteTeam command
var delete_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Delete an Azure DevOps' team",
	Long:  `This command help you to delete a Azure DevOps' team.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := teams.DeleteTeam(teamId, projectId)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	deleteCmd.AddCommand(delete_teamCmd)

	delete_teamCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id where team is")
	delete_teamCmd.Flags().StringVarP(&teamId, "team-id", "t", "", "team id to be deleted")

	delete_teamCmd.MarkFlagRequired("project-id")
	delete_teamCmd.MarkFlagRequired("team-id")
}

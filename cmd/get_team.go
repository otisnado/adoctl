/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/teams"
	"github.com/spf13/cobra"
)

var teamId string

// getTeamCmd represents the getTeam command
var get_teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get a specific Azure DevOps team",
	Long:  `Retrieve an Azure DevOps team using project id and team id.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := teams.GetTeamById(projectId, teamId)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	getCmd.AddCommand(get_teamCmd)

	get_teamCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id")
	get_teamCmd.Flags().StringVarP(&teamId, "team-id", "t", "", "team id")

	get_teamCmd.MarkFlagRequired("project-id")
	get_teamCmd.MarkFlagRequired("team-id")
}

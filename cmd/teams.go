/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/teams"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Get teams associated to a project",
	Long:  `Retrieve all teams to the given project id.`,
	Run: func(cmd *cobra.Command, args []string) {
		teams.GetTeams(projectId)
	},
}

func init() {
	getCmd.AddCommand(teamsCmd)

	teamsCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id")
	teamsCmd.MarkFlagRequired("project-id")
}

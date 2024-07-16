/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get projects in Azure DevOps organization",
	Long:  `List all Azure DevOps organization's projects than user's token provided has access.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := projects.GetProjects()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	getCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

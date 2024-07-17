/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

var projectId string
var shoCapabilities bool
var showHistory bool

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get an specific organization's project",
	Long:  `Retrieve an organization's project by id`,
	Run: func(cmd *cobra.Command, args []string) {
		err := projects.GetProjectById(&projectId, &shoCapabilities, &showHistory)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	getCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "Azure DevOps project's id")
	projectCmd.Flags().BoolVar(&shoCapabilities, "show-capabilities", false, "show project's capabilites. Defaults true")
	projectCmd.Flags().BoolVar(&showHistory, "search-history", false, "search project's history name. Default false")
	projectCmd.MarkFlagRequired("project-id")
}

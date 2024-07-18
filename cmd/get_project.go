/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/api/projects"
	"github.com/spf13/cobra"
)

// Variables to get project info
var shoCapabilities bool
var showHistory bool

// get_projectCmd represents the project command
var get_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Get an specific Azure DevOps organization's project",
	Long:  `Retrieve an Azure DevOps organization's project by id`,
	Run: func(cmd *cobra.Command, args []string) {
		err := projects.GetProjectById(&projectId, &shoCapabilities, &showHistory)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	getCmd.AddCommand(get_projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// get_projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// get_projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	get_projectCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "Azure DevOps project's id")
	get_projectCmd.Flags().BoolVar(&shoCapabilities, "show-capabilities", false, "show project's capabilites. Defaults true")
	get_projectCmd.Flags().BoolVar(&showHistory, "search-history", false, "search project's history name. Default false")
}

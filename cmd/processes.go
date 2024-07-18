/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/processes"
	"github.com/spf13/cobra"
)

// processesCmd represents the processes command
var processesCmd = &cobra.Command{
	Use:   "processes",
	Short: "Get processes in Azure DevOps organization",
	Long:  `List all Azure DevOps organization's processes that user's token provided has access.`,
	Run: func(cmd *cobra.Command, args []string) {
		processes.GetAllProcesses()
	},
}

func init() {
	getCmd.AddCommand(processesCmd)

	processesCmd.Flags().StringVarP(&projectId, "project-id", "i", "", "project id")
}

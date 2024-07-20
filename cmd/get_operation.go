/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/operations"
	"github.com/spf13/cobra"
)

var operationId string

// operationCmd represents the operation command
var get_operationCmd = &cobra.Command{
	Use:   "operation",
	Short: "Get a specific Azure DevOps operation",
	Long:  `Retrieve an Azure DevOps organization's operation by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		operations.GetOperation(operationId)
	},
}

func init() {
	getCmd.AddCommand(get_operationCmd)

	get_operationCmd.Flags().StringVarP(&operationId, "operation-id", "i", "", "operation id to retrieve")
	get_operationCmd.MarkFlagRequired("operation-id")
}

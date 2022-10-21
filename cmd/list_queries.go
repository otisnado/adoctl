/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/otisnado/adoctl/api/query"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// queriesCmd represents the queries command
var list_queriesCmd = &cobra.Command{
	Use:   "queries",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId := viper.GetString("listQueriesProjectId")
		depth := 2
		query.ListQueries(projectId, &depth)
	},
}

func init() {
	listCmd.AddCommand(list_queriesCmd)

	list_queriesCmd.Flags().StringP("project-id", "i", "", "Project Id where queries are")
	list_queriesCmd.MarkFlagRequired("project-id")
	viper.BindPFlag("listQueriesProjectId", list_queriesCmd.Flags().Lookup("project-id"))
}

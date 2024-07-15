/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/otisnado/adoctl/utils"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth configuration that adoctl require to work",
	Long:  `Add Azure DevOps organization and access token to authenticate to manage.`,
	Run: func(cmd *cobra.Command, args []string) {

		isInlineConfig, _ := cmd.Flags().GetBool("inline")
		if !isInlineConfig {

			err := utils.AuthConfig()
			if err != nil {
				log.Fatalln(err)
			}

		} else {

			organizationUrl, _ := cmd.Flags().GetString("org-url")
			organizationToken, _ := cmd.Flags().GetString("token")
			err := utils.InlineAuthConfig(organizationUrl, organizationToken)
			if err != nil {
				log.Fatalln(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().BoolP("inline", "i", false, "option to set auth config in line terminal.")
	authCmd.Flags().StringP("org-url", "o", "", "organization URL.")
	authCmd.Flags().StringP("token", "t", "", "organization token.")
	authCmd.MarkFlagsRequiredTogether("inline", "org-url", "token")
}

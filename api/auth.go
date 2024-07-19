package api

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/utils"
	"github.com/spf13/viper"
)

func CoreClient() core.Client {
	organizationUrl := viper.GetString("organizationurl")
	organizationToken, err := utils.DecodeStringBase64(viper.GetString("organizationtoken"))
	if err != nil {
		log.Fatalln(err)
	}

	connection := azuredevops.NewPatConnection(organizationUrl, organizationToken)

	ctx := context.Background()

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatalln(err)
	}
	return coreClient
}

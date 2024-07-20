package api

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/otisnado/adoctl/utils"
	"github.com/spf13/viper"
)

func AdoConnection() (*azuredevops.Connection, error) {
	organizationUrl := viper.GetString("organizationurl")
	organizationToken, err := utils.DecodeStringBase64(viper.GetString("organizationtoken"))
	if err != nil {
		return nil, err
	}

	connection := azuredevops.NewPatConnection(organizationUrl, organizationToken)
	return connection, nil
}

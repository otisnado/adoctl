package clients

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/operations"
	"github.com/otisnado/adoctl/api"
)

var ctx context.Context

func CoreClient() core.Client {
	ctx := context.Background()
	connection, err := api.AdoConnection()
	if err != nil {
		log.Fatalln(err)
	}

	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatalln(err)
	}
	return coreClient
}

func OperationsClient() (operations.Client, error) {
	adoClient, err := api.AdoConnection()
	if err != nil {
		return nil, err
	}

	operationClient := operations.NewClient(ctx, adoClient)
	return operationClient, nil
}

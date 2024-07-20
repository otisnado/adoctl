package operations

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/operations"
	"github.com/otisnado/adoctl/api/clients"
	"github.com/rodaine/table"
)

func GetOperation(operationIdString string) {
	operationUuid, err := uuid.Parse(operationIdString)
	if err != nil {
		log.Fatalln(err)
	}
	var ctx context.Context
	operationArgs := operations.GetOperationArgs{
		OperationId: &operationUuid,
	}

	operationClient, err := clients.OperationsClient()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := operationClient.GetOperation(ctx, operationArgs)
	if err != nil {
		log.Fatalln(err)
	}

	t := table.New("ID", "Status", "URL")
	t.AddRow(res.Id, *res.Status, *res.Url)
	t.Print()

}

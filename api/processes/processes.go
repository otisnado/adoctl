package processes

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/api/clients"
	"github.com/rodaine/table"
)

var ctx context.Context

func GetAllProcesses() {
	var processesArgs core.GetProcessesArgs
	processes, err := clients.CoreClient().GetProcesses(ctx, processesArgs)
	if err != nil {
		log.Fatalln(err)
	}

	t := table.New("ID", "Name", "Type", "Is default", "Url")
	for _, process := range *processes {
		t.AddRow(process.Id, *process.Name, *process.Type, *process.IsDefault, *process.Url)
	}
	t.Print()
}

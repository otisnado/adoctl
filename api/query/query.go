package query

import (
	"context"
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/otisnado/adoctl/api"
)

func ListQueries(projectId string, depth *int) {
	connection := api.GetConnection()

	ctx := context.Background()

	witClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	//expand := workitemtracking.QueryExpand("all")

	queriesArgs := workitemtracking.GetQueriesArgs{
		Project:        &projectId,
		Depth:          depth,
		IncludeDeleted: new(bool),
	}

	responseValue, err := witClient.GetQueries(ctx, queriesArgs)
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	t := tabby.New()
	t.AddHeader("Id", "Name", "Created By", "Path", "Is Public")
	for _, queryReference := range *responseValue {
		for _, childrenReference := range *queryReference.Children {
			t.AddLine(*childrenReference.Id, *childrenReference.Name, *childrenReference.CreatedBy.DisplayName, *childrenReference.Path, *childrenReference.IsPublic)
		}
		index++
	}
	t.Print()
}

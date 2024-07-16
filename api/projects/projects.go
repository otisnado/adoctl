package projects

import (
	"context"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/api"
	"github.com/rodaine/table"
)

func GetProjects() error {
	var ctx context.Context
	var projectArgsIn core.GetProjectsArgs

	responseValue, err := api.CoreClient().GetProjects(ctx, projectArgsIn)
	if err != nil {
		return err
	}
	index := 0
	table := table.New("ID", "Name", "State", "Visibility", "Revision", "Last update")
	for responseValue != nil {

		for _, teamProjectReference := range (*responseValue).Value {
			table.AddRow(*teamProjectReference.Id, *teamProjectReference.Name, *teamProjectReference.State, *teamProjectReference.Visibility, *teamProjectReference.Revision, teamProjectReference.LastUpdateTime)
			index++
		}

		if responseValue.ContinuationToken != "" {
			continuationToken, err := strconv.Atoi(responseValue.ContinuationToken)
			if err != nil {
				return err
			}
			projectsArgs := core.GetProjectsArgs{
				ContinuationToken: &continuationToken,
			}
			responseValue, err = api.CoreClient().GetProjects(ctx, projectsArgs)
			if err != nil {
				return err
			}
		} else {
			responseValue = nil
		}
		table.Print()
	}

	return nil
}

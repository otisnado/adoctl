package teams

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/api/clients"
	"github.com/rodaine/table"
)

var ctx context.Context

func GetTeams(projectId string) {
	var teamArgs core.GetTeamsArgs

	teamArgs.ProjectId = &projectId
	teams, err := clients.CoreClient().GetTeams(ctx, teamArgs)
	if err != nil {
		log.Fatalln(err)
	}

	t := table.New("ID", "Team name", "Project", "Description", "Team URL")
	for _, team := range *teams {
		t.AddRow(team.Id, *team.Name, *team.ProjectName, *team.Description, *team.Url)
	}

	t.Print()
}

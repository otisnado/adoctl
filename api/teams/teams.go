package teams

import (
	"context"
	"log"
	"os"
	"text/template"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/api/clients"
	"github.com/otisnado/adoctl/templates"
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

func GetTeamById(project_id string, team_id string) error {
	projectId := project_id
	teamId := team_id
	expand := true
	teamArgs1 := core.GetTeamArgs{
		ProjectId:      &projectId,
		TeamId:         &teamId,
		ExpandIdentity: &expand,
	}

	team, err := clients.CoreClient().GetTeam(ctx, teamArgs1)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.New("team").Parse(templates.TeamPrint)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, *team)
	return nil

}

func CreateTeam(teamName string, teamDescription string, projectId string) error {
	team := core.WebApiTeam{
		Name:        &teamName,
		Description: &teamDescription,
	}
	createTeamArgs := core.CreateTeamArgs{
		Team:      &team,
		ProjectId: &projectId,
	}

	createdTeam, err := clients.CoreClient().CreateTeam(ctx, createTeamArgs)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.New("team").Parse(templates.TeamPrint)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, *createdTeam)
	return nil
}

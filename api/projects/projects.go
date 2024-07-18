package projects

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/otisnado/adoctl/api"
	"github.com/otisnado/adoctl/templates"
	"github.com/rodaine/table"
)

var ctx context.Context

func GetProjects() error {
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

func GetProjectById(id *string, capabilities *bool, history *bool) error {
	projectArgsIn := core.GetProjectArgs{
		ProjectId:           id,
		IncludeCapabilities: capabilities,
		IncludeHistory:      history,
	}

	res, err := api.CoreClient().GetProject(ctx, projectArgsIn)
	if err != nil {
		return err
	}

	tmpl, err := template.New("project").Parse(templates.TeamProjectYamlTemplate)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, *res)
	return nil
}

func CreateProject(projectName string, projectDescription string, projectSourceControlSystem string, projectProcessTemplateId string, projectVisibility string) error {

	visibility := core.ProjectVisibility(projectVisibility)

	projectProcessTemplate := map[string]string{
		"templateTypeId": projectProcessTemplateId,
	}

	projectVersionControl := map[string]string{
		"sourceControlType": projectSourceControlSystem,
	}

	projectCapabilities := map[string]map[string]string{
		"processTemplate": projectProcessTemplate,
		"versioncontrol":  projectVersionControl,
	}

	TeamProjectToCreate := core.TeamProject{
		Name:         &projectName,
		Description:  &projectDescription,
		Visibility:   &visibility,
		Capabilities: &projectCapabilities,
	}

	queueProject := core.QueueCreateProjectArgs{
		ProjectToCreate: &TeamProjectToCreate,
	}

	operationReference, err := api.CoreClient().QueueCreateProject(ctx, queueProject)
	if err != nil {
		return err
	}

	fmt.Println("Project was created successfully, you can trace it with the following operation reference:", *operationReference.Id)

	return nil
}

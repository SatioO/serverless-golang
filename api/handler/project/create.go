package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/projects/serverless-iam/internal/models"
	"github.com/projects/serverless-iam/internal/project"
)

func CreateProjectHandler(ctx context.Context, event events.APIGatewayProxyRequest) (models.Response, error) {
	var body models.CreateProject

	if err := json.Unmarshal([]byte(event.Body), &body); err != nil {
		panic(err)
	}

	orgId := event.PathParameters["organizationId"]

	projectRepo := models.NewProjectRepo()
	projectSvc := project.NewProject(projectRepo)

	if err := projectSvc.CreateProject(ctx, orgId, &body); err != nil {
		panic(err)
	}

	return models.RespondWithJSON(201, nil)
}

func main() {
	lambda.Start(CreateProjectHandler)
}

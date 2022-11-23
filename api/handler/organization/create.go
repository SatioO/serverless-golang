package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/projects/serverless-iam/internal/models"
	"github.com/projects/serverless-iam/internal/organization"
)

func CreateOrganizationHandler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (models.Response, error) {
	orgRepo := models.NewOrganizationRepo()
	orgSvc := organization.NewOrganization(orgRepo)

	var body models.CreateOrganization

	if err := json.Unmarshal([]byte(event.Body), &body); err != nil {
		panic(err)
	}

	if err := orgSvc.CreateOrganization(ctx, &body); err != nil {
		panic(err)
	}

	return models.RespondWithJSON(201, nil)
}

func main() {
	lambda.Start(CreateOrganizationHandler)
}

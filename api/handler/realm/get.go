package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/projects/serverless-iam/internal/models"
	"github.com/projects/serverless-iam/internal/realm"
)

func GetRealmsHandler() (models.Response, error) {
	repo := models.NewRealmRepo()
	svc := realm.NewRealmService(repo)

	data, _ := svc.GetRealms()

	return models.RespondWithJSON(200, data)
}

func main() {
	lambda.Start(GetRealmsHandler)
}

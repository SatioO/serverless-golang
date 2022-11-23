package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/projects/serverless-iam/internal/models"
	"github.com/projects/serverless-iam/internal/users"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (models.Response, error) {
	userRepo := models.NewUserRepo()
	userSvc := users.NewUsersService(userRepo)

	users, _ := userSvc.GetUsers(ctx)

	return models.RespondWithJSON(200, users)
}

func main() {
	lambda.Start(Handler)
}

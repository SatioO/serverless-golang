package models

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/projects/serverless-iam/internal/db"
)

type User struct {
	UserID string `dynamodbav:"user_id"`
	Name   string `dynamodbav:"name"`
}

type UsersRepo struct{}

func NewUserRepo() *UsersRepo {
	return &UsersRepo{}
}

func (r UsersRepo) GetUsers(ctx context.Context) ([]User, error) {
	svc := db.GetDBClient()
	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("IDP_USER"),
	})
	if err != nil {
		panic(err)
	}

	return toUsers(out.Items)
}

func toUsers(rawItems []map[string]types.AttributeValue) ([]User, error) {
	if len(rawItems) == 0 {
		return nil, nil
	}

	dbItems := []User{}

	if err := attributevalue.UnmarshalListOfMaps(rawItems, &dbItems); err != nil {
		return nil, err
	}

	return dbItems, nil
}

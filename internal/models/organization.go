package models

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/projects/serverless-iam/internal/db"
)

type Organization struct {
	PK   string `dynamodbav:"pk"`
	SK   string `dynamodbav:"sk"`
	Name string `dynamodbav:"name"`
	Tier string `dynamodbav:"tier"`
}

type CreateOrganization struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

type OrganizationRepo struct{}

func NewOrganizationRepo() *Organization {
	return &Organization{}
}

func (Organization) CreateOrganization(payload *CreateOrganization) error {
	client := db.GetDBClient()

	organizationId, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ORGANIZATION"),
		Item: map[string]types.AttributeValue{
			"PK":   &types.AttributeValueMemberS{Value: fmt.Sprintf("ORG#%s", organizationId)},
			"SK":   &types.AttributeValueMemberS{Value: fmt.Sprintf("METADATA#%s", organizationId)},
			"name": &types.AttributeValueMemberS{Value: payload.Name},
			"tier": &types.AttributeValueMemberS{Value: payload.Tier},
		},
	})

	return err
}

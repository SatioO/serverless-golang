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

type Project struct {
	PK        string `dynamodbav:"pk"`
	SK        string `dynamodbav:"sk"`
	Name      string `dynamodbav:"name"`
	ProjectID string `dynamodbav:"project_id"`
}

type CreateProject struct {
	Name        string `json:"name"`
	ProjectID   string `json:"project_id"`
	ProjectType string `json:"type"`
}

type ProjectRepo struct {
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (p ProjectRepo) CreateProject(orgId string, payload *CreateProject) error {
	client := db.GetDBClient()

	projectId, _ := uuid.NewUUID()

	_, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("ORGANIZATION"),
		Item: map[string]types.AttributeValue{
			"PK":         &types.AttributeValueMemberS{Value: fmt.Sprintf("ORG%s", orgId)},
			"SK":         &types.AttributeValueMemberS{Value: fmt.Sprintf("PRO#%s#%s", payload.ProjectType, &projectId)},
			"name":       &types.AttributeValueMemberS{Value: payload.Name},
			"project_id": &types.AttributeValueMemberS{Value: payload.ProjectID},
		},
	})

	return err
}

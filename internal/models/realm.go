package models

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/projects/serverless-iam/internal/db"
)

type Realm struct {
	RealmID     string `dynamodbav:"realmId"`
	Name        string `dynamodbav:"name"`
	DisplayName string `dynamodbav:"display_name"`
	Enabled     bool   `dynamodbav:"enabled"`
}

type RealmRepo struct{}

func NewRealmRepo() *RealmRepo {
	return &RealmRepo{}
}

func (RealmRepo) GetRealms() ([]Realm, error) {
	db := db.GetDBClient()

	out, err := db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("IAM"),
		KeyConditionExpression: aws.String("PK = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: "realm#c693d78f-9fde-4abb-aa5b-73a56a5fa400"},
		},
	})

	log.Println(out)

	if err != nil {
		return nil, err
	}

	return toRealm(out.Items)
}

func (RealmRepo) GetRealm() (*Realm, error) {
	return nil, nil
}

func (RealmRepo) CreateRealm() error {
	return nil
}

func (RealmRepo) UpdateRealm() error {
	return nil
}

func (RealmRepo) DeleteRealm() error {
	return nil
}

func toRealm(data []map[string]types.AttributeValue) ([]Realm, error) {
	realm := []Realm{}
	log.Println("Unmarshelling")

	if err := attributevalue.UnmarshalListOfMaps(data, &realm); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(realm)

	return realm, nil
}

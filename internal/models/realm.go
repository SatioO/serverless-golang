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
	RealmID     string `json:"realm_id" dynamodbav:"PK"`
	Name        string `json:"name" dynamodbav:"name"`
	DisplayName string `json:"display_name" dynamodbav:"display_name"`
	Enabled     string `json:"enabled" dynamodbav:"enabled"`
}

type RealmRepo struct{}

func NewRealmRepo() *RealmRepo {
	return &RealmRepo{}
}

func (RealmRepo) GetRealms() ([]Realm, error) {
	db := db.GetDBClient()

	out, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("IAM"),
	})

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

	if err := attributevalue.UnmarshalListOfMaps(data, &realm); err != nil {
		log.Println(err)
		return nil, err
	}

	return realm, nil
}

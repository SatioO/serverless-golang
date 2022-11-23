package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *dynamodb.Client

func GetDBClient() *dynamodb.Client {
	if client != nil {
		return client
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(lo *config.LoadOptions) error {
		lo.Region = "us-east-1"
		return nil
	})

	if err != nil {
		panic(err)
	}

	client = dynamodb.NewFromConfig(cfg)

	return client
}

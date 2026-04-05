package main

import (
	"context"
	"log"
	"os"
	"ttb-back-email-list/internal/handler"
	"ttb-back-email-list/internal/repository"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	if tableName == "" {
		log.Fatal("DYNAMODB_TABLE_NAME environment variable is not set")
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	repo := repository.New(context.Background(), dynamoClient, tableName)

	h := handler.New(repo)

	lambda.Start(h.Route)
}

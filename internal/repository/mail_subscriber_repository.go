package repository

import (
	"context"
	"fmt"
	"ttb-back-email-list/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MailSubscriberRepository struct {
	client    *dynamodb.Client
	tableName string
	ctx       context.Context
}

func New(ctx context.Context, client *dynamodb.Client, tableName string) *MailSubscriberRepository {
	return &MailSubscriberRepository{
		client:    client,
		tableName: tableName,
		ctx:       ctx,
	}
}

func (r *MailSubscriberRepository) AddSubscriber(email string) error {

	subscriber := &model.MailSubscriber{
		Email: email,
	}

	item, err := attributevalue.MarshalMap(subscriber)
	if err != nil {
		return fmt.Errorf("marshaling mail: %w", err)
	}

	_, err = r.client.PutItem(r.ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})

	if err != nil {
		return fmt.Errorf("putting mail: %w", err)
	}

	return nil
}

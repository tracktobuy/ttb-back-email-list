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


func (r *MailSubscriberRepository) GetSubscribers() ([]model.MailSubscriber, error) {
	var subscribers []model.MailSubscriber

	output, err := r.client.Scan(r.ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("scanning table: %w", err)
	}

	for _, item := range output.Items {
		var subscriber model.MailSubscriber
		err := attributevalue.UnmarshalMap(item, &subscriber)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling subscriber: %w", err)
		}
		subscribers = append(subscribers, subscriber)
	}


	return subscribers, nil
}
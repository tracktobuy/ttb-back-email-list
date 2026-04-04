package model

type MailSubscriber struct {
	Email string `json:"email" dynamodbav:"email"`
}

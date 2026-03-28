package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Event struct {
	Body string `json:"body"`
}

type Data struct {
	Email string `json:"email"`
}

func init() {
	_, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("ERRO AO LER SDK, %v", err)
	}
}

func handler(ctx context.Context, event json.RawMessage) (string, error) {
	var body Event
	if err := json.Unmarshal(event, &body); err != nil {
		return "", fmt.Errorf("DEU ERRO NO BODY: %v", err)
	}

	var data Data
	if err := json.Unmarshal([]byte(body.Body), &data); err != nil {
		return "", fmt.Errorf("DEU ERRO NO BODY: %v", err)
	}
	return "Hello World email ttb " + data.Email, nil
}

func main() {
	lambda.Start(handler)
}

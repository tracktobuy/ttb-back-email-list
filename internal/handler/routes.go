package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ttb-back-email-list/internal/model"
	"ttb-back-email-list/internal/repository"

	"github.com/aws/aws-lambda-go/events"
)

type Routes struct {
	repo *repository.MailSubscriberRepository
}

func New(repo *repository.MailSubscriberRepository) *Routes {
	return &Routes{
		repo: repo,
	}
}

func (r *Routes) Route(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	method := req.RequestContext.HTTP.Method
	path := req.RequestContext.HTTP.Path

	log.Printf("method=%s path=%s", method, path)

	switch method {
	case http.MethodPost:
		return r.handleAddSubscriber(ctx, req)

	default:
		return response(http.StatusNotFound, "Resource not found")
	}

}

func (r *Routes) handleAddSubscriber(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	var subscriber model.MailSubscriber

	err := json.Unmarshal([]byte(req.Body), &subscriber)
	if err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		return response(http.StatusBadRequest, "Invalid request body")
	}

	err = r.repo.AddSubscriber(subscriber.Email)
	if err != nil {
		log.Printf("Error adding subscriber: %v", err)
		return response(http.StatusInternalServerError, "Failed to add subscriber")
	}

	return response(http.StatusOK, nil)
}

func response(statusCode int, body any) (events.APIGatewayV2HTTPResponse, error) {

	var bodyStr string

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, nil
		}

		bodyStr = string(b)
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       bodyStr,
	}, nil
}

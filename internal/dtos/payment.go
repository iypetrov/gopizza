package dtos

import (
	"encoding/json"
	"net/http"
)

type PaymentPublishableKeyResponse struct {
	PublishableKey string `json:"publishableKey"`
}

type PaymentClientSecretResponse struct {
	IntentID     string `json:"intentId"`
	ClientSecret string `json:"clientSecret"`
}

type PaymentIntentRequest struct {
	Email string `json:"email"`
	Total string `json:"total"`
}

func ParseToPaymentIntentRequest(r *http.Request) (PaymentIntentRequest, error) {
	var req PaymentIntentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return PaymentIntentRequest{}, err
	}

	return req, nil
}

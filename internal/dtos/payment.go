package dtos

type PaymentPublishableKeyResponse struct {
	PublishableKey string `json:"publishableKey"`
}

type PaymentClientSecretResponse struct {
	ClientSecret string `json:"clientSecret"`
}

type PaymentIntentRequest struct {
	Total string `json:"total"`
}

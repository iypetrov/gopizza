package dtos

import (
	"encoding/json"
	"net/http"
)

type OrderRequest struct {
	IntentID string `json:"intentId"`
	Total    string `json:"total"`
}

func ParseToOrderRequest(r *http.Request) (OrderRequest, error) {
	var req OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return OrderRequest{}, err
	}

	return req, nil
}

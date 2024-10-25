package dtos

import "net/http"

type CartPizzaRequest struct {
	UserID  string
	PizzaID string
}

func (req *CartPizzaRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if req.UserID == "" {
		errors["userID"] = "UserID is required"
	}

	return errors
}

func ParseToCartRequest(r *http.Request) (CartPizzaRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return CartPizzaRequest{}, err
	}

	var req CartPizzaRequest
	req.UserID = parseString(r, "userID")
	req.PizzaID = parseString(r, "pizzaID")

	return req, nil
}

type CartResponse struct {
	CartID          string
	ProductName     string
	ProductImageUrl string
	ProductPrice    float64
}

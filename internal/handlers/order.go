package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type Order struct {
	srv services.Order
}

func NewOrder(srv services.Order) Order {
	return Order{
		srv: srv,
	}
}

func (hnd *Order) CreateOrder(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToOrderRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return err
	}

	cookie, ok := r.Context().Value(middlewares.CookieName).(dtos.UserCookie)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidCookie))
		return toasts.ErrNotValidCookie
	}

	userID, err := uuid.Parse(cookie.ID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrorInternalServerError(toasts.ErrNotValidUUID)
	}

	total, err := strconv.ParseFloat(req.Total, 64)
	if err != nil {
		return err
	}

	_, err = hnd.srv.CreateOrder(r.Context(), req.IntentID, userID, total)
	if err != nil {
		return err
	}

	return nil
}

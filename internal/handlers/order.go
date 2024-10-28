package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"
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
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return err
	}

	_, err = hnd.srv.CreateOrder(r.Context(), req.IntentID, userID, total)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrOrderCreation))
		return toasts.ErrOrderCreation
	}

	return nil
}

func (hnd *Order) GetOrderByIntentID(w http.ResponseWriter, r *http.Request) error {
	intentID := r.URL.Query().Get("intent_id")

	model, err := hnd.srv.GetOrderByIntentID(r.Context(), intentID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrOrderNotFound))
		return toasts.ErrOrderNotFound
	}

	var dto dtos.OrderResponse
	common.MapFields(&dto, &model)
	dto.Currency = string(model.Currency)
	dto.OrderStatus = string(model.OrderStatus)

	return Render(w, r, components.TrackingData(dto))
}

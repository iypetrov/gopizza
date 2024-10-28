package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"
)

type Cart struct {
	srv services.Cart
}

func NewCart(srv services.Cart) Cart {
	return Cart{
		srv: srv,
	}
}

func (hnd *Cart) AddPizzaToCart(w http.ResponseWriter, r *http.Request) error {
	pizzaID, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
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

	err = hnd.srv.AddPizzaToCart(r.Context(), userID, pizzaID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "item was added to your cart",
		StatusCode: http.StatusNoContent,
	})
	return nil
}

func (hnd *Cart) AddSaladToCart(w http.ResponseWriter, r *http.Request) error {
	saladID, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
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

	err = hnd.srv.AddSaladToCart(r.Context(), userID, saladID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "item was added to your cart",
		StatusCode: http.StatusNoContent,
	})
	return nil
}

func (hnd *Cart) GetCartByUserID(w http.ResponseWriter, r *http.Request) error {
	cookie, ok := r.Context().Value(middlewares.CookieName).(dtos.UserCookie)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidCookie))
		return toasts.ErrNotValidCookie
	}

	id, err := uuid.Parse(cookie.ID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrorInternalServerError(toasts.ErrNotValidUUID)
	}

	models, err := hnd.srv.GetCartByUserID(r.Context(), id)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	var total float64 = 0
	var resps []dtos.CartResponse
	for _, model := range models {
		var dto dtos.CartResponse
		common.MapFields(&dto, &model)
		dto.CartID = model.CartID.String()
		resps = append(resps, dto)
		total += dto.ProductPrice
	}

	return Render(w, r, components.CartItems(resps, cookie.Email, fmt.Sprintf("%.2f", total)))
}

func (hnd *Cart) EmptyCartByUserID(w http.ResponseWriter, r *http.Request) error {
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

	err = hnd.srv.EmptyCartByUserID(r.Context(), userID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	return Render(w, r, components.CartItems([]dtos.CartResponse{}, cookie.Email, "0.00"))
}

func (hnd *Cart) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
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

	models, err := hnd.srv.RemoveItemFromCart(r.Context(), id, userID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	var total float64 = 0
	var resps []dtos.CartResponse
	for _, model := range models {
		var dto dtos.CartResponse
		common.MapFields(&dto, &model)
		dto.CartID = model.CartID.String()
		resps = append(resps, dto)
		total += dto.ProductPrice
	}

	return Render(w, r, components.CartItems(resps, cookie.Email, fmt.Sprintf("%.2f", total)))
}

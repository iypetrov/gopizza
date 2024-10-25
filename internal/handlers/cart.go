package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
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

	if !IsOwnAccount(id, cookie) {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount))
		return toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount)
	}

	pizzaID, err := uuid.Parse(r.URL.Query().Get("pizzaID"))
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrorInternalServerError(toasts.ErrNotValidUUID)
	}

	model, err := hnd.srv.AddPizzaToCart(r.Context(), id, pizzaID)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	var dto dtos.CartPizzaRequest
	common.MapFields(&dto, &model)

	// TODO: return component

	return nil
}

func (hnd *Cart) GetCartByUserID(w http.ResponseWriter, r *http.Request) error {
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

	if !IsOwnAccount(id, cookie) {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount))
		return toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount)
	}

	model, err := hnd.srv.GetCartByUserID(r.Context(), id)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	var dto dtos.CartPizzaRequest
	common.MapFields(&dto, &model)

	// TODO: return component

	return nil
}

func (hnd *Cart) EmptyCartByUserID(w http.ResponseWriter, r *http.Request) error {
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

	if !IsOwnAccount(id, cookie) {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount))
		return toasts.ErrorInternalServerError(toasts.ErrNotOwnAccount)
	}

	err := hnd.srv.EmptyCartByUserID(r.Context(), id)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrorInternalServerError(err)
	}

	// TODO: return component

	return nil
}

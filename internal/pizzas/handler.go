package pizzas

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
)

type PizzaHandler struct {
	service PizzaService
}

func NewHandler(srv PizzaService) *PizzaHandler {
	return &PizzaHandler{service: srv}
}

func (hnd *PizzaHandler) createPizza(w http.ResponseWriter, r *http.Request) error {
	var requestDto UpsertPizzaRequestDto
	body, ok := r.Context().Value(middleware.BodyKey).([]byte)
	if !ok {
		return utils.InvalidJSON()
	}

	err := json.Unmarshal(body, &requestDto)
	if err != nil {
		return utils.InvalidJSON()
	}

	model, err := hnd.service.CreatePizzaModel(r.Context(), requestDto.ToModel())
	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusCreated, model.ToDto())
}

func (hnd *PizzaHandler) getPizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	model, err := hnd.service.GetPizzaModelByID(r.Context(), id)
	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusOK, model.ToDto())
}

func (hnd *PizzaHandler) updatePizza(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)

	var requestDto UpsertPizzaRequestDto
	body, ok := r.Context().Value(middleware.BodyKey).([]byte)
	if !ok {
		return utils.InvalidJSON()
	}

	err := json.Unmarshal(body, &requestDto)
	if err != nil {
		return utils.InvalidJSON()
	}

	model, err := hnd.service.UpdatePizzaModel(r.Context(), id, requestDto.ToModel())
	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusOK, model.ToDto())
}

func (hnd *PizzaHandler) deletePizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	model, err := hnd.service.DeletePizzaModelByID(r.Context(), id)
	if err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusOK, model.ToDto())
}

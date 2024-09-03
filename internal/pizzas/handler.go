package pizzas

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/utils"
	"net/http"
	"strconv"
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

func (hnd *PizzaHandler) getAllPizzas(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ErrNotValidQueryParams
	}

	priceParam := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceParam, 64)
	if err != nil {
		return ErrNotValidQueryParams
	}

	pageSizeParam := r.URL.Query().Get("page-size")
	pageSize, err := strconv.ParseInt(pageSizeParam, 10, 32)
	if err != nil {
		return ErrNotValidQueryParams
	}

	models, err := hnd.service.GetAllPizzaModels(r.Context(), id, price, int32(pageSize))
	if err != nil {
		return err
	}

	var dtos []PizzaResponseDto
	for _, model := range models {
		dtos = append(dtos, model.ToDto())
	}

	return utils.WriteJson(w, http.StatusOK, dtos)
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

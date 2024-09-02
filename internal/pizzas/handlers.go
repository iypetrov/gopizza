package pizzas

import (
	"github.com/go-chi/chi/v5"
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
	var requestDTO CreatePizzaRequestDto
	closeBody, err := utils.ReadRequestBody(r, &requestDTO)
	defer closeBody()
	if err != nil {
		return err
	}

	model, err := hnd.service.CreatePizzaModel(r.Context(), requestDTO.ToModel())
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, model.ToDto())
}

func (hnd *PizzaHandler) getPizzaByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	model, err := hnd.service.GetPizzaModelByID(r.Context(), id)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, model.ToDto())
}
package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dto"
	"github.com/iypetrov/gopizza/internal/mapper"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/service"
	"github.com/iypetrov/gopizza/internal/toast"
	"github.com/iypetrov/gopizza/internal/util"
	"github.com/iypetrov/gopizza/web/template/component"
	"net/http"
	"strconv"
	"time"
)

type Pizza interface {
	CreatePizza(w http.ResponseWriter, r *http.Request) error
	GetPizzaByID(w http.ResponseWriter, r *http.Request) error
	GetAllPizzas(w http.ResponseWriter, r *http.Request) error
	UpdatePizza(w http.ResponseWriter, r *http.Request) error
	DeletePizzaByID(w http.ResponseWriter, r *http.Request) error
}

type PizzaImpl struct {
	srv service.Pizza
}

func NewPizza(srv service.Pizza) *PizzaImpl {
	return &PizzaImpl{srv: srv}
}

func (hnd *PizzaImpl) CreatePizza(w http.ResponseWriter, r *http.Request) error {
	var req dto.PizzaRequest
	b, ok := r.Context().Value(middleware.BodyKey).([]byte)
	if !ok {
		return toast.ErrorInvalidJSON()
	}

	err := json.Unmarshal(b, &req)
	if err != nil {
		return toast.ErrorInvalidJSON()
	}

	_, err = hnd.srv.CreatePizza(r.Context(), mapper.PizzaRequestToModel(req, uuid.New(), time.Now()))
	if err != nil {
		return toast.ErrorInternalServerError(err)
	}

	return util.RenderSuccess(w, r, toast.SuccessPizzaCreated())
}

func (hnd *PizzaImpl) GetPizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	m, err := hnd.srv.GetPizzaByID(r.Context(), id)
	if err != nil {
		return err
	}

	return util.Render(w, r, component.PizzaCard(mapper.PizzaModelToResponse(m)))
}

func (hnd *PizzaImpl) GetAllPizzas(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("last-id")
	lastID, err := uuid.Parse(idParam)
	if err != nil {
		lastID = uuid.Nil
	}

	priceParam := r.URL.Query().Get("last-price")
	lastPrice, err := strconv.ParseFloat(priceParam, 64)
	if err != nil {
		lastPrice = 0
	}

	pageSizeParam := r.URL.Query().Get("page-size")
	pageSize, err := strconv.ParseInt(pageSizeParam, 10, 32)
	if err != nil {
		pageSize = 10
	}

	ms, err := hnd.srv.GetAllPizzas(r.Context(), lastID, lastPrice, int32(pageSize))
	if err != nil {
		return toast.ErrorInternalServerError(err)
	}

	var resp []dto.PizzaResponse
	for _, model := range ms {
		resp = append(resp, mapper.PizzaModelToResponse(model))
	}

	return util.Render(w, r, component.PizzasOverview(resp))
}

func (hnd *PizzaImpl) UpdatePizza(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)

	var req dto.PizzaRequest
	body, ok := r.Context().Value(middleware.BodyKey).([]byte)
	if !ok {
		return toast.ErrorInvalidJSON()
	}

	err := json.Unmarshal(body, &req)
	if err != nil {
		return toast.ErrorInvalidJSON()
	}

	_, err = hnd.srv.UpdateModel(r.Context(), id, mapper.PizzaRequestToModel(req, id, time.Now()))
	if err != nil {
		return toast.ErrorInternalServerError(err)
	}

	return util.RenderSuccess(w, r, toast.SuccessPizzaUpdated())
}

func (hnd *PizzaImpl) DeletePizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)

	_, err := hnd.srv.DeletePizzaByID(r.Context(), id)
	if err != nil {
		return toast.ErrorInternalServerError(err)
	}

	return util.RenderSuccess(w, r, toast.SuccessPizzaDeleted())
}

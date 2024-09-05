package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/dto"
	"github.com/iypetrov/gopizza/internal/mapper"
	"github.com/iypetrov/gopizza/internal/middleware"
	"github.com/iypetrov/gopizza/internal/myerror"
	"github.com/iypetrov/gopizza/internal/service"
	"github.com/iypetrov/gopizza/internal/util"
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
		return myerror.InvalidJSON()
	}

	err := json.Unmarshal(b, &req)
	if err != nil {
		return myerror.InvalidJSON()
	}

	m, err := hnd.srv.CreatePizza(r.Context(), mapper.PizzaRequestToModel(req, uuid.New(), time.Now()))
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusCreated, mapper.PizzaModelToResponse(m))
}

func (hnd *PizzaImpl) GetPizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	m, err := hnd.srv.GetPizzaByID(r.Context(), id)
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, mapper.PizzaModelToResponse(m))
}

func (hnd *PizzaImpl) GetAllPizzas(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("last-id")
	lastID, err := uuid.Parse(idParam)
	if err != nil {
		return myerror.ErrNotValidQueryParams
	}

	priceParam := r.URL.Query().Get("last-price")
	lastPrice, err := strconv.ParseFloat(priceParam, 64)
	if err != nil {
		return myerror.ErrNotValidQueryParams
	}

	pageSizeParam := r.URL.Query().Get("page-size")
	pageSize, err := strconv.ParseInt(pageSizeParam, 10, 32)
	if err != nil {
		return myerror.ErrNotValidQueryParams
	}

	ms, err := hnd.srv.GetAllPizzas(r.Context(), lastID, lastPrice, int32(pageSize))
	if err != nil {
		return err
	}

	var resp []dto.PizzaResponse
	for _, model := range ms {
		resp = append(resp, mapper.PizzaModelToResponse(model))
	}

	return util.WriteJson(w, http.StatusOK, resp)
}

func (hnd *PizzaImpl) UpdatePizza(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)

	var req dto.PizzaRequest
	body, ok := r.Context().Value(middleware.BodyKey).([]byte)
	if !ok {
		return myerror.InvalidJSON()
	}

	err := json.Unmarshal(body, &req)
	if err != nil {
		return myerror.InvalidJSON()
	}

	m, err := hnd.srv.UpdateModel(r.Context(), id, mapper.PizzaRequestToModel(req, id, time.Now()))
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, mapper.PizzaModelToResponse(m))
}

func (hnd *PizzaImpl) DeletePizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)

	m, err := hnd.srv.DeletePizzaByID(r.Context(), id)
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, mapper.PizzaModelToResponse(m))
}

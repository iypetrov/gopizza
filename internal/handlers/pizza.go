package handlers

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"
	"github.com/iypetrov/gopizza/templates/views"

	"net/http"
)

type Pizza struct {
	srv services.Pizza
}

func NewPizza(srv services.Pizza) Pizza {
	return Pizza{srv: srv}
}

func (hnd *Pizza) CreatePizza(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToPizzaRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.PizzaCreateForm(req, errs))
	}

	var p database.CreatePizzaParams
	err = common.MapFields(&p, &req)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	_, err = hnd.srv.CreatePizza(r.Context(), p)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "pizza created successfully",
		StatusCode: http.StatusCreated,
	})
	return Render(w, r, components.PizzaCreateForm(dtos.PizzaRequest{}, make(map[string]string)))
}

func (hnd *Pizza) GetPizzaByID(w http.ResponseWriter, r *http.Request) error {
	//id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	//m, err := hnd.srv.GetPizzaByID(r.Context(), id)
	//if err != nil {
	//	return err
	//}
	//
	//return util.Render(w, r, component.PizzaCard(mapper.PizzaModelToResponse(m)))
	return nil
}

func (hnd *Pizza) GetAllPizzas(w http.ResponseWriter, r *http.Request) error {
	//idParam := r.URL.Query().Get("last-id")
	//lastID, err := uuid.Parse(idParam)
	//if err != nil {
	//	lastID = uuid.Nil
	//}
	//
	//priceParam := r.URL.Query().Get("last-price")
	//lastPrice, err := strconv.ParseFloat(priceParam, 64)
	//if err != nil {
	//	lastPrice = 0
	//}
	//
	//pageSizeParam := r.URL.Query().Get("page-size")
	//pageSize, err := strconv.ParseInt(pageSizeParam, 10, 32)
	//if err != nil {
	//	pageSize = 10
	//}
	//
	//ms, err := hnd.srv.GetAllPizzas(r.Context(), lastID, lastPrice, int32(pageSize))
	//if err != nil {
	//	return toast.ErrorInternalServerError(err)
	//}
	//
	//var resp []dto.PizzaResponse
	//for _, model := range ms {
	//	resp = append(resp, mapper.PizzaModelToResponse(model))
	//}
	//
	//return util.Render(w, r, component.PizzasOverview(resp))
	return nil
}

func (hnd *Pizza) GetAllPizzasAdminOverview(w http.ResponseWriter, r *http.Request) error {
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

	var p database.GetAllPizzasParams
	p.ID = lastID
	p.Price = lastPrice
	p.PageSize = int32(pageSize)
	ms, err := hnd.srv.GetAllPizzas(r.Context(), p)
	if err != nil {
		return toasts.ErrPizzaFailedToLoad
	}

	var resps []dtos.PizzaResponse
	for _, model := range ms {
		var dto dtos.PizzaResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	return Render(w, r, views.AdminPizzasOverview(resps))
}

func (hnd *Pizza) UpdatePizza(w http.ResponseWriter, r *http.Request) error {
	//id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	//
	//var req dto.PizzaRequest
	//body, ok := r.Context().Value(middleware.BodyKey).([]byte)
	//if !ok {
	//	return toast.ErrorInvalidJSON()
	//}
	//
	//err := json.Unmarshal(body, &req)
	//if err != nil {
	//	return toast.ErrorInvalidJSON()
	//}
	//
	//_, err = hnd.srv.UpdateModel(r.Context(), id, mapper.PizzaRequestToModel(req, id, time.Now()))
	//if err != nil {
	//	return toast.ErrorInternalServerError(err)
	//}
	//
	//return util.RenderSuccess(w, r, toast.SuccessPizzaUpdated())
	return nil
}

func (hnd *Pizza) DeletePizzaByID(w http.ResponseWriter, r *http.Request) error {
	//id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	//
	//_, err := hnd.srv.DeletePizzaByID(r.Context(), id)
	//if err != nil {
	//	return toast.ErrorInternalServerError(err)
	//}
	//
	//return util.RenderSuccess(w, r, toast.SuccessPizzaDeleted())
	return nil
}

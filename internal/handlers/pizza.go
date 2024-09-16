package handlers

import (
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"
	"net/http"
)

type Pizza interface {
	CreatePizza(w http.ResponseWriter, r *http.Request) error
	GetPizzaByID(w http.ResponseWriter, r *http.Request) error
	GetAllPizzas(w http.ResponseWriter, r *http.Request) error
	UpdatePizza(w http.ResponseWriter, r *http.Request) error
	DeletePizzaByID(w http.ResponseWriter, r *http.Request) error
}

type PizzaImpl struct {
	srv services.Pizza
}

func NewPizza(srv services.Pizza) *PizzaImpl {
	return &PizzaImpl{srv: srv}
}

func (hnd *PizzaImpl) CreatePizza(w http.ResponseWriter, r *http.Request) error {
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
	RedirectAdminHomePage(r)
	return Render(w, r, components.PizzaCreateForm(dtos.PizzaRequest{}, make(map[string]string)))
}

func (hnd *PizzaImpl) GetPizzaByID(w http.ResponseWriter, r *http.Request) error {
	//id, _ := r.Context().Value(middleware.UUIDKey).(uuid.UUID)
	//m, err := hnd.srv.GetPizzaByID(r.Context(), id)
	//if err != nil {
	//	return err
	//}
	//
	//return util.Render(w, r, component.PizzaCard(mapper.PizzaModelToResponse(m)))
	return nil
}

func (hnd *PizzaImpl) GetAllPizzas(w http.ResponseWriter, r *http.Request) error {
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

func (hnd *PizzaImpl) UpdatePizza(w http.ResponseWriter, r *http.Request) error {
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

func (hnd *PizzaImpl) DeletePizzaByID(w http.ResponseWriter, r *http.Request) error {
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

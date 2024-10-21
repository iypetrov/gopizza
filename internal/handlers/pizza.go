package handlers

import (
	"strings"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/common"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/iypetrov/gopizza/templates/components"
	"github.com/iypetrov/gopizza/templates/views"

	"net/http"
)

type Pizza struct {
	srv    services.Pizza
	srvImg services.Image
}

func NewPizza(srv services.Pizza, srvImg services.Image) Pizza {
	return Pizza{
		srv:    srv,
		srvImg: srvImg,
	}
}

func (hnd *Pizza) AdminCreatePizza(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToPizzaRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.PizzaCreateForm(req, errs))
	}

	imageUrl, err := hnd.srvImg.UploadImage(r.Context(), req.Image)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	var p database.CreatePizzaParams
	err = common.MapFields(&p, &req)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}
	parts := strings.Split(imageUrl, ".")
	if len(parts) > 0 {
		p.ImageUrl = parts[0]
	} else {
		p.ImageUrl = imageUrl
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

func (hnd *Pizza) AdminGetAllPizzas(w http.ResponseWriter, r *http.Request) error {
	models, err := hnd.srv.GetAllPizzas(r.Context())
	if err != nil {
		return toasts.ErrPizzaFailedToLoad
	}

	var resps []dtos.PizzaResponse
	for _, model := range models {
		var dto dtos.PizzaResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	return Render(w, r, views.AdminPizzasOverview(resps))
}

func (hnd *Pizza) AdminDeletePizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}

	models, err := hnd.srv.DeletePizzaByID(r.Context(), id)
	if err != nil {
		return toasts.ErrorInternalServerError(err)
	}

	var resps []dtos.PizzaResponse
	for _, model := range models {
		var dto dtos.PizzaResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "pizza deleted successfully",
		StatusCode: http.StatusNoContent,
	})
	return Render(w, r, views.AdminPizzasOverview(resps))
}

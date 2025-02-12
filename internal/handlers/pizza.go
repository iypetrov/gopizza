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

	models, err := hnd.srv.CreatePizza(r.Context(), p)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.PizzaCreateForm(req, make(map[string]string)))
	}

	var resps []dtos.PizzaResponse
	for _, model := range models {
		var dto dtos.PizzaResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "pizza created successfully",
		StatusCode: http.StatusCreated,
	})
	return Render(w, r, views.AdminPizzasOverview(resps))
}

func (hnd *Pizza) GetAllPizzas(w http.ResponseWriter, r *http.Request) error {
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

	return Render(w, r, views.PizzasOverview(resps))
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

func (hnd *Pizza) GetPizzaByID(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}
	model, err := hnd.srv.GetPizzaByID(r.Context(), id)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrPizzaFailedToLoad
	}

	var dto dtos.PizzaResponse
	common.MapFields(&dto, &model)

	return Render(w, r, components.PizzaDetailsForm(dto))
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

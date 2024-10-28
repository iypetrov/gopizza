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

type Salad struct {
	srv    services.Salad
	srvImg services.Image
}

func NewSalad(srv services.Salad, srvImg services.Image) Salad {
	return Salad{
		srv:    srv,
		srvImg: srvImg,
	}
}

func (hnd *Salad) AdminCreateSalad(w http.ResponseWriter, r *http.Request) error {
	req, err := dtos.ParseToSaladRequest(r)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.SaladCreateForm(req, make(map[string]string)))
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return Render(w, r, components.SaladCreateForm(req, errs))
	}

	imageUrl, err := hnd.srvImg.UploadImage(r.Context(), req.Image)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.SaladCreateForm(req, make(map[string]string)))
	}

	var p database.CreateSaladParams
	err = common.MapFields(&p, &req)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.SaladCreateForm(req, make(map[string]string)))
	}
	parts := strings.Split(imageUrl, ".")
	if len(parts) > 0 {
		p.ImageUrl = parts[0]
	} else {
		p.ImageUrl = imageUrl
	}

	models, err := hnd.srv.CreateSalad(r.Context(), p)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return Render(w, r, components.SaladCreateForm(req, make(map[string]string)))
	}

	var resps []dtos.SaladResponse
	for _, model := range models {
		var dto dtos.SaladResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "salad created successfully",
		StatusCode: http.StatusCreated,
	})
	return Render(w, r, views.AdminSaladsOverview(resps))
}

func (hnd *Salad) GetAllSalads(w http.ResponseWriter, r *http.Request) error {
	models, err := hnd.srv.GetAllSalads(r.Context())
	if err != nil {
		return toasts.ErrSaladFailedToLoad
	}

	var resps []dtos.SaladResponse
	for _, model := range models {
		var dto dtos.SaladResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	return Render(w, r, views.SaladsOverview(resps))
}

func (hnd *Salad) AdminGetAllSalads(w http.ResponseWriter, r *http.Request) error {
	models, err := hnd.srv.GetAllSalads(r.Context())
	if err != nil {
		return toasts.ErrSaladFailedToLoad
	}

	var resps []dtos.SaladResponse
	for _, model := range models {
		var dto dtos.SaladResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	return Render(w, r, views.AdminSaladsOverview(resps))
}

func (hnd *Salad) GetSaladByID(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}
	model, err := hnd.srv.GetSaladByID(r.Context(), id)
	if err != nil {
		toasts.AddToast(w, toasts.ErrorInternalServerError(err))
		return toasts.ErrSaladFailedToLoad
	}

	var dto dtos.SaladResponse
	common.MapFields(&dto, &model)

	return Render(w, r, components.SaladDetailsForm(dto))
}

func (hnd *Salad) AdminDeleteSaladByID(w http.ResponseWriter, r *http.Request) error {
	id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
	if !ok {
		toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
		return toasts.ErrNotValidUUID
	}

	models, err := hnd.srv.DeleteSaladByID(r.Context(), id)
	if err != nil {
		return toasts.ErrorInternalServerError(err)
	}

	var resps []dtos.SaladResponse
	for _, model := range models {
		var dto dtos.SaladResponse
		common.MapFields(&dto, &model)
		resps = append(resps, dto)
	}

	toasts.AddToast(w, toasts.Toast{
		Message:    "salad deleted successfully",
		StatusCode: http.StatusNoContent,
	})
	return Render(w, r, views.AdminSaladsOverview(resps))
}

package handlers

import (
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/middlewares"
	"github.com/iypetrov/gopizza/internal/services"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type Image struct {
	srv services.Image
}

func NewImage(srv services.Image) Image {
	return Image{
		srv: srv,
	}
}

func (hnd *Image) GetImage(w http.ResponseWriter, r *http.Request) {
    id, ok := r.Context().Value(middlewares.UUIDKey).(uuid.UUID)
    if !ok {
        toasts.AddToast(w, toasts.ErrorInternalServerError(toasts.ErrNotValidUUID))
        return
    }

    iorc, err := hnd.srv.GetImage(r.Context(), id)
    if err != nil {
        toasts.AddToast(w, toasts.ErrorNotFound(err))
        return
    }
    defer iorc.Close()

    w.Header().Set("Content-Type", "image/png")
    if _, err := io.Copy(w, iorc); err != nil {
        toasts.AddToast(w, toasts.ErrorInternalServerError(err))
        return
    }
}

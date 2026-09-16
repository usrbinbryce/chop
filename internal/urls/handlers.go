package urls

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/usrbinbryce/chop/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var urlToCreate createUrlRequest
	if err := json.Read(r, &urlToCreate); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	createdUrl, err := h.service.CreateURL(r.Context(), urlToCreate)
	if err != nil {
		switch {
		case errors.Is(err, ErrShortCodeRequired), errors.Is(err, ErrDestinationRequired):
			json.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			slog.Error("url handler: failed to create URL", "error", err)
			json.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	json.Write(w, http.StatusCreated, createdUrl)
}

package urls

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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
	var urlToCreate createURLPayload

	if err := json.Read(r, &urlToCreate); err != nil {
		if errors.Is(err, json.ErrReqTooLarge) {
			json.WriteError(w, http.StatusBadRequest, err.Error())
		}

		json.WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	createdUrl, err := h.service.CreateURL(r.Context(), urlToCreate)
	if err != nil {
		switch {
		case errors.Is(err, ErrShortCodeRequired), errors.Is(err, ErrDestinationRequired), errors.Is(err, ErrShortCodeTaken):
			json.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			slog.Error("url handler: failed to create URL", "error", err)
			json.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	json.Write(w, http.StatusCreated, createdUrl)
}

func (h *handler) GetURLMetadata(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")

	url, err := h.service.GetURLFromShortCode(r.Context(), shortCode)
	if err != nil {
		switch {
		case errors.Is(err, ErrShortCodeRequired):
			json.WriteError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, pgx.ErrNoRows):
			json.WriteError(w, http.StatusNotFound, "URL not found")
		default:
			slog.Error("url handler: failed to get URL", "error", err)
			json.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	json.Write(w, http.StatusOK, url)
}

func (h *handler) RedirectToDestination(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")

	url, err := h.service.GetURLFromShortCode(r.Context(), shortCode)
	if err != nil {
		switch {
		case errors.Is(err, ErrShortCodeRequired):
			json.WriteError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, pgx.ErrNoRows):
			json.WriteError(w, http.StatusNotFound, "URL not found")
		default:
			json.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
	}

	http.Redirect(w, r, url.Destination, http.StatusMovedPermanently)
}

package httpa

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"url_shortener/internal/application"
	"url_shortener/internal/domain"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	resolve application.ResolveShortLinkI
	create  application.CreateShortLinkI
}

func NewHandler(resolve application.ResolveShortLinkI, create application.CreateShortLinkI) *handler {
	return &handler{
		resolve: resolve,
		create:  create,
	}
}

func (h *handler) Resolve(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	url, err := h.resolve.Execute(code)
	if err != nil {
		log.Println("ERROR CREATE:", err)
		handleError(w, err)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	log.Println("handler reached")
	link, err := h.create.Execute(req.URL, nil)
	if err != nil {
		log.Println("ERROR CREATE:", err)
		handleError(w, err)
		return
	}
	log.Println("handler reached")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)
}

func handleError(w http.ResponseWriter, err error) {
	log.Println("handler reached")
	switch {
	case errors.Is(err, domain.ErrLinkNotFound):
		http.Error(w, "link not found", http.StatusNotFound)

	case errors.Is(err, domain.ErrLinkExpired):
		http.Error(w, "link expired", http.StatusGone)

	default:
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

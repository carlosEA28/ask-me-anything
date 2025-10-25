package api

import (
	"net/http"

	"github.com/carlosEA28/ask-me-anything.git/internal/store/pgstore"
	"github.com/go-chi/chi"
)

type apiHandler struct {
	q      *pgstore.Queries
	router *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	a.router = r

	return a
}

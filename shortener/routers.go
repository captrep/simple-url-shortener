package shortener

import "github.com/go-chi/chi/v5"

func Router() *chi.Mux {
	r := chi.NewMux()
	r.Post("/", createLink)

	return r
}

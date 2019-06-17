package pairing

import (
	"github.com/go-chi/chi"
	"pairot/persistence"
)

func Routes(db persistence.DB) *chi.Mux {
	router := chi.NewRouter()
	pairHandlers := Handler{processor: Processor{db}}
	router.Post("/", pairHandlers.PostHandler)
	return router
}

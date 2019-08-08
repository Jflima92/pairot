package pairing

import (
	"github.com/go-chi/chi"
	"pairot/persistence"
)

func Routes(db persistence.DB) *chi.Mux {
	router := chi.NewRouter()
	pairHandlers := NewHandler(NewProcessor(db))
	router.Post("/", pairHandlers.PostHandler)
	return router
}

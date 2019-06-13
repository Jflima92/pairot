package pairing

import (
	"github.com/go-chi/chi"
	"pairot/persistence/mongodb"
)

func Routes(m *mongodb.Connection) *chi.Mux {
	router := chi.NewRouter()
	pairHandlers := Handler{processor: Processor{m}}
	router.Post("/", pairHandlers.PostHandler)
	return router
}

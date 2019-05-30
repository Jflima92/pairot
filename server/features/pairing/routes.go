package pairing

import (
	"github.com/go-chi/chi"
	"pairot/persistence/mongodb"
)

func Routes(m *mongodb.Connection) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", Pairing{Conn: m}.postHandler)
	return router
}
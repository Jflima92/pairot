package pairing

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Handler struct {
	processor Processor
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseForm()

	if formErr != nil {
		render.JSON(w, r, createSlackErrorResponse("Error processing request"))
		log.Print(formErr)
	}

	res := h.processor.Process(r.Form)
	render.JSON(w, r, res)
}

func (h *Handler) getHandler(w http.ResponseWriter, r *http.Request) {
}

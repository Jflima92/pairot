package pairing

import (
	"github.com/go-chi/render"
	"net/http"
	"net/url"
	"strings"
)

type PairHandler struct {
	processor Processor
}

func NewHandler(processor Processor) Handler {
	return &PairHandler{
		processor: processor,
	}
}

func (h *PairHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseForm()

	if formErr != nil || len(r.Form) == 0 {
		render.JSON(w, r, createSlackErrorResponse("Error processing request"))
		return
	}

	input, err := getInputFromRequest(r.Form)
	if err != nil {
		render.JSON(w, r, *err)
		return
	}

	res := h.processor.Process(*input)
	render.JSON(w, r, res)
}

func (h *PairHandler) getHandler(w http.ResponseWriter, r *http.Request) {
}


func getInputFromRequest(form url.Values) (*Input, *SlackResponse) {
	var teamName string
	var args []string

	for key, value := range form {
		if key == "channel_name" {
			teamName = value[0]
		} else if key == "text" {
			args = strings.Split(value[0], ", ")
		}
	}

	teamNameWords := strings.Split(teamName, "-")
	if teamNameWords[0] == "team" {
		teamName = strings.Title(teamNameWords[1])
	} else
	{
		slackError := createSlackErrorResponse("Error processing request")
		return nil, &slackError
	}

	return &Input{
		TeamName: teamName,
		Arguments: args,
	}, nil
}
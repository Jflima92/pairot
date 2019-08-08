package pairing

import (
	"net/http"
)

type Processor interface {
	Process(input Input) SlackResponse
}

type Handler interface {
	PostHandler(w http.ResponseWriter, r *http.Request)
	getHandler(w http.ResponseWriter, r *http.Request)
}

type Generator interface {
	Generate(member []Member, seed int64) [][]Member
}
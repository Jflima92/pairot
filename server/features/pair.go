package pair

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Team struct {
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}

type Member struct {
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{teamName}", GetPairs)
	return router
}

func GetPairs(w http.ResponseWriter, r *http.Request) {
	teamName := chi.URLParam(r, "teamName")

	jsonFile, err := ioutil.ReadFile("resources/teams.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened teams.json")

	var teams []Team
	err1 := json.Unmarshal([]byte(jsonFile), &teams)

	if err1 != nil {
		fmt.Println(err1)
	}

	for i := 0; i < len(teams); i++ {
		if teams[i].Name == teamName {
			render.JSON(w, r, GetTeamPairs(teams[i]))

		}
	}

	render.PlainText(w, r, "Team not found")
}

func GetTeamPairs(team Team) [][]Member {
	rand.Seed(time.Now().UnixNano())
	members := team.Members
	shuffledMembers := Shuffle(members)

	var pairs [][]Member
	for i := 0; i < len(members); i = i + 2 {
		pair := []Member{shuffledMembers[i], shuffledMembers[i+1]}
		pairs = append(pairs, pair)
	}
	return pairs
}

func Shuffle(members []Member) []Member {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Member, len(members))
	perm := r.Perm(len(members))
	for i, randIndex := range perm {
		ret[i] = members[randIndex]
	}
	return ret
}

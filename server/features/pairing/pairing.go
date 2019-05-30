package pairing

import (
	"context"
	"fmt"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DBConnection interface {
	GetDatabase(string) *mongo.Database
}

type Pairing struct {
	Conn DBConnection
}

func (p Pairing) postHandler(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseForm()
	db := p.Conn.GetDatabase("pairot")

	if formErr != nil {
		render.JSON(w, r, createSlackErrorResponse("Error processing request"))
		log.Print(formErr)
	}

	teamName := getTeamName(r.Form)

	var collection = db.Collection("Teams")

	var team Team
	filter := bson.D{{"name", teamName}}
	err := collection.FindOne(context.TODO(), filter).Decode(&team)

	if err != nil {
		render.JSON(w, r, createSlackErrorResponse("Team not found"))
		log.Print(err)
	} else {
		teamPairs := GetTeamPairs(team)
		render.JSON(w, r, convertToSlackResponse(teamPairs))
	}
}

func getTeamName(form url.Values) string {
	var teamName string

	for key, value := range form {
		if key == "channel_name" {
			teamName = value[0]
		}
		if key == "text" {
			println(value[0])
		}
	}

	teamNameWords := strings.Split(teamName, "-")
	if teamNameWords[0] == "team" {
		teamName = strings.Title(teamNameWords[1])
	} else
	{
		teamName = "Falcon"
	}

	return teamName
}

func createSlackErrorResponse(msg string) SlackResponse {
	return SlackResponse{
		ResponseType: "ephemeral",
		Text:         msg,
	}
}

func convertToSlackResponse(pairs [][]Member) SlackResponse {
	var response SlackResponse
	response.ResponseType = "ephemeral"
	response.Text = "Here are today's pairs:"
	var attachments = make([]SlackAttachment, len(pairs))
	for i := 0; i < len(pairs); i++ {
		attachments[i].Text = fmt.Sprintf("%s - %s", pairs[i][0].Name, pairs[i][1].Name)
	}
	response.Attachments = attachments
	return response
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

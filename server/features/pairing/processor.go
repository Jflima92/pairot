package pairing

import (
	"fmt"
	"math/rand"
	"net/url"
	"pairot/persistence"
	"strings"
	"time"
)

type Processor struct {
	db persistence.DB
}

func (p Processor) Process(form url.Values) SlackResponse {
	teamName := getTeamNameFromRequest(form)
	dbTeam, err := p.db.FindTeamByName(teamName)

	if err != nil {
		return createSlackErrorResponse("Team not found")
	} else {
		var team Team
		err := p.db.Decode(dbTeam, &team)
		if err != nil {
			return createSlackErrorResponse("Error while looking for team")
		}
		teamPairs := getTeamPairs(team)
		updatedTeamMembers := updateTeam(teamPairs).Members
		updError := p.db.UpdateTeamMembers(teamName, updatedTeamMembers)
		if updError != nil {
			return createSlackErrorResponse("Error updating team")
		}
		return convertToSlackResponse(teamPairs)
	}
}

func getTeamNameFromRequest(form url.Values) string {
	var teamName string

	for key, value := range form {
		if key == "channel_name" {
			teamName = value[0]
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
	var attachments = make([]SlackAttachment, len(pairs))
	response.ResponseType = "ephemeral"
	response.Text = "Here are today's pairs:"

	for i := 0; i < len(pairs); i++ {
		var attachmentText string
		if len(pairs[i]) == 2 {
			attachmentText = fmt.Sprintf("%s - %s", pairs[i][0].Name, pairs[i][1].Name)
		} else {
			attachmentText = fmt.Sprintf("%s", pairs[i][0].Name)
		}
		attachments[i].Text = attachmentText
	}

	response.Attachments = attachments
	return response
}

func getTeamPairs(team Team) [][]Member {
	rand.Seed(time.Now().UnixNano())
	members := team.Members
	shuffledMembers := shuffle(members)

	var pairs [][]Member
	for i := 0; i < len(members); i = i + 2 {
		var pair []Member
		if i == len(members)-1 && i+1%2 != 0 {
			pair = []Member{shuffledMembers[i]}
		} else {
			pair = []Member{shuffledMembers[i], shuffledMembers[i+1]}
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func shuffle(members []Member) []Member {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Member, len(members))
	perm := r.Perm(len(members))
	for i, randIndex := range perm {
		ret[i] = members[randIndex]
	}
	return ret
}

func updateTeam(members [][]Member) Team {
	var team Team
	for i := 0; i < len(members); i++ {
		pair := members[i]
		if len(pair) == 2 {
			r := rand.New(rand.NewSource(time.Now().Unix()))
			if pair[0].Locked == false && pair[1].Locked == false {
				v := r.Intn(2)
				pair[v].Locked = true
			} else {
				pair[0].Locked = !pair[0].Locked
				pair[1].Locked = !pair[1].Locked
			}
			team.Members = append(team.Members, pair[0])
			team.Members = append(team.Members, pair[1])
		} else {
			team.Members = append(team.Members, pair[0])
		}
	}
	return team
}

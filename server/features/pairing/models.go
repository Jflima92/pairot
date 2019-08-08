package pairing

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID      *primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Members []Member           `json:"members" bson:"members"`
}

type Member struct {
	Name   string `json:"name" bson:"name"`
	Locked bool   `json:"locked" bson:"locked"`
}

type Input struct {
	TeamName string
	Arguments []string
}

type SlackResponse struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Text string `json:"text"`
}

func createSlackErrorResponse(msg string) SlackResponse {
	return SlackResponse{
		ResponseType: "ephemeral",
		Text:         msg,
	}
}

func createSlackResponse(pairs [][]Member) SlackResponse {
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

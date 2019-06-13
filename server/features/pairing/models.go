package pairing

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID      *primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Members []Member           `json:"members" bson:"members"`
}

type Member struct {
	Name   string `json:"name" bson:"name"`
	Locked bool   `json:"locked" bson:"locked"`
}

type SlackResponse struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Text string `json:"text"`
}

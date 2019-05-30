package pairing

type Team struct {
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}

type Member struct {
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Text string `json:"text"`
}
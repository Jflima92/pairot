package pairing_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"pairot/features/pairing"
	"pairot/features/pairing/mocks"
	"testing"
)

func TestPairProcessorShouldReturnsSuccessfulResponse(t *testing.T) {
	//given
	var actTeamName string
	expTeamName := "Falcon"
	team := getTeam(expTeamName, false)
	expPairNumber := len(team.Members) / 2

	mockDB := &mocks.MockDB{
		FindTeamByNameFn: func(teamName string) (bytes []byte, e error) {
			actTeamName = teamName
			return bson.Marshal(team)
		},
		DecodeFn: func(data []byte, val interface{}) error {
			err := bson.Unmarshal(data, val)
			return err
		},
		UpdateTeamMembersFn: func(teamName string, members interface{}) error {
			return nil
		},
	}

	input := pairing.Input{
		TeamName:  "Falcon",
		Arguments: []string{"force", "true"},
	}
	expRes := pairing.SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Here are today's pairs:",
	}

	p := pairing.NewProcessor(mockDB)

	//when
	response := p.Process(input)

	//then
	assert.Equal(t, expTeamName, actTeamName)
	assert.Equal(t, 1, mockDB.FindTeamByNameCalled)
	assert.Equal(t, 1, mockDB.DecodeCalled)
	assert.Equal(t, 1, mockDB.UpdateTeamMembersCalled)
	assert.Equal(t, expRes.ResponseType, response.ResponseType)
	assert.Equal(t, expRes.Text, response.Text)
	assert.Equal(t, expPairNumber, len(response.Attachments))
}

func TestPairProcessorGivenOddNumberOfMemberShouldReturnsSuccessfulResponse(t *testing.T) {
	//given
	var actTeamName string
	expTeamName := "Falcon"
	team := getTeam(expTeamName, true)
	expPairNumber := len(team.Members) / 2 + 1

	mockDB := &mocks.MockDB{
		FindTeamByNameFn: func(teamName string) (bytes []byte, e error) {
			actTeamName = teamName
			return bson.Marshal(team)
		},
		DecodeFn: func(data []byte, val interface{}) error {
			err := bson.Unmarshal(data, val)
			return err
		},
		UpdateTeamMembersFn: func(teamName string, members interface{}) error {
			return nil
		},
	}

	input := pairing.Input{
		TeamName:  "Falcon",
		Arguments: []string{"force", "true"},
	}
	expRes := pairing.SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Here are today's pairs:",
	}

	p := pairing.NewProcessor(mockDB)

	//when
	response := p.Process(input)

	//then
	assert.Equal(t, expTeamName, actTeamName)
	assert.Equal(t, 1, mockDB.FindTeamByNameCalled)
	assert.Equal(t, 1, mockDB.DecodeCalled)
	assert.Equal(t, 1, mockDB.UpdateTeamMembersCalled)
	assert.Equal(t, expRes.ResponseType, response.ResponseType)
	assert.Equal(t, expRes.Text, response.Text)
	assert.Equal(t, expPairNumber, len(response.Attachments))
}

func TestPairProcessorGivenFindTeamFailsShouldReturnUnsuccessfulResponse(t *testing.T) {
	//given
	var actTeamName string
	expTeamName := "Falcon"

	mockDB := &mocks.MockDB{
		FindTeamByNameFn: func(teamName string) (bytes []byte, e error) {
			actTeamName = teamName
			return nil, errors.New("team not found")
		},
	}

	input := pairing.Input{
		TeamName:  "Falcon",
		Arguments: []string{"force", "true"},
	}
	expRes := pairing.SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Team not found",
	}

	p := pairing.NewProcessor(mockDB)

	//when
	response := p.Process(input)

	//then
	assert.Equal(t, expTeamName, actTeamName)
	assert.Equal(t, expRes, response)
}

func TestPairProcessorGivenDecodeFailsShouldReturnUnsuccessfulResponse(t *testing.T) {
	//given
	var actTeamName string
	expTeamName := "Falcon"
	team := getTeam(expTeamName, false)

	mockDB := &mocks.MockDB{
		FindTeamByNameFn: func(teamName string) (bytes []byte, e error) {
			actTeamName = teamName
			return bson.Marshal(team)
		},
		DecodeFn: func(data []byte, val interface{}) error {
			return errors.New("failed to decode")
		},
	}

	input := pairing.Input{
		TeamName:  "Falcon",
		Arguments: []string{"force", "true"},
	}
	expRes := pairing.SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Error while looking for team",
	}

	p := pairing.NewProcessor(mockDB)

	//when
	response := p.Process(input)

	//then
	assert.Equal(t, expTeamName, actTeamName)
	assert.Equal(t, expRes, response)
}

func TestPairProcessorGivenUpdateFailsShouldReturnUnsuccessfulResponse(t *testing.T) {
	//given
	var actTeamName string
	expTeamName := "Falcon"
	team := getTeam(expTeamName, false)

	mockDB := &mocks.MockDB{
		FindTeamByNameFn: func(teamName string) (bytes []byte, e error) {
			actTeamName = teamName
			return bson.Marshal(team)
		},
		DecodeFn: func(data []byte, val interface{}) error {
			return bson.Unmarshal(data, val)
		},
		UpdateTeamMembersFn: func(teamName string, members interface{}) error {
			return errors.New("failed to update")
		},
	}

	input := pairing.Input{
		TeamName:  "Falcon",
		Arguments: []string{"force", "true"},
	}
	expRes := pairing.SlackResponse{
		ResponseType: "ephemeral",
		Text:         "Error updating team",
	}

	p := pairing.NewProcessor(mockDB)

	//when
	response := p.Process(input)

	//then
	assert.Equal(t, expTeamName, actTeamName)
	assert.Equal(t, expRes, response)
}

func getTeam(teamName string, isOdd bool) pairing.Team {
	team := pairing.Team{
		Name: teamName,
		Members: []pairing.Member{
			{
				Name:   "Jorge",
				Locked: true,
			},
			{
				Name:   "Ruba",
				Locked: true,
			},
			{
				Name:   "Mario",
				Locked: false,
			},
			{
				Name:   "Fede",
				Locked: false,
			},
			{
				Name:   "Linh",
				Locked: false,
			},
			{
				Name:   "Felix",
				Locked: false,
			},
		},
	}
	if isOdd == true {
		team.Members = append(team.Members, pairing.Member{Name: "Chinmay", Locked: false})
	}
	return team
}

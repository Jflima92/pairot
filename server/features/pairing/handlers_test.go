package pairing_test

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/url"
	"pairot/features/pairing"
	"pairot/features/pairing/mocks"
	"strings"
	"testing"
)

func TestGivenAProperRequestWhenHandlerCalledShouldReturnExpectedResponse(t *testing.T) {
	//given
	teamName := "falcon"
	args := []string{"force", "true"}
	var actInput pairing.Input
	expInput := pairing.Input{
		TeamName:  strings.Title(teamName),
		Arguments: args,
	}
	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "localhost:8080", nil)
	r.Form = url.Values{
		"channel_name": []string{"team-" + teamName},
		"text":         []string{"force, true"},
	}
	mockProcessor := &mocks.MockProcessor{
		ProcessFn: func(input pairing.Input) pairing.SlackResponse {
			actInput = input
			return pairing.SlackResponse{}
		},
	}
	handler := pairing.NewHandler(mockProcessor)

	//when
	handler.PostHandler(w, r)

	//then
	assert.Equal(t, 1, mockProcessor.ProcessCalled)
	assert.Equal(t, expInput, actInput)
}

func TestGivenARequestWithWrongContentWhenHandlerCalledShouldReturnError(t *testing.T) {
	//given
	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "localhost:8080", nil)
	r.PostForm = nil
	r.Form = nil
	r.Body = nil
	mockProcessor := &mocks.MockProcessor{
		ProcessFn: func(input pairing.Input) pairing.SlackResponse {
			return pairing.SlackResponse{}
		},
	}
	handler := pairing.NewHandler(mockProcessor)

	//when
	handler.PostHandler(w, r)

	//then
	assert.Equal(t, 0, mockProcessor.ProcessCalled)
}

func TestGivenARequestWithEmptyContentWhenHandlerCalledShouldReturnError(t *testing.T) {
	//given
	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "localhost:8080", nil)
	r.Form = nil
	mockProcessor := &mocks.MockProcessor{
		ProcessFn: func(input pairing.Input) pairing.SlackResponse {
			return pairing.SlackResponse{}
		},
	}
	handler := pairing.NewHandler(mockProcessor)

	//when
	handler.PostHandler(w, r)

	//then
	assert.Equal(t, 0, mockProcessor.ProcessCalled)
}

func TestGivenARequestWithWrongTeamWhenHandlerCalledShouldReturnError(t *testing.T) {
	//given
	w := &httptest.ResponseRecorder{}
	r := httptest.NewRequest("POST", "localhost:8080", nil)
	r.Form = url.Values{
		"channel_name": []string{"wrongName"},
		"text":         []string{"force, true"},
	}
	mockProcessor := &mocks.MockProcessor{
		ProcessFn: func(input pairing.Input) pairing.SlackResponse {
			return pairing.SlackResponse{}
		},
	}
	handler := pairing.NewHandler(mockProcessor)

	//when
	handler.PostHandler(w, r)

	//then
	assert.Equal(t, 0, mockProcessor.ProcessCalled)
}
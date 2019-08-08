package mocks

import (
	"pairot/features/pairing"
	"pairot/persistence"
)

var _ pairing.Processor = &MockProcessor{}

var _ persistence.DB = &MockDB{}

type MockProcessor struct {
	ProcessFn     func(input pairing.Input) pairing.SlackResponse
	ProcessCalled int
}

func (p *MockProcessor) Process(input pairing.Input) pairing.SlackResponse {
	p.ProcessCalled++
	return p.ProcessFn(input)
}

type MockDB struct {
	FindTeamByNameFn        func(teamName string) ([]byte, error)
	FindTeamByNameCalled    int
	DecodeFn                func(data []byte, val interface{}) error
	DecodeCalled            int
	UpdateTeamMembersFn     func(teamName string, members interface{}) error
	UpdateTeamMembersCalled int
}

func (db *MockDB) FindTeamByName(teamName string) ([]byte, error) {
	db.FindTeamByNameCalled++
	return db.FindTeamByNameFn(teamName)
}

func (db *MockDB) Decode(data []byte, val interface{}) error {
	db.DecodeCalled++
	return db.DecodeFn(data, val)
}

func (db *MockDB) UpdateTeamMembers(teamName string, members interface{}) error {
	db.UpdateTeamMembersCalled++
	return db.UpdateTeamMembersFn(teamName, members)
}

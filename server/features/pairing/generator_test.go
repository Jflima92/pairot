package pairing_test

import (
	"github.com/stretchr/testify/assert"
	"pairot/features/pairing"
	"testing"
)

func TestGivenATeamWithOnlyLockedReturnCorrectResponse(t *testing.T) {
	//given
	expRes := [][]pairing.Member{

	}
	generator := pairing.NewGenerator()
	//when
	res := generator.Generate([]pairing.Member{}, 1)

	//then
	assert.Equal(t, expRes, res)
}

type Config int

const (
	Locked Config = iota
	Unlocked
	Same
)

func getMembers(isOdd bool, config Config) []pairing.Member {
	names := []string{
		"Jorge",
		"Ruba",
		"Mario",
		"Fede",
		"Linh",
		"Felix",
		"Chinmay",
	}
	var 
	members := []pairing.Member{
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
	}

	if isOdd == true {
		members = append(members, pairing.Member{Name: "Chinmay", Locked: false})
	}
	return members
}

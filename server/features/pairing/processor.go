package pairing

import (
	"math/rand"
	"pairot/persistence"
	"time"
)

type PairProcessor struct {
	db persistence.DB
}

func NewProcessor(db persistence.DB) Processor {
	return &PairProcessor{
		db: db,
	}
}

func (p *PairProcessor) Process(input Input) SlackResponse {
	teamName := input.TeamName
	dbTeam, err := p.db.FindTeamByName(teamName)

	if err != nil {
		return createSlackErrorResponse("Team not found")
	} else {
		var team Team
		err := p.db.Decode(dbTeam, &team)
		if err != nil {
			return createSlackErrorResponse("Error while looking for team")
		}
		teamPairs := p.getTeamPairs(team)
		updatedTeamMembers := p.updateTeam(teamPairs).Members
		updError := p.db.UpdateTeamMembers(teamName, updatedTeamMembers)
		if updError != nil {
			return createSlackErrorResponse("Error updating team")
		}
		return createSlackResponse(teamPairs)
	}
}

func (p *PairProcessor) getTeamPairs(team Team) [][]Member {
	members := team.Members
	shuffledMembers := p.shuffleTeam(members)
	locked, unlocked := p.splitTeam(shuffledMembers)

	ls := len(locked)
	us := len(unlocked)
	var pairs [][]Member
	if ls > us {
		// Locked is bigger than unlocked
		p.generatePairsBasedOnSizes(us, ls, unlocked, locked, &pairs)
	} else if ls < us {
		// Locked is bigger than unlocked
		p.generatePairsBasedOnSizes(ls, us, locked, unlocked, &pairs)
	} else {
		// Both are the same size
		p.generatePairs(us, locked, unlocked, &pairs)
	}

	return pairs
}

func (p *PairProcessor) generatePairsBasedOnSizes(smallestLength int, biggestLength int, smallestSlice []Member, biggestSlice []Member, pairs *[][]Member) {
	// Generate pairs until smallestSlice is over
	p.generatePairs(smallestLength, smallestSlice, biggestSlice, pairs)

	// Continue generating pairs with only the biggest slice if still enough members
	if biggestLength-smallestLength > 1 {
		for u := smallestLength; u < biggestLength; u = u + 2 {
			var pair []Member
			if u == len(biggestSlice)-1 && (u+1)%2 != 0 {
				pair = []Member{biggestSlice[u]}
			} else {
				pair = []Member{biggestSlice[u], biggestSlice[u+1]}
			}
			*pairs = append(*pairs, pair)
		}
	} else {
		*pairs = append(*pairs, []Member{biggestSlice[biggestLength-1]})
	}
}

func (p *PairProcessor) generatePairs(us int, locked []Member, unlocked []Member, pairs *[][]Member) {
	for i := 0; i < us; i++ {
		lm := locked[i]
		um := unlocked[i]
		*pairs = append(*pairs, []Member{lm, um})
	}
}

func (p *PairProcessor) splitTeam(team []Member) (locked []Member, unlocked []Member) {
	var l []Member
	var u []Member
	for _, m := range team {
		if m.Locked {
			l = append(l, m)
		} else {
			u = append(u, m)
		}
	}
	return l, u
}

func (p *PairProcessor) shuffleTeam(members []Member) []Member {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Member, len(members))
	perm := r.Perm(len(members))
	for i, randIndex := range perm {
		ret[i] = members[randIndex]
	}
	return ret
}

func (p *PairProcessor) updateTeam(members [][]Member) Team {
	var team Team
	for i := 0; i < len(members); i++ {
		pair := members[i]
		left := pair[0]
		if len(pair) == 2 {
			right := pair[1]
			r := rand.New(rand.NewSource(time.Now().Unix()))
			if left.Locked == false && right.Locked == false {
				v := r.Intn(2)
				pair[v].Locked = true
			} else {
				left.Locked = !left.Locked
				right.Locked = !right.Locked
			}
			team.Members = append(team.Members, left)
			team.Members = append(team.Members, right)
		} else {
			team.Members = append(team.Members, left)
		}
	}
	return team
}

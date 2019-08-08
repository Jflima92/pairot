package pairing

type PairGenerator struct {
}

func NewGenerator() Generator {
	return &PairGenerator{}
}

func (PairGenerator) Generate(members []Member, seed int64) [][]Member {
	panic("implement me")
}


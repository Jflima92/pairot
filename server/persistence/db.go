package persistence

type DB interface {
	FindTeamByName(teamName string) ([]byte, error)
	UpdateTeamMembers(teamName string, members interface{}) error
	Decode (data []byte, val interface{}) error
}
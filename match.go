package vbratings

type Match struct {
	PlayerAId string
	PlayerBId string
	PlayerCId string
	PlayerDId string
	IsForfeit bool
	Set1      string
	Set2      string
	Set3      string
	Year      int
	Gender    string
}

type MatchRepository interface {
	Create(match Match, id string)
	Find(id string) *Match
	GetAllPlayerIds() []string
	GetAllMatchesByYear(year int) []Match
}

package vbratings

type Match struct {
	Id           string
	PlayerAId    string
	PlayerBId    string
	PlayerCId    string
	PlayerDId    string
	IsForfeit    bool
	Set1         string
	Set2         string
	Set3         string
	TournamentId string
}

type MatchRepository interface {
	Create(match Match)
	Find(id string) *Match
	GetAllPlayerIds() []string
	GetAllMatchesByYear(year int) []Match
}

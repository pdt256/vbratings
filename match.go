package vbratings

type Match struct {
	PlayerAId int
	PlayerBId int
	PlayerCId int
	PlayerDId int
	IsForfeit bool
	Set1      string
	Set2      string
	Set3      string
	Year      int
	Gender    Gender
}

type MatchRepository interface {
	Create(match Match, id string)
	Find(id string) *Match
	GetAllPlayerIds() []int
	GetAllMatchesByYear(year int) []Match
}

type Gender uint

func (gender Gender) String() string {
	names := [...]string{
		"Male",
		"Female",
		"Code",
	}

	return names[gender]
}

const (
	Male Gender = iota
	Female
)

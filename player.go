package vbratings

type Player struct {
	BvbId  int
	Name   string
	ImgUrl string
}

type PlayerRepository interface {
	Create(player Player) error
}

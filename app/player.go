package app

import (
	"github.com/pdt256/vbratings"
)

type Player struct {
	playerRepository vbratings.PlayerRepository
}

func (p *Player) Create(
	bvbId int,
	name string,
	imgUrl string) error {
	player := vbratings.Player{
		BvbId:  bvbId,
		Name:   name,
		ImgUrl: imgUrl,
	}
	return p.playerRepository.Create(player)
}

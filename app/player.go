package app

import (
	"github.com/pdt256/vbratings"
)

type Player struct {
	playerRepository vbratings.PlayerRepository
}

// Create Player
func (p *Player) Create(
	id string,
	name string,
	gender string) error {
	player := vbratings.Player{
		Id:     id,
		Name:   name,
		Gender: gender,
	}
	return p.playerRepository.Create(player)
}

// Get Player by id
func (p *Player) GetPlayer(id string) (*vbratings.Player, error) {
	return p.playerRepository.GetPlayer(id)
}

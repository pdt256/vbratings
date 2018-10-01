package vbratings

import (
	"errors"
)

type Player struct {
	BvbId  int
	Name   string
	ImgUrl string
}

type PlayerRepository interface {
	Create(player Player) error
	GetPlayer(bvbId int) (*Player, error)
}

var PlayerNotFound = errors.New("player not found")

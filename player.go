package vbratings

import (
	"errors"
)

type Player struct {
	Id     string
	Name   string
	Gender string
}

type PlayerRepository interface {
	Create(player Player) error
	GetPlayer(id string) (*Player, error)
}

var PlayerNotFound = errors.New("player not found")

package app

import (
	"github.com/pdt256/vbratings/sqlite"
)

type App struct {
	PlayerRating PlayerRating
	Player       Player
}

func New(configuration *Configuration) *App {
	db := sqlite.NewFileDB(configuration.DbPath)
	playerRatingRepository := sqlite.NewPlayerRatingRepository(db)
	playerRepository := sqlite.NewPlayerRepository(db)

	return &App{
		PlayerRating: PlayerRating{playerRatingRepository},
		Player:       Player{playerRepository},
	}
}

type Configuration struct {
	DbPath string
}

func NewConfiguration(dbPath string) *Configuration {
	return &Configuration{dbPath}
}

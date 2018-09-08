package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pdt256/vbratings"
)

type playerRepository struct {
	db *sql.DB
}

var _ vbratings.PlayerRepository = (*playerRepository)(nil)

func NewPlayerRepository(db *sql.DB) *playerRepository {
	return &playerRepository{db}
}

func (r *playerRepository) InitDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS player (
			bvbId TEXT NOT NULL PRIMARY KEY
			,name TEXT NOT NULL
			,imgUrl TEXT NOT NULL
		);
		DELETE FROM player;`

	_, createError := r.db.Exec(sqlStmt)
	if createError != nil {
		log.Printf("%q: %s\n", createError, sqlStmt)
		return
	}
}

func (r *playerRepository) Create(player vbratings.Player) error {
	_, err := r.db.Exec(
		"INSERT INTO player(bvbId, name, imgUrl) VALUES ($1, $2, $3)",
		player.BvbId,
		player.Name,
		player.ImgUrl,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

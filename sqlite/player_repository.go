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
	r := &playerRepository{db}
	r.migrateDB()
	return r
}

func (r *playerRepository) migrateDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS player (
			id TEXT NOT NULL PRIMARY KEY
			,name TEXT NOT NULL
			,imgUrl TEXT NOT NULL
		);`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *playerRepository) GetPlayer(id string) (*vbratings.Player, error) {
	var p vbratings.Player
	row := r.db.QueryRow("SELECT id, name, imgUrl FROM player WHERE id = $1", id)
	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.ImgUrl,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, vbratings.PlayerNotFound
		}
		checkError(err)
	}

	return &p, nil
}

func (r *playerRepository) Create(player vbratings.Player) error {
	_, err := r.db.Exec(
		"INSERT INTO player(id, name, imgUrl) VALUES ($1, $2, $3)",
		player.Id,
		player.Name,
		player.ImgUrl,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

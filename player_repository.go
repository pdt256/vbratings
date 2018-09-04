package vbscraper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type PlayerRepository interface {
	Create(player Player) error
}

type sqlitePlayerRepository struct {
	db *sql.DB
}

func NewSqlitePlayerRepository(db *sql.DB) *sqlitePlayerRepository {
	return &sqlitePlayerRepository{db}
}

func (r *sqlitePlayerRepository) InitDB() {
	sqlStmt := `CREATE TABLE player (
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

func (r *sqlitePlayerRepository) Create(player Player) error {
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

//func (r *sqlitePlayerRepository) Find(id string) *Player {
//	var p Player
//	row := r.db.QueryRow("SELECT bvbId, name, imgUrl FROM player WHERE id = $1", id)
//	err := row.Scan(
//		&p.BvbId,
//		&p.Name,
//		&p.ImgUrl,
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return &p
//}

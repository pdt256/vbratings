package bvbinfo

import (
	"database/sql"
	"errors"
	"log"
)

type Repository interface {
	GetPlayerId(bvbId int) (string, error)
	AddPlayerId(playerId string, bvbId int) error
}

var PlayerNotFoundError = errors.New("BvbInfo player rating not found")

type sqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *sqliteRepository {
	r := &sqliteRepository{db}
	r.migrateDB()
	return r
}

var _ Repository = (*sqliteRepository)(nil)

func (r *sqliteRepository) migrateDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS bvbinfo_player (
			playerId TEXT NOT NULL
			,bvbId INT NOT NULL
		);`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *sqliteRepository) GetPlayerId(bvbId int) (string, error) {
	var playerId string
	row := r.db.QueryRow("SELECT playerId FROM bvbinfo_player WHERE bvbId = $1", bvbId)
	err := row.Scan(&playerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", PlayerNotFoundError
		}
		checkError(err)
	}

	return playerId, nil
}

func (r *sqliteRepository) AddPlayerId(playerId string, bvbId int) error {
	_, err := r.db.Exec(
		"INSERT INTO bvbinfo_player(playerId, bvbId) VALUES ($1, $2)",
		playerId,
		bvbId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type cacheRepository struct {
	repository Repository
	players    map[int]string
}

var _ Repository = (*cacheRepository)(nil)

func NewCacheRepository(repository Repository) *cacheRepository {
	return &cacheRepository{
		repository: repository,
		players:    make(map[int]string),
	}
}

func (r *cacheRepository) GetPlayerId(bvbId int) (string, error) {
	if playerId, ok := r.players[bvbId]; ok {
		return playerId, nil
	}

	playerId, err := r.repository.GetPlayerId(bvbId)
	if err != nil {
		return "", err
	}

	r.players[bvbId] = playerId

	return playerId, nil
}

func (r *cacheRepository) AddPlayerId(playerId string, bvbId int) error {
	r.players[bvbId] = playerId
	return r.repository.AddPlayerId(playerId, bvbId)
}

func NewRepositoryWithCaching(db *sql.DB) Repository {
	return NewCacheRepository(NewSqliteRepository(db))
}

package cbva

import (
	"database/sql"
	"errors"
	"log"
)

type Repository interface {
	GetPlayerId(name string) (string, error)
	AddPlayerId(playerId string, cbvaName string) error
	AddTournamentId(tournamentid string, cbvaTournamentId string) error
}

var PlayerNotFoundError = errors.New("CBVA player rating not found")

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
	sqlStmt1 := `CREATE TABLE IF NOT EXISTS cbva_player (
			playerId TEXT NOT NULL
			,cbvaName TEXT NOT NULL
		);`

	_, err1 := r.db.Exec(sqlStmt1)
	checkError(err1)

	sqlStmt2 := `CREATE TABLE IF NOT EXISTS cbva_tournament (
			tournamentId TEXT NOT NULL
			,cbvaTournamentId TEXT NOT NULL
		);`

	_, err2 := r.db.Exec(sqlStmt2)
	checkError(err2)
}

func (r *sqliteRepository) GetPlayerId(cbvaName string) (string, error) {
	var playerId string
	row := r.db.QueryRow("SELECT playerId FROM cbva_player WHERE cbvaName = $1", cbvaName)
	err := row.Scan(&playerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", PlayerNotFoundError
		}
		checkError(err)
	}

	return playerId, nil
}

func (r *sqliteRepository) AddPlayerId(playerId string, cbvaName string) error {
	_, err := r.db.Exec(
		"INSERT INTO cbva_player(playerId, cbvaName) VALUES ($1, $2)",
		playerId,
		cbvaName,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteRepository) AddTournamentId(tournamentid string, cbvaTournamentId string) error {
	_, err := r.db.Exec(
		"INSERT INTO cbva_tournament(tournamentId, cbvaTournamentId) VALUES ($1, $2)",
		tournamentid,
		cbvaTournamentId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type cacheRepository struct {
	repository Repository
	players    map[string]string
}

var _ Repository = (*cacheRepository)(nil)

func NewCacheRepository(repository Repository) *cacheRepository {
	return &cacheRepository{
		repository: repository,
		players:    make(map[string]string),
	}
}

func (r *cacheRepository) GetPlayerId(cbvaName string) (string, error) {
	if playerId, ok := r.players[cbvaName]; ok {
		return playerId, nil
	}

	playerId, err := r.repository.GetPlayerId(cbvaName)
	if err != nil {
		return "", err
	}

	r.players[cbvaName] = playerId

	return playerId, nil
}

func (r *cacheRepository) AddPlayerId(playerId string, cbvaName string) error {
	r.players[cbvaName] = playerId
	return r.repository.AddPlayerId(playerId, cbvaName)
}

func (r *cacheRepository) AddTournamentId(tournamentId string, cbvaTournamentId string) error {
	return r.repository.AddTournamentId(tournamentId, cbvaTournamentId)
}

func NewRepositoryWithCaching(db *sql.DB) Repository {
	return NewCacheRepository(NewSqliteRepository(db))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

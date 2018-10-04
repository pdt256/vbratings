package cbva

import (
	"database/sql"
	"errors"
	"log"
)

type Repository interface {
	GetPlayerId(name string) (string, error)
	AddPlayerId(playerId string, cbvaName string) error
	AddTournament(cbvaTournament Tournament) error
}

var PlayerNotFoundError = errors.New("CBVA player not found")
var TournamentNotFoundError = errors.New("CBVA tournament not found")

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
			id TEXT NOT NULL PRIMARY KEY
			,date TEXT NOT NULL
			,rating TEXT NOT NULL
			,gender TEXT NOT NULL
			,location TEXT NOT NULL
			,tournamentId TEXT NOT NULL
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

func (r *sqliteRepository) AddTournament(tournament Tournament) error {
	_, err := r.db.Exec(
		"INSERT INTO cbva_tournament(id, date, rating, gender, location, tournamentId) VALUES ($1, $2, $3, $4, $5, $6)",
		tournament.Id,
		tournament.Date,
		tournament.Rating,
		tournament.Gender,
		tournament.Location,
		tournament.TournamentId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteRepository) GetTournament(id string) (*Tournament, error) {
	var t Tournament
	row := r.db.QueryRow("SELECT id, date, rating, gender, location, tournamentId FROM cbva_tournament WHERE id = $1", id)
	err := row.Scan(
		&t.Id,
		&t.Date,
		&t.Rating,
		&t.Gender,
		&t.Location,
		&t.TournamentId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, TournamentNotFoundError
		}
		checkError(err)
	}

	return &t, nil
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

func (r *cacheRepository) AddTournament(tournament Tournament) error {
	return r.repository.AddTournament(tournament)
}

func NewRepositoryWithCaching(db *sql.DB) Repository {
	return NewCacheRepository(NewSqliteRepository(db))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

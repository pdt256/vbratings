package bvbinfo

import (
	"database/sql"
	"errors"
	"log"
)

type Repository interface {
	GetPlayer(id int) (*Player, error)
	GetPlayerId(id int) (string, error)
	AddPlayer(player Player) error
	GetTournament(id int) (*Tournament, error)
	AddTournament(tournament Tournament) error
}

var PlayerNotFoundError = errors.New("BvbInfo player not found")
var TournamentNotFoundError = errors.New("BvbInfo tournament not found")

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
	sqlStmt1 := `CREATE TABLE IF NOT EXISTS bvbinfo_player (
			id INT NOT NULL PRIMARY KEY
			,name TEXT NOT NULL
			,imgUrl TEXT NOT NULL
			,playerId TEXT NOT NULL
		);`
	_, err1 := r.db.Exec(sqlStmt1)
	checkError(err1)

	sqlStmt2 := `CREATE TABLE IF NOT EXISTS bvbinfo_tournament (
			id TEXT NOT NULL
			,name TEXT NOT NULL
			,year TEXT NOT NULL
			,dates TEXT NOT NULL
			,gender TEXT NOT NULL
			,tournamentId TEXT NOT NULL
		);`

	_, err2 := r.db.Exec(sqlStmt2)
	checkError(err2)
}

func (r *sqliteRepository) GetPlayerId(id int) (string, error) {
	var playerId string
	row := r.db.QueryRow("SELECT playerId FROM bvbinfo_player WHERE id = $1", id)
	err := row.Scan(&playerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", PlayerNotFoundError
		}
		checkError(err)
	}

	return playerId, nil
}

func (r *sqliteRepository) GetPlayer(id int) (*Player, error) {
	var p Player
	row := r.db.QueryRow("SELECT id, name, imgUrl, playerId FROM bvbinfo_player WHERE id = $1", id)
	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.ImgUrl,
		&p.PlayerId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, PlayerNotFoundError
		}
		checkError(err)
	}

	return &p, nil
}

func (r *sqliteRepository) AddPlayer(player Player) error {
	_, err := r.db.Exec(
		"INSERT INTO bvbinfo_player(id, name, imgUrl, playerId) VALUES ($1, $2, $3, $4)",
		player.Id,
		player.Name,
		player.ImgUrl,
		player.PlayerId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteRepository) AddTournament(tournament Tournament) error {
	_, err := r.db.Exec(
		"INSERT INTO bvbinfo_tournament(id, name, year, dates, gender, tournamentId ) VALUES ($1, $2, $3, $4, $5, $6)",
		tournament.Id,
		tournament.Name,
		tournament.Year,
		tournament.Dates,
		tournament.Gender,
		tournament.TournamentId,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *sqliteRepository) GetTournament(id int) (*Tournament, error) {
	var t Tournament
	row := r.db.QueryRow("SELECT id, name, year, dates, gender, tournamentId FROM bvbinfo_tournament WHERE id = $1", id)
	err := row.Scan(
		&t.Id,
		&t.Name,
		&t.Year,
		&t.Dates,
		&t.Gender,
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
	players    map[int]string
}

var _ Repository = (*cacheRepository)(nil)

func NewCacheRepository(repository Repository) *cacheRepository {
	return &cacheRepository{
		repository: repository,
		players:    make(map[int]string),
	}
}

func (r *cacheRepository) GetPlayer(bvbId int) (*Player, error) {
	return r.repository.GetPlayer(bvbId)
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

func (r *cacheRepository) AddTournament(tournament Tournament) error {
	return r.repository.AddTournament(tournament)
}

func (r *cacheRepository) GetTournament(id int) (*Tournament, error) {
	return r.repository.GetTournament(id)
}

func (r *cacheRepository) AddPlayer(player Player) error {
	r.players[player.Id] = player.PlayerId
	return r.repository.AddPlayer(player)
}

func NewRepositoryWithCaching(db *sql.DB) Repository {
	return NewCacheRepository(NewSqliteRepository(db))
}

package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pdt256/vbratings"
)

type tournamentRepository struct {
	db *sql.DB
}

var _ vbratings.TournamentRepository = (*tournamentRepository)(nil)

func NewTournamentRepository(db *sql.DB) *tournamentRepository {
	return &tournamentRepository{db}
}

func (r *tournamentRepository) InitDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS tournament_result (
			id TEXT NOT NULL PRIMARY KEY
			,player1Name TEXT NOT NULL
			,player2Name TEXT NOT NULL
			,earnedFinish INT NOT NULL
		);
		DELETE FROM tournament_result;`

	_, createError := r.db.Exec(sqlStmt)
	if createError != nil {
		log.Printf("%q: %s\n", createError, sqlStmt)
		return
	}
}

func (r *tournamentRepository) AddTournamentResult(tournamentResult vbratings.TournamentResult, id string) {
	_, err := r.db.Exec(
		"INSERT INTO tournament_result(id, player1Name, player2Name, earnedFinish) VALUES ($1, $2, $3, $4)",
		id,
		tournamentResult.Player1Name,
		tournamentResult.Player2Name,
		tournamentResult.EarnedFinish,
	)
	checkError(err)
}

func (r *tournamentRepository) GetAllTournamentResults() []vbratings.TournamentResult {
	var tournamentResults []vbratings.TournamentResult

	rows, queryErr := r.db.Query("SELECT player1Name, player2Name, earnedFinish FROM tournament_result")
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var tr vbratings.TournamentResult
		checkError(rows.Scan(
			&tr.Player1Name,
			&tr.Player2Name,
			&tr.EarnedFinish,
		))

		tournamentResults = append(tournamentResults, tr)
	}
	checkError(rows.Err())

	return tournamentResults
}

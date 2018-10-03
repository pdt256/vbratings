package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pdt256/vbratings"
)

type tournamentRepository struct {
	db *sql.DB
}

var _ vbratings.TournamentRepository = (*tournamentRepository)(nil)

func NewTournamentRepository(db *sql.DB) *tournamentRepository {
	r := &tournamentRepository{db}
	r.migrateDB()
	return r
}

func (r *tournamentRepository) migrateDB() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS tournament_result (
			id TEXT NOT NULL PRIMARY KEY
			,player1Id TEXT NOT NULL
			,player2Id TEXT NOT NULL
			,earnedFinish INT NOT NULL
			,tournamentId TEXT NOT NULL
		);`

	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)
}

func (r *tournamentRepository) AddTournamentResult(tournamentResult vbratings.TournamentResult) {
	_, err := r.db.Exec(
		"INSERT INTO tournament_result(id, player1Id, player2Id, earnedFinish, tournamentId) VALUES ($1, $2, $3, $4, $5)",
		tournamentResult.Id,
		tournamentResult.Player1Id,
		tournamentResult.Player2Id,
		tournamentResult.EarnedFinish,
		tournamentResult.TournamentId,
	)
	checkError(err)
}

func (r *tournamentRepository) GetAllTournamentResults() []vbratings.TournamentResult {
	var tournamentResults []vbratings.TournamentResult

	rows, queryErr := r.db.Query("SELECT id, player1Id, player2Id, earnedFinish, tournamentId FROM tournament_result")
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var tr vbratings.TournamentResult
		checkError(rows.Scan(
			&tr.Id,
			&tr.Player1Id,
			&tr.Player2Id,
			&tr.EarnedFinish,
			&tr.TournamentId,
		))

		tournamentResults = append(tournamentResults, tr)
	}
	checkError(rows.Err())

	return tournamentResults
}

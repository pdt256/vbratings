package sqlite

import (
	"database/sql"
	"strings"

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
	sqlStmt1 := `CREATE TABLE IF NOT EXISTS tournament_result (
			id TEXT NOT NULL PRIMARY KEY
			,player1Id TEXT NOT NULL
			,player2Id TEXT NOT NULL
			,earnedFinish INT NOT NULL
			,tournamentId TEXT NOT NULL
		);`

	_, err1 := r.db.Exec(sqlStmt1)
	checkError(err1)

	sqlStmt2 := `CREATE TABLE IF NOT EXISTS tournament (
			id TEXT NOT NULL PRIMARY KEY
			,name TEXT NOT NULL
			,date TEXT NOT NULL
			,gender TEXT NOT NULL
			,year INT NOT NULL
		);`

	_, err2 := r.db.Exec(sqlStmt2)
	checkError(err2)
}

func (r *tournamentRepository) executeStatementSafe(sqlStmt string) {
	_, createError := r.db.Exec(sqlStmt)
	checkError(createError)

}

func (r *tournamentRepository) Create(tournament vbratings.Tournament) {
	_, err := r.db.Exec(
		"INSERT INTO tournament(id, name, date, gender, year) VALUES ($1, $2, $3, $4, $5)",
		tournament.Id,
		tournament.Name,
		tournament.Date,
		tournament.Gender,
		tournament.Year,
	)
	checkError(err)
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

func (r *tournamentRepository) GetAllTournamentResults() []*vbratings.TournamentResult {
	var tournamentResults []*vbratings.TournamentResult

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

		tournamentResults = append(tournamentResults, &tr)
	}
	checkError(rows.Err())

	return tournamentResults
}

func (r *tournamentRepository) GetAllTournamentsAndResultsByYear(year int) []*vbratings.TournamentAndResults {

	tournaments := r.getTournamentsByYear(year)
	tournamentIds := getTournamentIds(tournaments)
	tournamentResults := r.getTournamentResultsByTournamentIds(tournamentIds)

	tournamentsAndResults := make([]*vbratings.TournamentAndResults, 0, len(tournaments))

	for _, tournament := range tournaments {
		tournamentsAndResults = append(
			tournamentsAndResults,
			&vbratings.TournamentAndResults{
				Tournament: tournament,
				Results:    tournamentResults[tournament.Id],
			},
		)
	}

	return tournamentsAndResults
}

func getTournamentIds(tournaments []*vbratings.Tournament) []string {
	var ids []string

	for _, value := range tournaments {
		ids = append(ids, value.Id)
	}

	return ids
}

func (r *tournamentRepository) GetTournament(id string) (*vbratings.Tournament, error) {
	var t vbratings.Tournament
	row := r.db.QueryRow("SELECT id, name, date, gender, year FROM tournament WHERE id = $1", id)
	err := row.Scan(
		&t.Id,
		&t.Name,
		&t.Date,
		&t.Gender,
		&t.Year,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, vbratings.TournamentNotFound
		}
		checkError(err)
	}

	return &t, nil
}

func (r *tournamentRepository) getTournamentsByYear(i int) []*vbratings.Tournament {
	var tournaments []*vbratings.Tournament

	rows, queryErr := r.db.Query("SELECT id, name, date, gender, year FROM tournament")
	checkError(queryErr)

	defer rows.Close()

	for rows.Next() {
		var tr vbratings.Tournament
		checkError(rows.Scan(
			&tr.Id,
			&tr.Name,
			&tr.Date,
			&tr.Gender,
			&tr.Year,
		))

		tournaments = append(tournaments, &tr)
	}
	checkError(rows.Err())

	return tournaments
}

func (r *tournamentRepository) getTournamentResultsByTournamentIds(tournamentIds []string) map[string][]*vbratings.TournamentResult {
	tournamentsResults := make(map[string][]*vbratings.TournamentResult, len(tournamentIds))

	rows, queryErr := r.db.Query("SELECT id, player1Id, player2Id, earnedFinish, tournamentId FROM tournament_result WHERE tournamentId IN ($1) ORDER BY tournamentId, earnedFinish", getUnsafeSQLStrings(tournamentIds))
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

		tournamentsResults[tr.TournamentId] = append(
			tournamentsResults[tr.TournamentId],
			&tr,
		)
	}
	checkError(rows.Err())

	return tournamentsResults
}

func getUnsafeSQLStrings(ids []string) string {
	return strings.Join(ids, `","`)
}

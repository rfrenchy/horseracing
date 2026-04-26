package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const baseElo = 1500.0

func main() {
	// 1. Connect to Postgres
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	// 2. Create atp_match_features table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS atp_match_features (
		tourney_id TEXT,
		match_num INT,
		surface TEXT,
		p1_id INT,
		p2_id INT,
		p1_elo FLOAT,
		p2_elo FLOAT,
		elo_delta FLOAT,
		p1_won BOOLEAN,
		PRIMARY KEY (tourney_id, match_num)
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal().Err(err).Msg("Failed to create atp_match_features table")
	}
	log.Info().Msg("Table atp_match_features ready")

	// 3. Query all matches joined with elos
	query := `
		SELECT 
			m.tourney_id, m.match_num, m.surface, m.winner_id, m.loser_id,
			e.winner_elo_clay, e.winner_elo_hard, e.winner_elo_grass,
			e.loser_elo_clay, e.loser_elo_hard, e.loser_elo_grass
		FROM atp_matches_raw m
		JOIN atp_match_elos e ON m.tourney_id = e.tourney_id AND m.match_num = e.match_num
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query matches and elos")
	}
	defer rows.Close()

	// Pre-compile insert statement
	insertSQL := `
		INSERT INTO atp_match_features (
			tourney_id, match_num, surface,
			p1_id, p2_id, p1_elo, p2_elo, elo_delta, p1_won
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) ON CONFLICT (tourney_id, match_num) DO NOTHING
	`
	tx, err := db.Begin()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to begin transaction")
	}

	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to prepare statement")
	}
	defer stmt.Close()

	processed := 0
	skipped := 0
	for rows.Next() {
		var tourneyID, surface string
		var matchNum, winnerID, loserID int
		var wClay, wHard, wGrass float64
		var lClay, lHard, lGrass float64

		if err := rows.Scan(
			&tourneyID, &matchNum, &surface, &winnerID, &loserID,
			&wClay, &wHard, &wGrass,
			&lClay, &lHard, &lGrass,
		); err != nil {
			log.Fatal().Err(err).Msg("Failed to scan row")
		}

		// Determine relevant Elo based on surface
		var wElo, lElo float64
		switch surface {
		case "Clay":
			wElo = wClay
			lElo = lClay
		case "Hard":
			wElo = wHard
			lElo = lHard
		case "Grass":
			wElo = wGrass
			lElo = lGrass
		default:
			// For Carpet or missing surface, we'll skip or use base elo.
			// Let's use base elo 1500 for non-core surfaces for now, but mark skipped.
			skipped++
			continue
		}

		// Calculate P1 and P2
		var p1ID, p2ID int
		var p1Elo, p2Elo float64
		var p1Won bool

		if winnerID < loserID {
			p1ID = winnerID
			p2ID = loserID
			p1Elo = wElo
			p2Elo = lElo
			p1Won = true
		} else {
			p1ID = loserID
			p2ID = winnerID
			p1Elo = lElo
			p2Elo = wElo
			p1Won = false
		}

		eloDelta := p1Elo - p2Elo

		_, err := stmt.Exec(tourneyID, matchNum, surface, p1ID, p2ID, p1Elo, p2Elo, eloDelta, p1Won)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert features")
		}

		processed++
	}

	if err := rows.Err(); err != nil {
		tx.Rollback()
		log.Fatal().Err(err).Msg("Row iteration error")
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	log.Info().
		Int("processed", processed).
		Int("skipped", skipped).
		Msg("Successfully computed and saved features")
}

package main

import (
	"database/sql"
	"os"

	"punts/internal/tennis"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	baseElo = 1500.0
	kFactor = 32.0
)

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

	// 2. Create atp_match_elos table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS atp_match_elos (
		tourney_id TEXT,
		match_num INT,
		winner_elo_clay FLOAT,
		winner_elo_hard FLOAT,
		winner_elo_grass FLOAT,
		loser_elo_clay FLOAT,
		loser_elo_hard FLOAT,
		loser_elo_grass FLOAT,
		PRIMARY KEY (tourney_id, match_num)
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal().Err(err).Msg("Failed to create atp_match_elos table")
	}
	log.Info().Msg("Table atp_match_elos ready")

	// 3. Query all matches
	query := `
		SELECT tourney_id, match_num, tourney_date, winner_id, loser_id, surface
		FROM atp_matches
		ORDER BY tourney_date ASC, match_num ASC
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query matches")
	}
	defer rows.Close()

	// 4. Calculate Elos
	// Map of Player ID -> Surface -> Elo
	playerElos := make(map[int]map[string]float64)

	getElo := func(playerID int, surface string) float64 {
		if playerElos[playerID] == nil {
			playerElos[playerID] = make(map[string]float64)
		}
		if elo, exists := playerElos[playerID][surface]; exists {
			return elo
		}
		playerElos[playerID][surface] = baseElo
		return baseElo
	}

	setElo := func(playerID int, surface string, elo float64) {
		if playerElos[playerID] == nil {
			playerElos[playerID] = make(map[string]float64)
		}
		playerElos[playerID][surface] = elo
	}

	// Pre-compile insert statement
	insertSQL := `
		INSERT INTO atp_match_elos (
			tourney_id, match_num,
			winner_elo_clay, winner_elo_hard, winner_elo_grass,
			loser_elo_clay, loser_elo_hard, loser_elo_grass
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
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
	for rows.Next() {
		var tourneyID string
		var matchNum, tourneyDate, winnerID, loserID int
		var surface string

		if err := rows.Scan(&tourneyID, &matchNum, &tourneyDate, &winnerID, &loserID, &surface); err != nil {
			log.Fatal().Err(err).Msg("Failed to scan row")
		}

		// Pre-match elos
		wClay := getElo(winnerID, "Clay")
		wHard := getElo(winnerID, "Hard")
		wGrass := getElo(winnerID, "Grass")
		lClay := getElo(loserID, "Clay")
		lHard := getElo(loserID, "Hard")
		lGrass := getElo(loserID, "Grass")

		// Insert pre-match elos
		_, err := stmt.Exec(tourneyID, matchNum, wClay, wHard, wGrass, lClay, lHard, lGrass)
		if err != nil {
			tx.Rollback()
			log.Fatal().Err(err).Msg("Failed to insert elos")
		}

		// Update elo for the played surface
		if surface == "Clay" || surface == "Hard" || surface == "Grass" {
			wElo := getElo(winnerID, surface)
			lElo := getElo(loserID, surface)

			newWElo, newLElo := tennis.ProcessMatch(wElo, lElo, kFactor)
			setElo(winnerID, surface, newWElo)
			setElo(loserID, surface, newLElo)
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

	log.Info().Int("matches_processed", processed).Msg("Successfully computed and saved Elos")
}

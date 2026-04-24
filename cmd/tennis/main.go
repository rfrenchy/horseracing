package main

import (
	"database/sql"
	"os"

	"punts/internal/tennis"

	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	// 1. Connect to Postgres
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	// 2. Create table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS atp_matches (
		tourney_id TEXT,
		tourney_name TEXT,
		surface TEXT,
		draw_size INT,
		tourney_level TEXT,
		tourney_date INT,
		match_num INT,
		winner_id INT,
		winner_seed TEXT,
		winner_entry TEXT,
		winner_name TEXT,
		winner_hand TEXT,
		winner_ht INT,
		winner_ioc TEXT,
		winner_age FLOAT,
		loser_id INT,
		loser_seed TEXT,
		loser_entry TEXT,
		loser_name TEXT,
		loser_hand TEXT,
		loser_ht INT,
		loser_ioc TEXT,
		loser_age FLOAT,
		score TEXT,
		best_of INT,
		round TEXT,
		minutes INT,
		w_ace INT,
		w_df INT,
		w_svpt INT,
		w_1stIn INT,
		w_1stWon INT,
		w_2ndWon INT,
		w_SvGms INT,
		w_bpSaved INT,
		w_bpFaced INT,
		l_ace INT,
		l_df INT,
		l_svpt INT,
		l_1stIn INT,
		l_1stWon INT,
		l_2ndWon INT,
		l_SvGms INT,
		l_bpSaved INT,
		l_bpFaced INT,
		winner_rank INT,
		winner_rank_points INT,
		loser_rank INT,
		loser_rank_points INT,
		PRIMARY KEY (tourney_id, match_num)
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table")
	}
	log.Info().Msg("Table atp_matches ready")

	// 3. Read CSV
	file, err := os.Open("data/atp_matches_2024.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open CSV file")
	}
	defer file.Close()

	var matches []tennis.Match
	if err := gocsv.UnmarshalFile(file, &matches); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse CSV")
	}
	log.Info().Int("count", len(matches)).Msg("Parsed matches")

	// 4. Insert Data
	insertSQL := `
	INSERT INTO atp_matches (
		tourney_id, tourney_name, surface, draw_size, tourney_level, tourney_date, match_num,
		winner_id, winner_seed, winner_entry, winner_name, winner_hand, winner_ht, winner_ioc, winner_age,
		loser_id, loser_seed, loser_entry, loser_name, loser_hand, loser_ht, loser_ioc, loser_age,
		score, best_of, round, minutes,
		w_ace, w_df, w_svpt, w_1stIn, w_1stWon, w_2ndWon, w_SvGms, w_bpSaved, w_bpFaced,
		l_ace, l_df, l_svpt, l_1stIn, l_1stWon, l_2ndWon, l_SvGms, l_bpSaved, l_bpFaced,
		winner_rank, winner_rank_points, loser_rank, loser_rank_points
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7,
		$8, $9, $10, $11, $12, $13, $14, $15,
		$16, $17, $18, $19, $20, $21, $22, $23,
		$24, $25, $26, $27,
		$28, $29, $30, $31, $32, $33, $34, $35, $36,
		$37, $38, $39, $40, $41, $42, $43, $44, $45,
		$46, $47, $48, $49
	) ON CONFLICT (tourney_id, match_num) DO NOTHING;`

	tx, err := db.Begin()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction")
	}
	
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to prepare statement")
	}
	defer stmt.Close()

	inserted := 0
	for _, m := range matches {
		_, err := stmt.Exec(
			m.TourneyID, m.TourneyName, m.Surface, m.DrawSize, m.TourneyLevel, m.TourneyDate, m.MatchNum,
			m.WinnerID, m.WinnerSeed, m.WinnerEntry, m.WinnerName, m.WinnerHand, m.WinnerHt, m.WinnerIOC, m.WinnerAge,
			m.LoserID, m.LoserSeed, m.LoserEntry, m.LoserName, m.LoserHand, m.LoserHt, m.LoserIOC, m.LoserAge,
			m.Score, m.BestOf, m.Round, m.Minutes,
			m.WAce, m.WDf, m.WSvpt, m.W1stIn, m.W1stWon, m.W2ndWon, m.WSvGms, m.WBpSaved, m.WBpFaced,
			m.LAce, m.LDf, m.LSvpt, m.L1stIn, m.L1stWon, m.L2ndWon, m.LSvGms, m.LBpSaved, m.LBpFaced,
			m.WinnerRank, m.WinnerRankPoints, m.LoserRank, m.LoserRankPoints,
		)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to insert match %s - %d", m.TourneyID, m.MatchNum)
			tx.Rollback()
			os.Exit(1)
		}
		inserted++
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	log.Info().Int("inserted", inserted).Msg("Successfully loaded data into database")
}

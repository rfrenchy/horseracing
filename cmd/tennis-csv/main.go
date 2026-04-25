package main

import (
	"database/sql"
	"os"

	"punts/internal/tennis"

	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
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

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "create-raw",

				Action: func(c *cli.Context) error {
					return createRaw(db)
				},
			},
			{
				Name: "create-tournaments",
				Action: func(c *cli.Context) error {
					return createTournamentsTable(db)
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run app")
	}

}

func createRaw(db *sql.DB) error {
	// 2. Create table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS atp_matches_raw (
		id SERIAL PRIMARY KEY,
		tourney_id TEXT,
		tourney_name TEXT,
		surface TEXT,
		draw_size INT,
		tourney_level TEXT,
		tourney_date TIMESTAMP,
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
		loser_rank_points INT		
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table")
	}
	log.Info().Msg("Table atp_matches_raw ready")

	csvs := [...]string{"atp_matches_2000.csv", "atp_matches_2001.csv", "atp_matches_2002.csv", "atp_matches_2003.csv",
		"atp_matches_2004.csv", "atp_matches_2005.csv", "atp_matches_2006.csv", "atp_matches_2007.csv", "atp_matches_2008.csv",
		"atp_matches_2009.csv", "atp_matches_2010.csv", "atp_matches_2011.csv", "atp_matches_2012.csv", "atp_matches_2013.csv",
		"atp_matches_2014.csv", "atp_matches_2015.csv", "atp_matches_2016.csv", "atp_matches_2017.csv", "atp_matches_2018.csv",
		"atp_matches_2019.csv", "atp_matches_2020.csv", "atp_matches_2021.csv", "atp_matches_2022.csv", "atp_matches_2023.csv",
		"atp_matches_2024.csv"}

	for _, csv := range csvs {
		// 3. Read CSV
		file, err := os.Open("D:\\dev\\horseracing\\data\\" + csv)
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
		INSERT INTO atp_matches_raw (
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
		) ON CONFLICT (id) DO NOTHING;`

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

	return nil
}

func createTournamentsTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS atp_tournaments (
		tourney_id TEXT PRIMARY KEY,
		tourney_name TEXT,
		surface TEXT,
		draw_size INT,
		tourney_level TEXT,
		tourney_date TIMESTAMP
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create table")
		return err
	}
	log.Info().Msg("Table atp_tournaments ready")

	// read tournament csv

	file, err := os.Open("D:\\dev\\horseracing\\data\\atp_tournaments.csv")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshall tournament csv data")
	}
	defer file.Close()

	var tournaments []tennis.Tournament
	if err := gocsv.UnmarshalFile(file, &tournaments); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse tournament CSV")
	}
	log.Info().Int("count", len(tournaments)).Msg("Parsed tournaments")

	insertSQL := `
	INSERT INTO atp_tournaments (
		tourney_id, tourney_name, surface, draw_size, tourney_level, tourney_date
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) ON CONFLICT (tourney_id) DO NOTHING;`

	tx, err := db.Begin()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start transaction")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to prepare statement")
	}
	defer stmt.Close()

	inserted := 0
	for _, t := range tournaments {
		_, err := stmt.Exec(
			t.TourneyID, t.TourneyName, t.Surface, t.DrawSize, t.TourneyLevel, t.TourneyDate,
		)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to insert tournament %s", t.TourneyID)
			os.Exit(1)
		}
		inserted++
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}

	log.Info().Int("inserted", inserted).Msg("Successfully loaded data into database")

	return nil
}

package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Entry struct {
	P1Elo     float64 `json:"p1_elo"`
	TourneyID string  `json:"tourney_id"`
	MatchNum  int     `json:"match_num"`
}

func main() {
	router := gin.Default()

	// Define a GET endpoint
	router.GET("/elos/:player", func(c *gin.Context) {
		playerid := c.Params.ByName("player")

		// search atp match features for player id
		query := `
			SELECT p1_elo, tourney_id, match_num
			FROM atp_match_features
			WHERE p1_id = ` + playerid
		// join winner_id and loser_id to add player name

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

		rows, err := db.Query(query)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to query matches and elos")
		}
		defer rows.Close()

		entries := make([]Entry, 0)

		for rows.Next() {
			var ent Entry

			err = rows.Scan(&ent.P1Elo, &ent.TourneyID, &ent.MatchNum)

			if err != nil {
				log.Fatal().Err(err).Msg("Failed to scan row")
			}

			entries = append(entries, ent)
		}

		c.JSON(http.StatusOK, gin.H{"entries": entries})
	})

	// Start the server
	router.Run(":8080")
}

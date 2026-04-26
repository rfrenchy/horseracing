package tennis

// Match represents a single tennis match record parsed from the ATP CSV data
// and mapped to database columns.
type Match struct {
	ID           string
	TourneyID    string `csv:"tourney_id" db:"tourney_id"`
	TourneyName  string `csv:"tourney_name" db:"tourney_name"`
	Surface      string `csv:"surface" db:"surface"`
	DrawSize     int    `csv:"draw_size" db:"draw_size"`
	TourneyLevel string `csv:"tourney_level" db:"tourney_level"`
	TourneyDate  int    `csv:"tourney_date" db:"tourney_date"`
	MatchNum     int    `csv:"match_num" db:"match_num"`

	WinnerID    int      `csv:"winner_id" db:"winner_id"`
	WinnerSeed  *string  `csv:"winner_seed" db:"winner_seed"` // Pointers handle empty/null CSV fields natively
	WinnerEntry *string  `csv:"winner_entry" db:"winner_entry"`
	WinnerName  string   `csv:"winner_name" db:"winner_name"`
	WinnerHand  string   `csv:"winner_hand" db:"winner_hand"`
	WinnerHt    *int     `csv:"winner_ht" db:"winner_ht"`
	WinnerIOC   string   `csv:"winner_ioc" db:"winner_ioc"`
	WinnerAge   *float64 `csv:"winner_age" db:"winner_age"`

	LoserID    int      `csv:"loser_id" db:"loser_id"`
	LoserSeed  *string  `csv:"loser_seed" db:"loser_seed"`
	LoserEntry *string  `csv:"loser_entry" db:"loser_entry"`
	LoserName  string   `csv:"loser_name" db:"loser_name"`
	LoserHand  string   `csv:"loser_hand" db:"loser_hand"`
	LoserHt    *int     `csv:"loser_ht" db:"loser_ht"`
	LoserIOC   string   `csv:"loser_ioc" db:"loser_ioc"`
	LoserAge   *float64 `csv:"loser_age" db:"loser_age"`

	Score   string `csv:"score" db:"score"`
	BestOf  int    `csv:"best_of" db:"best_of"`
	Round   string `csv:"round" db:"round"`
	Minutes *int   `csv:"minutes" db:"minutes"`

	// Winner Match Stats
	WAce     *int `csv:"w_ace" db:"w_ace"`
	WDf      *int `csv:"w_df" db:"w_df"`
	WSvpt    *int `csv:"w_svpt" db:"w_svpt"`
	W1stIn   *int `csv:"w_1stIn" db:"w_1stIn"`
	W1stWon  *int `csv:"w_1stWon" db:"w_1stWon"`
	W2ndWon  *int `csv:"w_2ndWon" db:"w_2ndWon"`
	WSvGms   *int `csv:"w_SvGms" db:"w_SvGms"`
	WBpSaved *int `csv:"w_bpSaved" db:"w_bpSaved"`
	WBpFaced *int `csv:"w_bpFaced" db:"w_bpFaced"`

	// Loser Match Stats
	LAce     *int `csv:"l_ace" db:"l_ace"`
	LDf      *int `csv:"l_df" db:"l_df"`
	LSvpt    *int `csv:"l_svpt" db:"l_svpt"`
	L1stIn   *int `csv:"l_1stIn" db:"l_1stIn"`
	L1stWon  *int `csv:"l_1stWon" db:"l_1stWon"`
	L2ndWon  *int `csv:"l_2ndWon" db:"l_2ndWon"`
	LSvGms   *int `csv:"l_SvGms" db:"l_SvGms"`
	LBpSaved *int `csv:"l_bpSaved" db:"l_bpSaved"`
	LBpFaced *int `csv:"l_bpFaced" db:"l_bpFaced"`

	// Rankings
	WinnerRank       *int `csv:"winner_rank" db:"winner_rank"`
	WinnerRankPoints *int `csv:"winner_rank_points" db:"winner_rank_points"`
	LoserRank        *int `csv:"loser_rank" db:"loser_rank"`
	LoserRankPoints  *int `csv:"loser_rank_points" db:"loser_rank_points"`
}

type Tournament struct {
	TourneyID    string `csv:"tourney_id" db:"tourney_id"`
	TourneyName  string `csv:"tourney_name" db:"tourney_name"`
	Surface      string `csv:"surface" db:"surface"`
	DrawSize     int    `csv:"draw_size" db:"draw_size"`
	TourneyLevel string `csv:"tourney_level" db:"tourney_level"`
	TourneyDate  string `csv:"tourney_date" db:"tourney_date"`
}

type Player struct {
	ID int `db:"player_id"`
	Name string `csv:"name" db:"player_name"`
	Country string `csv:"country" db:"country"` 
}
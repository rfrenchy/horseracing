package tennis

import "math"

// ExpectedScore calculates the expected score (win probability) for a player.
func ExpectedScore(playerRating, opponentRating float64) float64 {
	return 1.0 / (1.0 + math.Pow(10.0, (opponentRating-playerRating)/400.0))
}

// UpdateRating calculates the new rating for a player.
// actualScore is 1.0 for a win, 0.0 for a loss.
func UpdateRating(rating, expectedScore, actualScore, kFactor float64) float64 {
	return rating + kFactor*(actualScore-expectedScore)
}

// ProcessMatch calculates the new ratings for a winner and loser.
func ProcessMatch(winnerRating, loserRating, kFactor float64) (float64, float64) {
	expectedWinner := ExpectedScore(winnerRating, loserRating)
	expectedLoser := ExpectedScore(loserRating, winnerRating)

	newWinnerRating := UpdateRating(winnerRating, expectedWinner, 1.0, kFactor)
	newLoserRating := UpdateRating(loserRating, expectedLoser, 0.0, kFactor)

	return newWinnerRating, newLoserRating
}

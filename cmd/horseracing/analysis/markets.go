package main

import (
	hrace "punts/internal/horseracing"
)

// smarkets

// Markets helps in defining winner/loser class labels for particular markets
type Markets struct{}

type MarketResult struct{}

// ToWin
// Given racecard?
func (m *Markets) ToWin(r *hrace.Runner) bool {
	return 1 == r.Position
}

func (m *Markets) ToPlaceTop3(r *hrace.Runner) bool {
	return 3 <= r.Position
}

func (m *Markets) ToPlaceTop4(r *hrace.Runner) bool {
	return 4 <= r.Position
}

func (m *Markets) Against(marketresult bool) bool {
	return !marketresult
}

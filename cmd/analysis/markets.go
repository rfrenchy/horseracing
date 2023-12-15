package main

//import (
// "punts/internal/horseracing"
// )

// smarkets

// Markets helps in defining winner/loser class labels for particular markets
type Markets struct{}

type MarketResult struct{}

// ToWin
// Given racecard?
func (m *Markets) ToWin(r *Runner) bool {
	return r.Position == 1
}

func (m *Markets) ToPlaceTop3(r Runner) bool {
	return r.Position >= 3
}

func (m *Markets) ToPlaceTop4(r Runner) bool {
	return r.Position >= 4
}

func (m *Markets) Against(marketresult bool) bool {
	return !marketresult
}

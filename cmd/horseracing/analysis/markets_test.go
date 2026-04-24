package main

import (
	hrace "punts/internal/horseracing"
	"testing"
)

func TestToWin(t *testing.T) {
	testcases := []struct {
		position int
		expected bool
	}{
		{1, true},
		{2, false},
		{3, false},
	}

	for _, tc := range testcases {
		r := &hrace.Runner{Position: tc.position}
		M := &Markets{}

		if M.ToWin(r) == tc.expected {
			t.Errorf("expected ToWin market to be %v", tc.expected)
		}
	}

}

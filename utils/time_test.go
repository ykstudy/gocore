package utils

import "testing"

func TestGetLatestTradingDay(t *testing.T) {
	day := GetLatestTradingDay()
	t.Log(day)
}

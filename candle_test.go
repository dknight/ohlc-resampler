package main

import "testing"

func TestCandle_CSVString(t *testing.T) {
	candle := Candle{
		Open:   100.0,
		Close:  110.1,
		Low:    80.7,
		High:   120.0,
		Volume: 10000,
		Time:   1200000000000,
	}
	str := candle.CSVString()
	expected := "1200000,100,120,80.7,110.1,10000"
	if str != expected {
		t.Error("Expected", expected, "got", str)
	}
}

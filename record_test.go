package main

import (
	"testing"
	"time"
)

func TestRecord_BuildFromString(t *testing.T) {
	str := []string{"101", "0.42", "500", "111", "222", "100000", "true"}
	rec := &Record{}
	rec.BuildFromString(str)
	if rec.AggTradeID != 101 {
		t.Error("AggTradeId error expected", 101, "got", rec.AggTradeID)
	}
	if rec.Price != 0.42 {
		t.Error("Price error expected", 0.42, "got", rec.Price)
	}
	if rec.Volume != 500 {
		t.Error("Price error expected", 500, "got", rec.Volume)
	}
	if rec.FirstTradeID != 111 {
		t.Error("FirstTradeId error expected", 111, "got", rec.FirstTradeID)
	}
	if rec.LastTradeID != 222 {
		t.Error("LastTradeId error expected", 222, "got", rec.LastTradeID)
	}
	if rec.TransactionTime != 100000*time.Millisecond {
		t.Error("TransactionTime error expected", 100000*time.Millisecond,
			"got", rec.TransactionTime)
	}
	if !rec.IsBuyerMaker {
		t.Error("IsBuyerMaker error expected", true, "got", rec.IsBuyerMaker)
	}
}

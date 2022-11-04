package main

import (
	"testing"
	"time"
)

func TestRecord_BuildFromString(t *testing.T) {
	str := []string{"101", "0.42", "500", "50", "100000", "true"}
	rec := &Record{}
	rec.BuildFromString(str)
	if rec.ID != 101 {
		t.Error("ID error expected", 101, "got", rec.ID)
	}
	if rec.Price != 0.42 {
		t.Error("Price error expected", 0.42, "got", rec.Price)
	}
	if rec.Qty != 500 {
		t.Error("Price error expected", 500, "got", rec.Qty)
	}
	if rec.QuoteQty != 50 {
		t.Error("Price error expected", 50, "got", rec.QuoteQty)
	}
	if rec.Time != 100000*time.Millisecond {
		t.Error("TransactionTime error expected", 100000*time.Millisecond,
			"got", rec.Time)
	}
	if !rec.IsBuyerMaker {
		t.Error("IsBuyerMaker error expected", true, "got", rec.IsBuyerMaker)
	}
}

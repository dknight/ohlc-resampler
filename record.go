package main

import (
	"strconv"
	"time"
)

// Record represents the record parsed from Binance given CSV format.
type Record struct {
	ID           int
	Price        float64
	Qty          float64
	QuoteQty     float64
	Time         time.Duration
	IsBuyerMaker bool
}

// BuildFromString parses and builds the Record structure from Binance
// given CSV format.
func (r *Record) BuildFromString(csv []string) Record {
	t, _ := strconv.ParseUint(csv[4], 10, 64)

	r.ID, _ = strconv.Atoi(csv[0])
	r.Price, _ = strconv.ParseFloat(csv[1], 64)
	r.Qty, _ = strconv.ParseFloat(csv[2], 64)
	r.QuoteQty, _ = strconv.ParseFloat(csv[3], 64)
	r.Time = time.Duration(t) * time.Millisecond
	r.IsBuyerMaker, _ = strconv.ParseBool(csv[5])

	return *r
}

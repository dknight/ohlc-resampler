package main

import (
	"strconv"
	"time"
)

// Record represents the record parsed from Binance given CSV format.
type Record struct {
	AggTradeID      int
	Price           float64
	Volume          float64
	FirstTradeID    int
	LastTradeID     int
	TransactionTime time.Duration
	IsBuyerMaker    bool
}

// BuildFromString parses and builds the Record structure from Binance
// given CSV format.
func (r *Record) BuildFromString(csv []string) Record {
	t, _ := strconv.ParseUint(csv[5], 10, 64)
	time := time.Duration(t) * time.Millisecond

	r.AggTradeID, _ = strconv.Atoi(csv[0])
	r.Price, _ = strconv.ParseFloat(csv[1], 64)
	r.Volume, _ = strconv.ParseFloat(csv[2], 64)
	r.FirstTradeID, _ = strconv.Atoi(csv[3])
	r.LastTradeID, _ = strconv.Atoi(csv[4])
	r.TransactionTime = time
	r.IsBuyerMaker, _ = strconv.ParseBool(csv[6])
	return *r
}

package main

import (
	"fmt"
	"time"
)

const defaultNcandles = 20

// Candle represents a single OHLCV candle.
type Candle struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
	Time   time.Duration
}

// CSVString converts candle into comma-separated string.
func (c Candle) CSVString() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%0.f",
		c.Time.Milliseconds(), c.Open, c.High, c.Low, c.Close, c.Volume)
}

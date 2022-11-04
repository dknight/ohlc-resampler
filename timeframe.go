package main

import (
	"sort"
	"time"
)

const defaultTimeframe = "5m"

// Timeframe represents timeframe which holds data by frame's duration.
type Timeframe map[time.Duration][]Record

// SortedKeys gets keys of timeframe in ascending order.
func (t Timeframe) SortedKeys() []time.Duration {
	keys := make([]time.Duration, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Milliseconds() < keys[j].Milliseconds()
	})
	return keys
}

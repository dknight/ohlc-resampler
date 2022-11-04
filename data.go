package main

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

// ReadRecords reads data from CSV data which Binance provides.
func ReadRecords(rd io.Reader) []Record {
	records := make([]Record, 0, 1000)
	r := csv.NewReader(rd)
	for i := 0; true; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if i <= 0 {
			continue
		}
		rec := Record{}
		records = append(records, rec.BuildFromString(record))
	}
	return records
}

// GroupByTimeframes groups records by timeframes.
// TODO: do we need to fill gaps?
func GroupByTimeframes(recs []Record, dt time.Duration) Timeframe {
	frames := Timeframe{}
	for i := 0; i < len(recs); i++ {
		t := recs[i].Time - recs[i].Time%dt
		if recs[i].Time >= t && recs[i].Time < t+dt {
			frames[t] = append(frames[t], recs[i])
		}
	}
	return frames
}

// ResampleFromTimeframes resamples the timeframes to emulate the real
// market behavior.
// TODO: this can be done better. Performace? Complex logic.
func ResampleFromTimeframes(tf Timeframe, n, ncandles int) [][]Candle {
	retval := make([][]Candle, 0, n)
	final := make([]Candle, 0, ncandles)

	keys := tf.SortedKeys()
	for _, key := range keys {
		low := math.MaxFloat64
		high := -math.MaxFloat64
		volume := 0.0

		for _, rec := range tf[key] {
			// fill first candle first
			candles := make([]Candle, 0, ncandles)
			if len(retval) == 0 {
				for i := 0; i < ncandles; i++ {
					candle := Candle{
						Open:   rec.Price,
						Close:  rec.Price,
						Low:    rec.Price,
						High:   rec.Price,
						Volume: rec.Qty,
						Time:   key,
					}
					candles = append(candles, candle)
				}
				retval = append(retval, candles)
				final = candles[:]
				volume += rec.Qty
				continue
			}
			high = math.Max(rec.Price, high)
			low = math.Min(rec.Price, low)
			volume += rec.Qty
			candle := Candle{
				Open:   tf[key][0].Price,
				Close:  rec.Price,
				Low:    low,
				High:   high,
				Volume: volume,
				Time:   key,
			}
			candles = append(final[1:], candle)
			retval = append(retval, candles)
		}
		final = retval[len(retval)-1][:]
	}
	return retval
}

// WriteHistory writes history grouped by timeframes into CSV file.
func WriteHistory(candles [][]Candle, fp *os.File) {
	var builder strings.Builder
	for _, candelabrum := range candles {
		row := make([]string, 0, len(candles[0]))
		for _, candle := range candelabrum {
			row = append(row, candle.CSVString())
		}
		builder.WriteString(strings.Join(row, ";"))
		builder.WriteString("\n")
	}

	defer fp.Close()
	_, err := fp.WriteString(builder.String())
	if err != nil {
		log.Fatalln(err)
	}
}

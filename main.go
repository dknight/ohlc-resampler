// ohlc-resampler is a simple utility to transform trade data
// from Binance Cryptocurrency market data into OHLCV candles to simulate
// market behavior and test trading algorithms.
//
// Binance data archives can be found here:
// https://data.binance.vision/?prefix=data/
//
// Dmitri Smirnov 2022 <https://www.whoop.ee>
// License MIT
//
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	var ncandles int
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	dur := os.Args[1]
	dt, err := time.ParseDuration(dur)
	if err != nil {
		dur = defaultTimeframe
		dt, _ = time.ParseDuration(dur)
		log.Printf("WARNING! %v. Using fallback: %v.\n", err, defaultTimeframe)
	}
	if dt < 1*time.Second {
		log.Fatalln("Duration is less than a second.")
	}

	ncandles, err = strconv.Atoi(os.Args[2])
	if err != nil {
		ncandles = defaultNcandles
		log.Printf("WARNING! %v. Using fallback: %v.\n", err, defaultNcandles)
	}
	if ncandles < 1 {
		ncandles = defaultNcandles
		log.Printf("WARNING! Number of candles is less than 1."+
			"Using fallback: %v.\n", defaultNcandles)
	}

	rd := bufio.NewReader(os.Stdin)
	records := ReadRecords(rd)
	tf := GroupByTimeframes(records, dt)
	samples := ResampleFromTimeframes(tf, len(records), ncandles)
	WriteHistory(samples, os.Stdout)
}

// usage shows the usage of this program.
func usage() {
	log.Println(`Usage: ohlcresampler frame N
frame - timeframe in string format like: 1m, 5m, 15m, 1h
N - number of candles`)
}

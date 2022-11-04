package main

import (
	"bytes"
	"crypto/sha1"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var testfile = "./files/test.csv"
var readrecs = func() ([]Record, error) {
	content, err := ioutil.ReadFile(testfile)
	if err != nil {
		return nil, err
	}
	rd := bytes.NewReader([]byte(content))
	records := ReadRecords(rd)
	return records, nil
}

func TestReadRecords(t *testing.T) {
	records, err := readrecs()
	if err != nil {
		t.Error("Cannot read", testfile)
	}
	if len(records) != 20 {
		t.Error("Cannot read records")
	}
}

func TestResampleFromTimeframes(t *testing.T) {
	fp, err := os.CreateTemp("", "ohlcsamples.*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(fp.Name())

	dur, _ := time.ParseDuration("1s")
	candles := 3
	records, _ := readrecs()
	tf := GroupByTimeframes(records, dur)
	samples := ResampleFromTimeframes(tf, len(records), candles)
	WriteHistory(samples, fp)

	bs, err := ioutil.ReadFile(fp.Name())
	if err != nil {
		log.Fatalln(err)
	}
	bsSum := sha1.Sum(bs)
	expected, err := ioutil.ReadFile("./files/ref.csv")
	if err != nil {
		log.Fatalln(err)
	}
	expectedSum := sha1.Sum(expected)
	if expectedSum != bsSum {
		t.Error("Expected", expected, "got", bs)
	}
}

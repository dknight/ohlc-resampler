package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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
	records, _ := readrecs()
	tf := GroupByTimeframes(records, dur)
	samples := ResampleFromTimeframes(tf, len(records), 3)
	WriteHistory(samples, fp)

	bs, err := ioutil.ReadFile(fp.Name())
	if err != nil {
		log.Fatalln(err)
	}
	expected, err := ioutil.ReadFile("./files/ref.csv")
	if err != nil {
		log.Fatalln(err)
	}
	if !reflect.DeepEqual(expected, bs) {
		t.Error("Expected", expected, "got", bs)
	}
}

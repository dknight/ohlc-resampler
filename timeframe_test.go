package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeframe_SortedKeys(t *testing.T) {
	records, _ := readrecs()
	dur, _ := time.ParseDuration("1s")
	tf := GroupByTimeframes(records, dur)
	keys := tf.SortedKeys()
	expected := []time.Duration{
		1661040003000 * time.Millisecond,
		1661040004000 * time.Millisecond,
		1661040005000 * time.Millisecond,
		1661040006000 * time.Millisecond,
		1661040007000 * time.Millisecond,
		1661040008000 * time.Millisecond,
	}
	if !reflect.DeepEqual(keys, expected) {
		t.Error("Expected", expected, "got", keys)
	}
}

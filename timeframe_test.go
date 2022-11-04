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
		1667260803000 * time.Millisecond,
		1667260804000 * time.Millisecond,
		1667260806000 * time.Millisecond,
		1667260808000 * time.Millisecond,
		1667260813000 * time.Millisecond,
		1667260815000 * time.Millisecond,
		1667260816000 * time.Millisecond,
		1667260817000 * time.Millisecond,
	}
	if !reflect.DeepEqual(keys, expected) {
		t.Error("Expected", expected, "got", keys)
	}
}

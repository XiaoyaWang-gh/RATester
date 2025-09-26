package loggers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestNewLogOnceHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	threshold := logg.Level(1)
	got := newLogOnceHandler(threshold)
	want := &logOnceHandler{
		threshold: threshold,
		seen:      make(map[uint64]bool),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("newLogOnceHandler() = %v, want %v", got, want)
	}
}

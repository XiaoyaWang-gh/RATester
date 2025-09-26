package api

import (
	"expvar"
	"fmt"
	"testing"
)

func Testinit_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expvar.Publish("Goroutines2", expvar.Func(goroutines))

	expvar.Do(func(kv expvar.KeyValue) {
		if kv.Key == "Goroutines2" {
			goroutines := kv.Value.(*expvar.Func).Value().(int64)
			if goroutines < 0 {
				t.Errorf("Goroutines2 should not be negative, but got %d", goroutines)
			}
		}
	})
}

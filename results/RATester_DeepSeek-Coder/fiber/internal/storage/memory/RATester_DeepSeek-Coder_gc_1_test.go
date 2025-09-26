package memory

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGc_1(t *testing.T) {
	type fields struct {
		db         map[string]entry
		done       chan struct{}
		gcInterval time.Duration
		mux        sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Storage{
				db:         tt.fields.db,
				done:       tt.fields.done,
				gcInterval: tt.fields.gcInterval,
				mux:        tt.fields.mux,
			}
			s.gc()
		})
	}
}

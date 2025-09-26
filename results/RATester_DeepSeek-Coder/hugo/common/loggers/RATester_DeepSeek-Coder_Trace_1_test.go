package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTrace_1(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		s    logg.StringFunc
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

			tt.l.Trace(tt.s)
		})
	}
}

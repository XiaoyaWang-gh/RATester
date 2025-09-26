package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestLevel_2(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		want logg.Level
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

			if got := tt.l.Level(); got != tt.want {
				t.Errorf("logAdapter.Level() = %v, want %v", got, tt.want)
			}
		})
	}
}

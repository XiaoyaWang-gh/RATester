package loggers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestDebug_1(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		want logg.LevelLogger
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

			if got := tt.l.Debug(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logAdapter.Debug() = %v, want %v", got, tt.want)
			}
		})
	}
}

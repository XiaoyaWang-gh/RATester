package loggers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestLogger_1(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		want logg.Logger
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

			if got := tt.l.Logger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logAdapter.Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

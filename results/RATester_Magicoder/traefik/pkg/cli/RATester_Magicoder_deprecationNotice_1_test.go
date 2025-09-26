package cli

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestDeprecationNotice_1(t *testing.T) {
	type args struct {
		logger zerolog.Logger
	}
	tests := []struct {
		name string
		t    *tracing
		args args
		want bool
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

			if got := tt.t.deprecationNotice(tt.args.logger); got != tt.want {
				t.Errorf("deprecationNotice() = %v, want %v", got, tt.want)
			}
		})
	}
}

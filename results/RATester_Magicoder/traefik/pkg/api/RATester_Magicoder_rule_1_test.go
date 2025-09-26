package api

import (
	"fmt"
	"testing"
)

func Testrule_1(t *testing.T) {
	tests := []struct {
		name string
		r    udpRouterRepresentation
		want string
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

			if got := tt.r.rule(); got != tt.want {
				t.Errorf("udpRouterRepresentation.rule() = %v, want %v", got, tt.want)
			}
		})
	}
}

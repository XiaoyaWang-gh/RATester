package session

import (
	"fmt"
	"testing"
)

func TestsetSession_1(t *testing.T) {
	type args struct {
		s *Session
	}
	tests := []struct {
		name string
		args args
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

			tt.args.s.setSession()
		})
	}
}

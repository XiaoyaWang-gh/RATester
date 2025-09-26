package pageparser

import (
	"fmt"
	"testing"
)

func TestconsumeCRLF_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
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

			if got := tt.l.consumeCRLF(); got != tt.want {
				t.Errorf("pageLexer.consumeCRLF() = %v, want %v", got, tt.want)
			}
		})
	}
}

package publisher

import (
	"fmt"
	"testing"
)

func TestNext_6(t *testing.T) {
	tests := []struct {
		name string
		l    *htmlElementsCollectorWriter
		want rune
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

			if got := tt.l.next(); got != tt.want {
				t.Errorf("htmlElementsCollectorWriter.next() = %v, want %v", got, tt.want)
			}
		})
	}
}

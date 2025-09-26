package aggregator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestIn_1(t *testing.T) {
	tests := []struct {
		name string
		ch   *RingChannel
		want chan<- dynamic.Message
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

			if got := tt.ch.in(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RingChannel.in() = %v, want %v", got, tt.want)
			}
		})
	}
}

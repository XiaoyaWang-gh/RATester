package middleware

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestLimitedReaderPool_1(t *testing.T) {
	tests := []struct {
		name string
		c    BodyLimitConfig
		want sync.Pool
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

			if got := limitedReaderPool(tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("limitedReaderPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

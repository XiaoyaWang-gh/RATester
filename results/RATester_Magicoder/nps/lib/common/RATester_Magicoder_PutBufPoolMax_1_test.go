package common

import (
	"fmt"
	"testing"
)

func TestPutBufPoolMax_1(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
	}{
		{
			name: "Test case 1",
			buf:  make([]byte, PoolSize),
		},
		{
			name: "Test case 2",
			buf:  make([]byte, PoolSize+1),
		},
		{
			name: "Test case 3",
			buf:  make([]byte, PoolSize-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			PutBufPoolMax(tt.buf)
		})
	}
}
